package web

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/pscn/web-mpc/helpers"

	"github.com/gorilla/websocket"
	"github.com/pscn/web-mpc/mpc"
)

// Handler a websocket, a logger and two channels come into a bar
type Handler struct {
	upgrader *websocket.Upgrader
	logger   *log.Logger
}

// New *Handler
func New(upgrader *websocket.Upgrader, logger *log.Logger) *Handler {
	return &Handler{
		upgrader: upgrader,
		logger:   logger,
	}
}

// Channel to websocket
func (h *Handler) Channel(w http.ResponseWriter, r *http.Request) {
	h.logger.Printf("handling: %s", r.Host)
	defer h.logger.Printf("stop handling: %s", r.Host)

	// open websocket
	c, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		h.logger.Println("upgrade:", err)
		return
	}
	defer c.Close()

	// open connection to mpc
	client, err := mpc.New("192.168.0.102", 6600, "", h.logger)
	if err != nil {
		h.logger.Println("mpc:", err)
		return
	}
	defer client.Close()

	rc := make(chan *mpc.Event, 1)
	defer close(rc)

	go func() { // read events from mpc
		go func() {
			client.EventLoop(rc)
			return
		}()
		for event := range rc {
			h.logger.Printf("Event: %d\n", event.Type)
			var data []byte
			var err error
			switch event.Type {
			case mpc.EventTypeError:
				h.logger.Println("error:", event.Error())
				break
			case mpc.EventTypeString:
				h.logger.Println("string:", event.String())
			case mpc.EventTypeStatus:
				h.logger.Println("status:", event.Status())
			case mpc.EventTypeCurrentSong:
				h.logger.Println("current song:", event.CurrentSong())
			case mpc.EventTypeCurrentPlaylist:
				h.logger.Println("current playlist:", event.CurrentPlaylist())
			}
			data, err = json.Marshal(event)
			if err != nil {
				h.logger.Println("marshal:", err)
				break
			}
			if data != nil {
				h.logger.Println("writing:", string(data))
				c.WriteMessage(websocket.TextMessage, []byte(data))
				if err != nil {
					h.logger.Println("write:", err)
					break
				}
			}
		}
	}()

	for { // read commands from the webpage
		_, data, err := c.ReadMessage()
		if err != nil {
			h.logger.Println("read:", err)
			break
		}
		h.logger.Printf("recv: %v", string(data))
		var cmd mpc.Command
		err = json.Unmarshal(data, &cmd)
		h.logger.Printf("Command: %v", cmd.Command)
		switch cmd.Command {
		case "play":
			err = client.Play()
		case "resume":
			err = client.Resume()
		case "pause":
			err = client.Pause()
		case "stop":
			err = client.Stop()
		case "next":
			err = client.Next()
		case "previous":
			err = client.Previous()
		case "status":
			rc <- mpc.NewStatusEvent(client.Status())
		}
		if strings.HasPrefix(cmd.Command, "remove") {
			nr := helpers.ToInt64(cmd.Command[6:])
			h.logger.Printf("%s => %s == %d", cmd.Command, cmd.Command[6:], nr)
			err = client.RemovePlaylistEntry(nr)
		}
		if err != nil {
			h.logger.Printf("Command error: %v", err)
		}
	}
}
