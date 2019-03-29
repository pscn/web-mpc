package web

import (
	"log"
	"net/http"

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

	go func() { // read events from mpc
		rc := make(chan *mpc.Event, 1)
		defer close(rc)
		go func() {
			client.EventLoop(rc)
			return
		}()
		for event := range rc {
			h.logger.Printf("Event: %d\n", event.Type)
			switch event.Type {
			case mpc.EventError:
				h.logger.Println("error:", event.Data)
				h.logger.Println("error:", event.Error())
				break
			case mpc.EventString:
				h.logger.Println("string:", event.String())
				c.WriteMessage(websocket.TextMessage, []byte(event.String()))
				if err != nil {
					h.logger.Println("write:", err)
					break
				}
			}
		}
		// open watcher for mpc
		//			c.WriteMessage(websocket.TextMessage, msg.Data)
		//			if err != nil {
		//				h.logger.Println("write:", err)
		//				break
		//			}
	}()

	for { // read commands from the webpage
		_, data, err := c.ReadMessage()
		if err != nil {
			h.logger.Println("read:", err)
			break
		}
		h.logger.Printf("recv: %v", string(data))
	}
}
