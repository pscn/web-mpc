package main

import (
	"log"
	"net/http"

	"github.com/karrick/golf"
	"github.com/pscn/web-mpc/templates"
	"github.com/pscn/web-mpc/web"
)

//go:generate file2go -v -t -o templates/files.go templates/*.html templates/*.js templates/*.css templates/*.ico

var addr = golf.StringP('a', "addr", ":8666", "http service address")
var mpdHost = golf.StringP('m', "mpd", "127.0.0.1:6600", "MPD service address")
var mpdPass = golf.StringP('p', "password", "", "MPD password")
var devel = golf.BoolP('l', "local", false,
	"serves html, jss & css from the local templates directory")

var verbosity = 2

func main() {
	golf.Parse()
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
		mux.HandleFunc("/favicon.ico", h.StaticFile("image/x-icon", "favicon.ico"))
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
		mux.HandleFunc("/favicon.ico", h.StaticPacked("image/x-icon",
			templates.ContentMust("templates/favicon.ico")))

	}
	mux.HandleFunc("/ws", h.Channel())
	log.Printf("Web MPC starting…")
	log.Fatal(http.ListenAndServe(*addr, mux))
}

// eof
