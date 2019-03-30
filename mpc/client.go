package mpc

import (
	"fmt"
	"log"

	"github.com/fhs/gompd/mpd"
	"github.com/pkg/errors"
)

// Client with host, port & password & mpc reference
type Client struct {
	host     string
	port     int
	password string
	logger   *log.Logger
	mpc      *mpd.Client
}

// New with host, port, password and in & out channels
func New(host string, port int, password string, logger *log.Logger) (*Client, error) {
	mpc := &Client{
		host:     host,
		port:     port,
		password: password,
		logger:   logger,
	}
	return mpc, mpc.reConnect()
}

func (client *Client) reConnect() (err error) {
	if client.password != "" {
		client.logger.Printf("connecting to %s:%d with %s", client.host, client.port, client.password)
		client.mpc, err = mpd.DialAuthenticated("tcp", fmt.Sprintf("%s:%d", client.host, client.port), client.password)
	} else {
		client.logger.Printf("connecting to %s:%d", client.host, client.port)
		client.mpc, err = mpd.Dial("tcp", fmt.Sprintf("%s:%d", client.host, client.port))
	}
	if err == nil {
		client.logger.Printf("connected to %s:%d", client.host, client.port)
	}
	return
}

// Close the MPDClient
func (client *Client) Close() error {
	client.logger.Println("closing connection")
	return client.mpc.Close() // close client
}

// Ping and try to re-connect if ping fails
// Panics if we can not re-connect FIXME: don't panic
func (client *Client) Ping() (err error) {
	if err = client.mpc.Ping(); err != nil {
		if err = client.reConnect(); err != nil {
			panic(err) // FIXME: no panic
		}
		if err = client.mpc.Ping(); err != nil {
			panic(err) // FIXME: no panic
		}
	}
	return
}

// Status returns mpd.Attrs
func (client *Client) Status() *mpd.Attrs {
	// we get EOF here sometimes.  why?
	client.Ping()
	status, err := client.mpc.Status()
	if err != nil {
		panic(err) // FIXME: no panic
	}
	return &status
}

// Play start playing
func (client *Client) Play() error {
	client.Ping()
	return client.mpc.Play(-1)
}

// Pause playing
func (client *Client) Pause() error {
	client.Ping()
	return client.mpc.Pause(true)
}

// Resume playing
func (client *Client) Resume() error {
	client.Ping()
	return client.mpc.Pause(false)
}

// Stop stops playing
func (client *Client) Stop() error {
	client.Ping()
	return client.mpc.Stop()
}

// Next song in playlist
func (client *Client) Next() error {
	client.Ping()
	return client.mpc.Next()
}

// Previous song in playlist
func (client *Client) Previous() error {
	client.Ping()
	return client.mpc.Previous()
}

// func (client *Client) Cover() []byte {
// return client.mpc.
// }

// CurrentSong returns the currently active song
func (client *Client) CurrentSong() *mpd.Attrs {
	client.Ping()
	attrs, err := client.mpc.CurrentSong()
	if err != nil {
		client.logger.Println("currentsong:", err)
	}
	return &attrs
}

// CurrentPlaylist returns the currently active playlist / queue
func (client *Client) CurrentPlaylist() *[]mpd.Attrs {
	client.Ping()
	attrs, err := client.mpc.PlaylistInfo(-1, -1)
	if err != nil {
		client.logger.Println("currentsong:", err)
	}
	return &attrs
}

// Search for str (as tokens)
func (client *Client) Search() *[]mpd.Attrs {
	client.Ping()
	attrs, err := client.mpc.PlaylistInfo(-1, -1)
	if err != nil {
		client.logger.Println("currentsong:", err)
	}
	return &attrs
}

// RemovePlaylistEntry nr
func (client *Client) RemovePlaylistEntry(nr int64) error {
	client.Ping()
	return client.mpc.Delete(int(nr), -1)
	//PlaylistDelete("", int(nr))
}

func (client *Client) watcher() (*mpd.Watcher, error) {
	mpw, err := mpd.NewWatcher("tcp", fmt.Sprintf("%s:%d", client.host, client.port), client.password, "")
	return mpw, err
}

// EventLoop with a return channel for messages
func (client *Client) EventLoop(rc chan *Event) {
	defer close(rc)
	mpw, err := mpd.NewWatcher("tcp", fmt.Sprintf("%s:%d", client.host, client.port), client.password, "")
	if err != nil { // FIXME: error recovery
		client.logger.Println("error", err)
		rc <- NewErrorEvent(errors.Wrapf(err, "failed to start watcher for MPD events"))
	}
	defer mpw.Close()
	rc <- NewStringEvent("listening for MPD events")

	go func() { // event loop
		rc <- NewStatusEvent(client.Status())
		rc <- NewCurrentSongEvent(client.CurrentSong())
		rc <- NewCurrentPlaylistEvent(client.CurrentPlaylist())
		for subsystem := range mpw.Event {
			rc <- NewStringEvent(fmt.Sprintf("MPD subsystem: %s", subsystem))
			switch subsystem {
			case "update":
				status := *client.Status()
				client.logger.Printf("Status: %v\n", status)
				if _, ok := status["updating_db"]; !ok { // if present, it's still in progress
					rc <- NewStringEvent("database updating")
				}
			case "player":
				status := *client.Status()
				rc <- NewStatusEvent(client.Status())
				rc <- NewCurrentSongEvent(client.CurrentSong())
				client.logger.Printf("Status: %v\n", status)
			case "playlist":
				rc <- NewCurrentPlaylistEvent(client.CurrentPlaylist())
			}
		}
	}()

	// error loop
	for err := range mpw.Error {
		// Seen so far:
		// mpd shutdown → write: broken pipe
		rc <- NewErrorEvent(errors.Wrapf(err, "MPDClient error loop"))
	}
	rc <- NewStringEvent("client shutdown")
	return
}
