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
}

// New *Handler
func New(upgrader *websocket.Upgrader, verbosity int) *Handler {
	return &Handler{
		upgrader:  upgrader,
		verbosity: verbosity,
	}
}

// Channel to websocket
func (h *Handler) Channel(w http.ResponseWriter, r *http.Request) {
	logger := log.New(os.Stdout, fmt.Sprintf("web-mpc %s ", r.RemoteAddr), log.LstdFlags|log.Lshortfile)
	defer func() {
		if r := recover(); r != nil {
			logger.Println("recovered", r)
		}
	}()
	logger.Printf("handling")
	defer logger.Printf("stop handling")

	// open websocket
	c, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Println("upgrade:", err)
		return
	}
	defer c.Close()

	// open connection to mpc
	client, err := mpc.New("192.168.0.102", 6600, "", logger)
	if err != nil {
		logger.Println("mpc:", err)
		return
	}
	defer client.Close()

	rc := make(chan *mpc.Message, 1)
	defer close(rc)

	go client.EventLoop(rc)

	go func() { // read events from mpc

		for event := range rc {
			if h.verbosity > 5 {
				logger.Printf("Event: %d\n", event.Type)
			}
			var data []byte
			var err error
			switch event.Type {
			case mpc.Error:
				logger.Println("error:", event.Error())
				break
			case mpc.Info:
				if h.verbosity > 5 {
					logger.Println("string:", event.String())
				}
			case mpc.Status:
				if h.verbosity > 5 {
					logger.Println("status:", event.Status())
				}
			case mpc.CurrentSong:
				if h.verbosity > 5 {
					logger.Println("current song:", event.CurrentSong())
				}
			case mpc.Playlist:
				if h.verbosity > 5 {
					logger.Println("current playlist:", event.CurrentPlaylist())
				}
			}
			data, err = json.Marshal(event)
			if err != nil {
				logger.Println("marshal:", err)
				break
			}
			if data != nil {
				if h.verbosity > 5 {
					logger.Println("writing:", string(data))
				}
				c.WriteMessage(websocket.TextMessage, []byte(data))
				if err != nil {
					logger.Println("write:", err)
					break
				}
			}
		}
	}()

	for { // read commands from the webpage
		_, data, err := c.ReadMessage()
		if err != nil {
			logger.Println("read:", err)
			break
		}
		if h.verbosity > 5 {
			logger.Printf("recv: %v", string(data))
		}
		var cmd mpc.Command
		err = json.Unmarshal(data, &cmd)
		logger.Printf("Command: %v", cmd)
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
			logger.Printf("Command error: %v", err)
		}
	}
}
