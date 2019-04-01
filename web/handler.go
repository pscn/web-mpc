package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

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
		if h.verbosity > 5 {
			logger.Printf("Command: %v", cmd.Command)
		}
		switch cmd.Command {
		case "play":
			err = client.Play(-1)
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
			rc <- mpc.NewStatus(client.Status())
		}
		if strings.HasPrefix(cmd.Command, "add") {
			file := cmd.Command[3:]
			if h.verbosity > 5 {
				logger.Printf("%s => %s == %s", cmd.Command, cmd.Command[3:], file)
			}
			err = client.Add(file)
		}
		if strings.HasPrefix(cmd.Command, "play") && len(cmd.Command) > 4 {
			nr := helpers.ToInt(cmd.Command[4:])
			if h.verbosity > 5 {
				logger.Printf("%s => %s == %d", cmd.Command, cmd.Command[4:], nr)
			}
			err = client.Play(nr)
		}
		if strings.HasPrefix(cmd.Command, "remove") {
			nr := helpers.ToInt64(cmd.Command[6:])
			if h.verbosity > 5 {
				logger.Printf("%s => %s == %d", cmd.Command, cmd.Command[6:], nr)
			}
			err = client.RemovePlaylistEntry(nr)
		}
		if strings.HasPrefix(cmd.Command, "search") {
			search := cmd.Command[6:]
			logger.Printf("%s => %s == %s", cmd.Command, cmd.Command[6:], search)
			rc <- mpc.NewSearchResult(client.Search(search))
		}

		if err != nil {
			logger.Printf("Command error: %v", err)
		}
	}
}
