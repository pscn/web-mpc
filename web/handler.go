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
	"time"

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

// StaticPacked serves content with contenType
func (h *Handler) StaticPacked(contentType string, fileName string, box *packr.Box) http.HandlerFunc {
	tmplStr, err := (*box).FindString(fileName)
	if err != nil {
		h.logger.Println("box error:", err)
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
		h.logger.Println("box error:", err)
		return nil
	}
	tmpl := template.Must(template.New("").Parse(tmplStr))

	return func(w http.ResponseWriter, r *http.Request) {
		p := map[string]interface{}{"ws": "ws://" + r.Host + "/echo"}
		w.Header().Set("Content-type", contentType)
		tmpl.Execute(w, p)
	}
}

// StaticTemplateFile serves content with contenType
func (h *Handler) StaticTemplateFile(contentType string, fileName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := map[string]interface{}{"ws": "ws://" + r.Host + "/echo"}
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
	if msg == nil {
		h.logger.Println("cowardly refusing to work with nil")
		return nil
	}
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
			// FIXME: either the host & port for MPD is wrong, or MPD is
			// down / restarting
			// we could try again after some time?
			// right now the user needs to reload the page to try again
			h.writeMessage(ws, mpc.NewInfo(
				fmt.Sprintf("failed to connect to MPD: %v", err)))
			return
		}
		defer client.Close()

		// channel for commands from the webclient
		wc := make(chan *mpc.Command, 10)

		go func() {
			defer close(wc)
			for {
				cmd, err := h.readCommand(ws)
				if err != nil {
					h.logger.Println("read error:", err)
					return
				}
				wc <- cmd
			}
		}()

		// update the web client with the current status
		h.writeMessage(ws, client.Status())
		h.writeMessage(ws, client.CurrentSong())
		h.writeMessage(ws, client.CurrentPlaylist())

		ping := time.Tick(5 * time.Second)
		for {
			select {
			case event := <-*client.Event:
				h.logger.Println("event:", event)
				switch event {
				case "player", "playlist":
					h.writeMessage(ws, client.Status())
					h.writeMessage(ws, client.CurrentSong())
					h.writeMessage(ws, client.CurrentPlaylist())
				}
			case cmd := <-wc:
				h.logger.Printf("cmd: %s\n", cmd)
				switch cmd.Command {
				case mpc.Play:
					if cmd.Data != "" {
						err = client.Play(helpers.ToInt(cmd.Data))
					} else {
						err = client.Play(-1)
					}
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
				case mpc.Add:
					err = client.Add(cmd.Data)
				case mpc.Remove:
					err = client.RemovePlaylistEntry(helpers.ToInt(cmd.Data))
				case mpc.StatusRequest:
					h.writeMessage(ws, client.Status())
				case mpc.Search:
					h.writeMessage(ws, client.Search(cmd.Data))
				}
				if err != nil {
					h.logger.Println("command error:", err)
				}
			case <-ping:
				err := client.Ping()
				if err != nil {
					h.logger.Println("ping failed:", err)
				}
			}
		}
	}
}

// eof
