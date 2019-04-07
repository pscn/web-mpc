package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gobuffalo/packr"
	"github.com/pscn/web-mpc/web"
)

var addr = flag.String("addr", ":8666", "http service address")
var mpdHost = flag.String("mpd", "127.0.0.1:6600", "MPD service address")
var mpdPass = flag.String("password", "", "MPD password")
var devel = flag.Bool("local", false,
	"serves html, jss & css from the local templates directory")

var box = packr.NewBox("./templates")
var verbosity = 2

func main() {
	flag.Parse()
	log.SetFlags(0)
	h := web.New(verbosity, !*devel, *mpdHost, *mpdPass)
	mux := http.NewServeMux()
	// read templates and add listener
	if *devel {
		mux.HandleFunc("/", h.EnsureCookie(h.StaticTemplateFile("text/html", "index.html")))
		mux.HandleFunc("/script.js", h.StaticFile("text/javascript", "script.js"))
		mux.HandleFunc("/theme-default.css", h.StaticFile("text/css", "theme-default.css"))
		mux.HandleFunc("/theme-juri.css", h.StaticFile("text/css", "theme-juri.css"))
		mux.HandleFunc("/style.css", h.StaticFile("text/css", "style.css"))
	} else {
		mux.HandleFunc("/", h.EnsureCookie(h.StaticTemplatePacked("text/html", "index.html", &box)))
		mux.HandleFunc("/script.js", h.StaticPacked("text/javascript", "script.js", &box))
		mux.HandleFunc("/theme-default.css", h.StaticPacked("text/css", "theme-default.css", &box))
		mux.HandleFunc("/theme-juri.css", h.StaticPacked("text/css", "theme-juri.css", &box))
		mux.HandleFunc("/style.css", h.StaticPacked("text/css", "style.css", &box))
	}
	mux.HandleFunc("/ws", h.Channel())

	log.Fatal(http.ListenAndServe(*addr, mux))
}

// eof
