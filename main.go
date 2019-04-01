package main

import (
	"flag"
	"fmt"
	"io/ioutil"
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
	p := map[string]interface{}{
		"error":           mpc.Error,
		"string":          mpc.Info,
		"status":          mpc.Status,
		"currentSong":     mpc.CurrentSong,
		"currentPlaylist": mpc.Playlist,
	}
	for i := range tmplName {
		if verbosity > 5 {
			logger.Printf("preparing handler for: %s", tmplName[i])
		}
		{
			// FIXME: this is hackish, maybe use gorillas muxer or ...
			var tmplNr = i // copy to new scope so we can use it safelly in the callback functions
			if verbosity > 5 {
				logger.Printf("adding handler: %s", tmplName[tmplNr])
			}
			if tmplName[tmplNr] == "index.html" {
				if *devel {
					mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
						if verbosity > 5 {
							logger.Printf("serving / to %s", r.RemoteAddr)
						}
						p["ws"] = "ws://" + r.Host + "/echo"
						w.Header().Set("Content-type", tmplType[tmplNr])
						dat, err := ioutil.ReadFile(fmt.Sprintf("templates/%s", tmplName[tmplNr]))
						if err != nil {
							logger.Fatal(err)
						}
						tmpl := template.Must(template.New("").Parse(string(dat)))
						tmpl.Execute(w, p)
					})
				} else {
					tmplStr, err := box.FindString(tmplName[i])
					if err != nil {
						logger.Panicf("Failed to load template '%s': %v", tmplName[i], err)
					}
					tmpl := template.Must(template.New("").Parse(tmplStr))
					mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
						if verbosity > 5 {
							logger.Printf("serving / to %s", r.RemoteAddr)
						}
						p["ws"] = "ws://" + r.Host + "/echo"
						w.Header().Set("Content-type", tmplType[tmplNr])
						tmpl.Execute(w, p)
					})
				}
			} else {
				if *devel {
					mux.HandleFunc(fmt.Sprintf("/%s", tmplName[tmplNr]), func(w http.ResponseWriter, r *http.Request) {
						if verbosity > 5 {
							logger.Printf("serving /%s to %s", tmplName[tmplNr],
								r.RemoteAddr)
						}
						p["ws"] = "ws://" + r.Host + "/echo"
						w.Header().Set("Content-type", tmplType[tmplNr])
						dat, err := ioutil.ReadFile(fmt.Sprintf("templates/%s", tmplName[tmplNr]))
						if err != nil {
							logger.Fatal(err)
						}
						w.Write(dat)
					})
				} else {
					tmplByte, err := box.Find(tmplName[i])
					if err != nil {
						logger.Panicf("Failed to load template '%s': %v", tmplName[i], err)
					}
					mux.HandleFunc(fmt.Sprintf("/%s", tmplName[tmplNr]), func(w http.ResponseWriter, r *http.Request) {
						if verbosity > 5 {
							logger.Printf("serving /%s to %s", tmplName[tmplNr], r.RemoteAddr)
						}
						p["ws"] = "ws://" + r.Host + "/echo"
						w.Header().Set("Content-type", tmplType[tmplNr])
						w.Write(tmplByte)
					})
				}
			}
		}
	}
	mux.HandleFunc("/echo", h.Channel)

	log.Fatal(http.ListenAndServe(*addr, mux))
}
