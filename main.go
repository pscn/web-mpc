package main

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/karrick/golf" // flags with short options
	"github.com/pscn/web-mpc/backend"
	"github.com/pscn/web-mpc/backend/server"
)

//go:generate file2go -v -t -o backend/files.go frontend/dist/*.html frontend/dist/*.ico frontend/dist/js/*.js frontend/dist/js/*.map frontend/dist/css/*.css

var addr = golf.StringP('a', "addr", ":8666", "http service address")
var mpdHost = golf.StringP('m', "mpd", "127.0.0.1:6600", "MPD service address")
var mpdPass = golf.StringP('p', "password", "", "MPD password")
var devel = golf.BoolP('d', "devel", false, "development mode (do not check origin for websocket)")

var verbosity = 2

func main() {
	golf.Parse()
	log.SetFlags(0)
	h := server.New(verbosity, *mpdHost, *mpdPass, !*devel)
	mux := http.NewServeMux()
	// read templates and add listener
	suffix2contentType := map[string]string{
		".html": "text/html",
		".js":   "text/javascript",
		".map":  "application/octet-stream",
		".css":  "text/css",
		".ico":  "image/x-icon",
	}
	for _, file := range backend.Filenames() {
		f, err := filepath.Rel("frontend/dist", file)
		if err != nil {
			panic(err)
		}
		f = filepath.ToSlash(f)
		if f == "index.html" {
			mux.HandleFunc("/", h.StaticTemplatePacked("text/html",
				backend.ContentMust(file)))
			continue
		}
		if ct, ok := suffix2contentType[filepath.Ext(f)]; ok {
			mux.HandleFunc("/"+f, h.StaticPacked(ct,
				backend.ContentMust(file)))
		}
	}
	mux.HandleFunc("/ws", h.Channel())
	log.Printf("Web MPC starting…")
	log.Fatal(http.ListenAndServe(*addr, mux))
}

// eof
