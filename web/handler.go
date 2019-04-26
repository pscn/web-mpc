package web

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/pscn/web-mpc/conv"
	"github.com/pscn/web-mpc/msg"

	"github.com/gorilla/websocket"
	"github.com/pscn/web-mpc/mpc"
)

// Handler a websocket, a logger and two channels come into a bar
type Handler struct {
	mpdHost   string
	mpdPort   int
	mpdPass   string
	upgrader  *websocket.Upgrader
	verbosity int
	logger    *log.Logger
}

// New handler
func New(verbosity int, checkOrigin bool, mpdHost string, mpdPass string) *Handler {
	upgrader := websocket.Upgrader{}
	if !checkOrigin {
		// disable origin check to test from static html, css & js
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	}
	host, port, _ := net.SplitHostPort(mpdHost) // FIXME: handle err
	return &Handler{
		mpdHost:   host,
		mpdPort:   conv.ToInt(port),
		mpdPass:   mpdPass,
		upgrader:  &upgrader,
		verbosity: verbosity,
		logger:    log.New(os.Stdout, fmt.Sprintf("web-mpc "), log.LstdFlags|log.Lshortfile),
	}
}

// StaticPacked serves content with contenType
func (h *Handler) StaticPacked(contentType string, content *[]byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", contentType)
		w.Write(*content)
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
func (h *Handler) StaticTemplatePacked(contentType string, content *[]byte) http.HandlerFunc {
	tmpl := template.Must(template.New("").Parse(string(*content)))

	return func(w http.ResponseWriter, r *http.Request) {
		p := map[string]interface{}{
			"ws": "ws://" + r.Host + "/ws",
		}
		w.Header().Set("Content-type", contentType)
		tmpl.Execute(w, p)
	}
}

// StaticTemplateFile serves content with contenType
func (h *Handler) StaticTemplateFile(contentType string, fileName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := map[string]interface{}{
			"ws": "ws://" + r.Host + "/ws",
		}
		w.Header().Set("Content-type", contentType)
		dat, err := ioutil.ReadFile(path.Join("templates", fileName))
		if err != nil {
			h.logger.Fatal(err)
		}
		tmpl := template.Must(template.New("").Parse(string(dat)))
		tmpl.Execute(w, p)
	}
}

func (h *Handler) writeMessage(ws *websocket.Conn, msg *msg.Message) error {
	if msg == nil {
		h.logger.Println("cowardly refusing to work with nil")
		return nil
	}
	data, err := msg.JSON()
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
func (h *Handler) Channel() http.HandlerFunc {
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
		client, err := mpc.New(h.mpdHost, h.mpdPort, h.mpdPass, h.logger)
		if err != nil {
			h.logger.Println("mpc:", err)
			// FIXME: either the host & port for MPD is wrong, or MPD is
			// down / restarting
			// we could try again after some time?
			// right now the user needs to reload the page to try again
			h.writeMessage(ws, msg.InfoMsg(
				fmt.Sprintf("failed to connect to MPD: %v", err)))
			return
		}
		client.Stats()
		defer client.Close()

		// channel for commands from the webclient
		wc := make(chan *mpc.Command, 10)

		go func() {
			defer func() {
				close(wc)
				h.logger.Println("stopping webclient loop")
			}()
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
		currentPage := 1
		currentSearchPage := 1
		lastSearch := ""
		h.writeMessage(ws, client.Version())
		h.writeMessage(ws, client.Update(currentPage))

		ping := time.Tick(5 * time.Second)
		for {
			select {
			case <-ping:
				err := client.Ping()
				if err != nil {
					h.logger.Println("ping failed:", err)
					err = client.Connect()
					if err != nil {
						h.logger.Println("connect failed:", err)
					}
				}
			case event := <-client.Event:
				h.logger.Println("event:", event)
				switch event {
				case "player", "playlist", "options":
					// send the playlist before the status, to have "nextsong" work correctly
					h.writeMessage(ws, client.Update(currentPage))
				}

			case cmd := <-wc:
				if cmd == nil {
					// wc closed → exit
					return
				}
				h.logger.Printf("cmd: %s\n", cmd)
				switch cmd.Command {
				case mpc.Play:
					if cmd.Data != "" {
						err = client.Play(conv.ToInt(cmd.Data))
					} else {
						err = client.Play(-1)
					}

				case mpc.Resume:
					err = client.Pause(false)

				case mpc.Pause:
					err = client.Pause(true)

				case mpc.Stop:
					err = client.Stop()

				case mpc.Next:
					err = client.Next()

				case mpc.Previous:
					err = client.Previous()

				case mpc.Add:
					err = client.Add(cmd.Data)

				case mpc.Remove:
					err = client.Delete(conv.ToInt(cmd.Data), -1)

				case mpc.Clean:
					err = client.Clean()

				case mpc.Prio:
					args := strings.Split(cmd.Data, ":")
					err = client.SetPriority(conv.ToInt(args[0]), conv.ToInt(args[1]), -1)

				case mpc.AddPrio:
					args := strings.Split(cmd.Data, ":")
					err = client.AddPrio(conv.ToInt(args[0]), args[1])

				case mpc.ModeConsume, mpc.ModeRepeat, mpc.ModeSingle, mpc.ModeRandom:
					target := true
					if cmd.Data == "disable" {
						target = false
					}
					switch cmd.Command {
					case mpc.ModeConsume:
						client.Consume(target)

					case mpc.ModeRepeat:
						client.Repeat(target)

					case mpc.ModeSingle:
						client.Single(target)

					case mpc.ModeRandom:
						client.Random(target)
					}

				case mpc.UpdateRequest:
					currentPage = conv.ToInt(cmd.Data)
					h.writeMessage(ws, client.Update(currentPage))

				case mpc.Search:
					currentSearchPage = 1
					lastSearch = cmd.Data
					h.writeMessage(ws, client.Search(cmd.Data, currentSearchPage))

				case mpc.SearchPage:
					currentSearchPage = conv.ToInt(cmd.Data)
					h.writeMessage(ws, client.Search(lastSearch, currentSearchPage))

				case mpc.Browse:
					h.writeMessage(ws, client.ListDirectory(cmd.Data))

				case mpc.Playlists:
					currentSearchPage = 1
					h.writeMessage(ws, client.ListPlaylists(currentSearchPage))
				case mpc.PlaylistsPage:
					currentSearchPage = conv.ToInt(cmd.Data)
					h.writeMessage(ws, client.ListPlaylists(currentSearchPage))
				}
				if err != nil {
					h.logger.Println("command error:", err)
				}
			}
		}
	}
}

// eof
