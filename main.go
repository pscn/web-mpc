package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/pscn/web-mpc/web"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{} // FIXME: what is this, what does it do?

func home(w http.ResponseWriter, r *http.Request) {
	web.Template.Execute(w, "ws://"+r.Host+"/echo")
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
