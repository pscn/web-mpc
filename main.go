package main

import (
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/karrick/golf"
	"github.com/pscn/web-mpc/server"
	"github.com/pscn/web-mpc/templates"
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
	h := server.New(verbosity, !*devel, *mpdHost, *mpdPass)
	mux := http.NewServeMux()
	// read templates and add listener
	if *devel {
		mux.HandleFunc("/", h.StaticTemplateFile("text/html", "index.html"))
		for _, file := range templates.Filenames() {
			f := filepath.Base(file)
			switch {
			case strings.HasSuffix(f, "js"):
				mux.HandleFunc("/"+f, h.StaticFile("text/javascript", f))
			case strings.HasSuffix(f, "css"):
				mux.HandleFunc("/"+f, h.StaticFile("text/css", f))
			case strings.HasSuffix(f, "ico"):
				mux.HandleFunc("/"+f, h.StaticFile("image/x-icon", f))
			}
		}
	} else {
		mux.HandleFunc("/", h.StaticTemplatePacked("text/html",
			templates.ContentMust("templates/index.html")))
		for _, file := range templates.Filenames() {
			f := filepath.Base(file)
			switch {
			case strings.HasSuffix(f, "js"):
				mux.HandleFunc("/"+f, h.StaticPacked("text/javascript",
					templates.ContentMust(file)))
			case strings.HasSuffix(f, "css"):
				mux.HandleFunc("/"+f, h.StaticPacked("text/css",
					templates.ContentMust(file)))
			case strings.HasSuffix(f, "ico"):
				mux.HandleFunc("/"+f, h.StaticPacked("image/x-icon",
					templates.ContentMust(file)))
			}
		}
	}
	mux.HandleFunc("/ws", h.Channel())
	log.Printf("Web MPC starting…")
	log.Fatal(http.ListenAndServe(*addr, mux))
}

// eof
