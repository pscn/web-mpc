package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/pscn/web-mpc/helpers"

	"github.com/gorilla/websocket"
	"github.com/pscn/web-mpc/mpc"
)

// Handler a websocket, a logger and two channels come into a bar
type Handler struct {
	upgrader  *websocket.Upgrader
	verbosity int
	websocket *websocket.Conn
	logger    *log.Logger
}

// New *Handler
func New(upgrader *websocket.Upgrader, verbosity int) *Handler {
	return &Handler{
		upgrader:  upgrader,
		verbosity: verbosity,
	}
}

func (h *Handler) writeMessage(msg *mpc.Message) error {
	data, err := json.Marshal(msg)
	if err != nil {
		h.logger.Println("marshal:", err)
		return err
	}
	err = h.websocket.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		h.logger.Println("write:", err)
	}
	return err
}
func (h *Handler) readCommand() (*mpc.Command, error) {
	var cmd mpc.Command
	_, data, err := h.websocket.ReadMessage()
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
		var err error
		h.logger = log.New(os.Stdout, fmt.Sprintf("web-mpc %s ", r.RemoteAddr), log.LstdFlags|log.Lshortfile)
		defer func() {
			if r := recover(); r != nil {
				h.logger.Println("recovered", r)
			}
		}()
		h.logger.Printf("handling")
		defer h.logger.Printf("stop handling")

		// open websocket
		h.websocket, err = h.upgrader.Upgrade(w, r, nil)
		if err != nil {
			h.logger.Println("upgrade:", err)
			return
		}
		defer h.websocket.Close()

		// open connection to mpc
		client, err := mpc.New(mpdHost, mpdPass, h.logger)
		if err != nil {
			h.logger.Println("mpc:", err)
			h.writeMessage(mpc.NewStringEvent(
				fmt.Sprintf("failed to connect to MPD: %v", err)))
			return
		}
		defer client.Close()

		// return channel
		rc := make(chan *mpc.Message, 1)
		defer close(rc)

		go client.EventLoop(rc)

		go func() { // read events from mpc

			for event := range rc {
				if h.verbosity > 5 {
					h.logger.Printf("Event: %d\n", event.Type)
				}
				switch event.Type {
				case mpc.Error:
					h.logger.Println("error:", event.Error())
					break
				case mpc.Info:
					if h.verbosity > 5 {
						h.logger.Println("string:", event.String())
					}
				case mpc.Status:
					if h.verbosity > 5 {
						h.logger.Println("status:", event.Status())
					}
				case mpc.CurrentSong:
					if h.verbosity > 5 {
						h.logger.Println("current song:", event.CurrentSong())
					}
				case mpc.Playlist:
					if h.verbosity > 5 {
						h.logger.Println("current playlist:", event.CurrentPlaylist())
					}
				}
				h.writeMessage(event)
			}
		}()

		for { // read commands from the webpage
			cmd, err := h.readCommand() // FIXME: handle err
			if h.verbosity > 5 {
				h.logger.Printf("recv: %v", *cmd)
			}
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
