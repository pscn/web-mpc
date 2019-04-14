package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/pscn/web-mpc/templates"
	"github.com/pscn/web-mpc/web"
)

//go:generate file2go -v -t -o templates/files.go templates/*.html templates/*.js templates/*.css

var addr = flag.String("addr", ":8666", "http service address")
var mpdHost = flag.String("mpd", "127.0.0.1:6600", "MPD service address")
var mpdPass = flag.String("password", "", "MPD password")
var devel = flag.Bool("local", false,
	"serves html, jss & css from the local templates directory")

var verbosity = 2

func main() {
	flag.Parse()
	log.SetFlags(0)
	h := web.New(verbosity, !*devel, *mpdHost, *mpdPass)
	mux := http.NewServeMux()
	// read templates and add listener
	if *devel {
		mux.HandleFunc("/", h.StaticTemplateFile("text/html", "index.html"))
		mux.HandleFunc("/javascript.js", h.StaticFile("text/javascript", "javascript.js"))
		mux.HandleFunc("/theme-default.css", h.StaticFile("text/css", "theme-default.css"))
		mux.HandleFunc("/theme-juri.css", h.StaticFile("text/css", "theme-juri.css"))
		mux.HandleFunc("/style.css", h.StaticFile("text/css", "style.css"))
	} else {
		mux.HandleFunc("/", h.StaticTemplatePacked("text/html",
			templates.ContentMust("templates/index.html")))
		mux.HandleFunc("/javascript.js", h.StaticPacked("text/javascript",
			templates.ContentMust("templates/javascript.js")))
		mux.HandleFunc("/theme-default.css", h.StaticPacked("text/css",
			templates.ContentMust("templates/theme-default.css")))
		mux.HandleFunc("/theme-juri.css", h.StaticPacked("text/css",
			templates.ContentMust("templates/theme-juri.css")))
		mux.HandleFunc("/style.css", h.StaticPacked("text/css",
			templates.ContentMust("templates/style.css")))

	}
	mux.HandleFunc("/ws", h.Channel())
	log.Printf("Web MPC starting…")
	log.Fatal(http.ListenAndServe(*addr, mux))
}

// eof
