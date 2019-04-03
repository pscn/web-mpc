package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gobuffalo/packr/v2"
	"github.com/gorilla/websocket"
	"github.com/pscn/web-mpc/web"
)

var addr = flag.String("addr", "127.0.0.1:8080", "http service address")
var mpdHost = flag.String("mpd", "127.0.0.1:6000", "MPD service address")
var pass = flag.String("password", "", "MPD password")
var devel = flag.Bool("devel", false, "serves html, jss & css from the src templates directory")

var upgrader = websocket.Upgrader{} // FIXME: what is this, what does it do?

var box = packr.New("templates", "./templates")
var tmplName = []string{"index.html", "script.js", "style.css"}
var tmplType = []string{"text/html", "text/javascript", "text/css"}
var verbosity = 2

func main() {
	flag.Parse()
	log.SetFlags(0)
	logger := log.New(os.Stdout, "web-mpc ", log.LstdFlags|log.Lshortfile)
	// disable origin check to test from static html, css & js (FIXME: remove this)
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	h := web.New(&upgrader, verbosity)
	mux := http.NewServeMux()
	// read templates and add listener
	for i := range tmplName {
		if verbosity > 5 {
			logger.Printf("preparing handler for: %s", tmplName[i])
		}
		if verbosity > 5 {
			logger.Printf("adding handler: %s", tmplName[i])
		}
		if tmplName[i] == "index.html" {
			if *devel {
				mux.HandleFunc("/", h.StaticTemplateFile(tmplType[i], tmplName[i]))
			} else {
				tmplStr, err := box.FindString(tmplName[i])
				if err != nil {
					logger.Panicf("Failed to load template '%s': %v", tmplName[i], err)
				}
				mux.HandleFunc("/", h.StaticTemplateString(tmplType[i], tmplStr))
			}
		} else {
			if *devel {
				mux.HandleFunc(fmt.Sprintf("/%s", tmplName[i]),
					h.StaticFile(tmplType[i], tmplName[i]))
			} else {
				tmplByte, err := box.Find(tmplName[i])
				if err != nil {
					logger.Panicf("Failed to load template '%s': %v", tmplName[i], err)
				}
				mux.HandleFunc(fmt.Sprintf("/%s", tmplName[i]),
					h.StaticString(tmplType[i], string(tmplByte)))
			}
		}
	}
	mux.HandleFunc("/echo", h.Channel(*mpdHost, *pass))

	log.Fatal(http.ListenAndServe(*addr, mux))
}
