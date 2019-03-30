package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/pscn/web-mpc/mpc"

	"github.com/gorilla/websocket"
	"github.com/pscn/web-mpc/web"
)

var addr = flag.String("addr", "192.168.0.111:8080", "http service address")

var upgrader = websocket.Upgrader{} // FIXME: what is this, what does it do?

func home(w http.ResponseWriter, r *http.Request) {
	p := map[string]interface{}{
		"ws":              "ws://" + r.Host + "/echo",
		"error":           mpc.EventTypeError,
		"string":          mpc.EventTypeString,
		"status":          mpc.EventTypeStatus,
		"currentSong":     mpc.EventTypeCurrentSong,
		"currentPlaylist": mpc.EventTypeCurrentPlaylist,
	}
	web.Template.Execute(w, p)
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	logger := log.New(os.Stdout, "web-mpc ", log.LstdFlags|log.Lshortfile)
	h := web.New(&upgrader, logger)
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/echo", h.Channel)

	log.Fatal(http.ListenAndServe(*addr, mux))
}
