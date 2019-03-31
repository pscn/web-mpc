package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/gobuffalo/packr/v2"
	"github.com/gorilla/websocket"
	"github.com/pscn/web-mpc/mpc"
	"github.com/pscn/web-mpc/web"
)

var addr = flag.String("addr", "192.168.0.111:8080", "http service address")

var upgrader = websocket.Upgrader{} // FIXME: what is this, what does it do?

var box = packr.New("templates", "./templates")
var tmplName = []string{"index.html", "script.js", "style.css"}
var tmplType = []string{"text/html", "text/javascript", "text/css"}

func main() {
	flag.Parse()
	log.SetFlags(0)
	logger := log.New(os.Stdout, "web-mpc ", log.LstdFlags|log.Lshortfile)
	// disable origin check to test from static html, css & js (FIXME: remove this)
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	h := web.New(&upgrader, logger)
	mux := http.NewServeMux()
	// read templates and add listener
	p := map[string]interface{}{
		"error":           mpc.EventTypeError,
		"string":          mpc.EventTypeString,
		"status":          mpc.EventTypeStatus,
		"currentSong":     mpc.EventTypeCurrentSong,
		"currentPlaylist": mpc.EventTypeCurrentPlaylist,
	}
	for i := range tmplName {
		logger.Printf("reading: %s", tmplName[i])
		{
			var tmplNr = i
			logger.Printf("adding handler: %s", tmplName[tmplNr])
			if tmplName[i] == "index.html" {
				tmplStr, err := box.FindString(tmplName[i])
				if err != nil {
					logger.Panicf("Failed to load template '%s': %v", tmplName[i], err)
				}
				tmpl := template.Must(template.New("").Parse(tmplStr))
				mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
					logger.Printf("serving / to %s", r.Host)
					p["ws"] = "ws://" + r.Host + "/echo"
					w.Header().Set("Content-type", tmplType[tmplNr])
					tmpl.Execute(w, p)
				})
			} else {
				tmplByte, err := box.Find(tmplName[i])
				if err != nil {
					logger.Panicf("Failed to load template '%s': %v", tmplName[i], err)
				}
				mux.HandleFunc(fmt.Sprintf("/%s", tmplName[tmplNr]), func(w http.ResponseWriter, r *http.Request) {
					logger.Printf("serving /%s to %s", tmplName[tmplNr], r.Host)
					p["ws"] = "ws://" + r.Host + "/echo"
					w.Header().Set("Content-type", tmplType[tmplNr])
					w.Write(tmplByte)
				})

			}
		}
	}
	mux.HandleFunc("/echo", h.Channel)

	log.Fatal(http.ListenAndServe(*addr, mux))
}
