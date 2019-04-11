package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/pscn/web-mpc/templates"
	"github.com/pscn/web-mpc/web"
)

//go:generate file2go -verbose -prefix Index templates/index.html
//go:generate file2go -verbose -prefix CSS templates/style.css
//go:generate file2go -verbose -prefix JavaScript templates/javascript.js
//go:generate file2go -verbose -prefix ThemeDefault templates/theme-default.css
//go:generate file2go -verbose -prefix ThemeJuri templates/theme-juri.css

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
		mux.HandleFunc("/", h.StaticTemplatePacked("text/html", templates.IndexContent()))
		mux.HandleFunc("/javascript.js", h.StaticPacked("text/javascript", templates.JavaScriptContent()))
		mux.HandleFunc("/theme-default.css", h.StaticPacked("text/css", templates.ThemeDefaultContent()))
		mux.HandleFunc("/theme-juri.css", h.StaticPacked("text/css", templates.ThemeJuriContent()))
		mux.HandleFunc("/style.css", h.StaticPacked("text/css", templates.CSSContent()))

	}
	mux.HandleFunc("/ws", h.Channel())
	log.Printf("Web MPC starting…")
	log.Fatal(http.ListenAndServe(*addr, mux))
}

// eof
