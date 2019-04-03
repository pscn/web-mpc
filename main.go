package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gobuffalo/packr"
	"github.com/gorilla/websocket"
	"github.com/pscn/web-mpc/web"
)

var addr = flag.String("addr", "127.0.0.1:8080", "http service address")
var mpdHost = flag.String("mpd", "127.0.0.1:6600", "MPD service address")
var pass = flag.String("password", "", "MPD password")
var devel = flag.Bool("local", false,
	"serves html, jss & css from the local templates directory")

var upgrader = websocket.Upgrader{} // FIXME: what is this, what does it do?

var box = packr.NewBox("./templates")
var verbosity = 2

func main() {
	flag.Parse()
	log.SetFlags(0)
	// disable origin check to test from static html, css & js (FIXME: remove this)
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	h := web.New(&upgrader, verbosity)
	mux := http.NewServeMux()
	// read templates and add listener
	if *devel {
		mux.HandleFunc("/", h.StaticTemplateFile("text/html", "index.html"))
		mux.HandleFunc("/script.js", h.StaticFile("text/javascript", "script.js"))
		mux.HandleFunc("/style.css", h.StaticFile("text/css", "style.css"))
	} else {
		mux.HandleFunc("/", h.StaticTemplatePacked("text/html", "index.html", &box))
		mux.HandleFunc("/script.js", h.StaticPacked("text/javascript", "script.js", &box))
		mux.HandleFunc("/style.css", h.StaticPacked("text/css", "style.css", &box))
	}
	mux.HandleFunc("/echo", h.Channel(*mpdHost, *pass))

	log.Fatal(http.ListenAndServe(*addr, mux))
}
