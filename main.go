package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/karrick/golf"
	"github.com/pscn/web-mpc/server"
	"github.com/pscn/web-mpc/templates"
)

//go:generate npm run build
//go:generate file2go -v -t -o templates/files.go dist/*.html dist/*.ico dist/js/*.js dist/js/*.map dist/css/*.css

var addr = golf.StringP('a', "addr", ":8666", "http service address")
var mpdHost = golf.StringP('m', "mpd", "127.0.0.1:6600", "MPD service address")
var mpdPass = golf.StringP('p', "password", "", "MPD password")

var verbosity = 2

func main() {
	golf.Parse()
	log.SetFlags(0)
	h := server.New(verbosity, *mpdHost, *mpdPass)
	mux := http.NewServeMux()
	// read templates and add listener
	for _, file := range templates.Filenames() {
		f, err := filepath.Rel("dist", file)
		if err != nil {
			panic(err)
		}
		f = filepath.ToSlash(f)
		switch {
		case f == "index.html":
			fmt.Printf("%s\n", f)
			mux.HandleFunc("/", h.StaticPacked("text/html",
				templates.ContentMust(file)))
		case strings.HasSuffix(f, "html"):
			fmt.Printf("%s\n", f)
			mux.HandleFunc("/"+f, h.StaticPacked("text/html",
				templates.ContentMust(file)))
		case strings.HasSuffix(f, "js"):
			fmt.Printf("%s\n", f)
			mux.HandleFunc("/"+f, h.StaticPacked("text/javascript",
				templates.ContentMust(file)))
		case strings.HasSuffix(f, "map"):
			fmt.Printf("%s\n", f)
			mux.HandleFunc("/"+f, h.StaticPacked("application/octet-stream",
				templates.ContentMust(file)))
		case strings.HasSuffix(f, "css"):
			fmt.Printf("%s\n", f)
			mux.HandleFunc("/"+f, h.StaticPacked("text/css",
				templates.ContentMust(file)))
		case strings.HasSuffix(f, "ico"):
			fmt.Printf("%s\n", f)
			mux.HandleFunc("/"+f, h.StaticPacked("image/x-icon",
				templates.ContentMust(file)))
		}
	}
	mux.HandleFunc("/ws", h.Channel())
	log.Printf("Web MPC starting…")
	log.Fatal(http.ListenAndServe(*addr, mux))
}

// eof
