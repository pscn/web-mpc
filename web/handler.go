package web

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/gobuffalo/packr"
	"github.com/pscn/web-mpc/helpers"

	"github.com/gorilla/websocket"
	"github.com/pscn/web-mpc/mpc"
)

// Handler a websocket, a logger and two channels come into a bar
type Handler struct {
	upgrader  *websocket.Upgrader
	verbosity int
	logger    *log.Logger
}

// New *Handler
func New(upgrader *websocket.Upgrader, verbosity int) *Handler {
	return &Handler{
		upgrader:  upgrader,
		verbosity: verbosity,
	}
}

// getTemplateParameters returns common "constants" shared between go
// and javascript
func getTemplateParameters() *map[string]interface{} {
	p := map[string]interface{}{
		"error":           mpc.Error,
		"string":          mpc.Info,
		"status":          mpc.Status,
		"currentSong":     mpc.CurrentSong,
		"currentPlaylist": mpc.Playlist,
	}
	return &p
}

// StaticPacked serves content with contenType
func (h *Handler) StaticPacked(contentType string, fileName string, box *packr.Box) http.HandlerFunc {
	tmplStr, err := (*box).FindString(fileName)
	if err != nil {
		h.logger.Printf("box error: %v", err)
		return nil
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", contentType)
		w.Write([]byte(tmplStr))
	}
}

// StaticFile serves fileName with contenType
func (h *Handler) StaticFile(contentType string, fileName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", contentType)
		dat, err := ioutil.ReadFile(path.Join("templates", fileName))
		if err != nil {
			h.logger.Fatal(err)
		}
		w.Write(dat)
	}
}

// StaticTemplatePacked serves content with contenType
func (h *Handler) StaticTemplatePacked(contentType string, fileName string, box *packr.Box) http.HandlerFunc {
	tmplStr, err := (*box).FindString(fileName)
	if err != nil {
		h.logger.Printf("box error: %v", err)
		return nil
	}
	tmpl := template.Must(template.New("").Parse(tmplStr))
	p := *getTemplateParameters()
	return func(w http.ResponseWriter, r *http.Request) {
		p["ws"] = "ws://" + r.Host + "/echo"
		w.Header().Set("Content-type", contentType)
		tmpl.Execute(w, p)
	}
}

// StaticTemplateFile serves content with contenType
func (h *Handler) StaticTemplateFile(contentType string, fileName string) http.HandlerFunc {
	p := *getTemplateParameters()
	return func(w http.ResponseWriter, r *http.Request) {
		p["ws"] = "ws://" + r.Host + "/echo"
		w.Header().Set("Content-type", contentType)
		dat, err := ioutil.ReadFile(path.Join("templates", fileName))
		if err != nil {
			h.logger.Fatal(err)
		}
		tmpl := template.Must(template.New("").Parse(string(dat)))
		tmpl.Execute(w, p)
	}
}

func (h *Handler) writeMessage(ws *websocket.Conn, msg *mpc.Message) error {
	data, err := json.Marshal(msg)
	if err != nil {
		h.logger.Println("marshal:", err)
		return err
	}
	err = ws.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		h.logger.Println("write:", err)
	}
	return err
}

func (h *Handler) readCommand(ws *websocket.Conn) (*mpc.Command, error) {
	var cmd mpc.Command
	_, data, err := ws.ReadMessage()
	if err != nil {
		h.logger.Println("read:", err)
		return &cmd, err
	}
	err = json.Unmarshal(data, &cmd)
	return &cmd, err
}

// Channel to websocket
func (h *Handler) Channel(mpdHost string, mpdPass string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.logger = log.New(os.Stdout, fmt.Sprintf("web-mpc %s ", r.RemoteAddr), log.LstdFlags|log.Lshortfile)

		defer func() { // FIXME: not very nice, but better then crashing eh?
			if r := recover(); r != nil {
				h.logger.Println("recovered", r)
			}
		}()
		h.logger.Printf("handling")
		defer h.logger.Printf("stop handling")

		// open websocket
		ws, err := h.upgrader.Upgrade(w, r, nil)
		if err != nil {
			h.logger.Println("upgrade:", err)
			return
		}
		defer ws.Close()

		// open connection to mpc
		client, err := mpc.New(mpdHost, mpdPass, h.logger)
		if err != nil {
			h.logger.Println("mpc:", err)
			h.writeMessage(ws, mpc.NewInfo(
				fmt.Sprintf("failed to connect to MPD: %v", err)))
			return
		}
		defer client.Close()

		// return channel mpc
		rc := make(chan *mpc.Message, 10)
		defer close(rc)

		// channel for websocket
		wc := make(chan *mpc.Command, 10)
		defer close(wc)

		go client.EventLoop(rc)

		go func() {
			for {
				cmd, err := h.readCommand(ws)
				if err != nil {
					h.logger.Printf("read error: %v", err)
					break
				}
				wc <- cmd
			}
			close(wc)
		}()

		for {
			select {
			case msg := <-rc:
				h.logger.Printf("event: %s\n", msg.String()) // FIXME: why does it not use String() automatically?
				switch msg.Type {
				case mpc.Error:
					h.logger.Println("error:", msg.Error())
					break
				}
				h.writeMessage(ws, msg)
			case cmd := <-wc:
				h.logger.Printf("cmd: %v\n", *cmd)
				switch cmd.Command {
				case mpc.Play:
					var nr int
					if cmd.Data != "" {
						nr = helpers.ToInt(cmd.Data)
					} else {
						nr = -1
					}
					err = client.Play(nr)
				case mpc.Resume:
					err = client.Resume()
				case mpc.Pause:
					err = client.Pause()
				case mpc.Stop:
					err = client.Stop()
				case mpc.Next:
					err = client.Next()
				case mpc.Previous:
					err = client.Previous()
				case mpc.StatusRequest:
					rc <- mpc.NewStatus(client.Status())
				case mpc.Add:
					err = client.Add(cmd.Data)
				case mpc.Remove:
					nr := helpers.ToInt64(cmd.Data)
					err = client.RemovePlaylistEntry(nr)
				case mpc.Search:
					rc <- mpc.NewSearchResult(client.Search(cmd.Data))
				}
				if err != nil {
					h.logger.Printf("Command error: %v", err)
				}
			}
		}
	}
}
