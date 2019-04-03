package mpc

import (
	"fmt"
	"log"
	"strings"

	"github.com/fhs/gompd/mpd"
)

// Client with host, port & password & mpc reference
type Client struct {
	addr     string // host:port
	password string
	logger   *log.Logger
	mpc      *mpd.Client
	mpw      *mpd.Watcher
}

// New with host, port, password and in & out channels
func New(addr string, password string, logger *log.Logger) (*Client, error) {
	mpc := &Client{
		addr:     addr,
		password: password,
		logger:   logger,
	}
	return mpc, mpc.reConnect()
}

func (client *Client) reConnect() (err error) {
	if client.password != "" {
		client.logger.Printf("connecting to %s with %s", client.addr, client.password)
		client.mpc, err = mpd.DialAuthenticated("tcp", client.addr, client.password)
	} else {
		client.logger.Printf("connecting to %s", client.addr)
		client.mpc, err = mpd.Dial("tcp", client.addr)
	}
	if err == nil {
		client.logger.Printf("connected to %s", client.addr)
		client.mpw, err = mpd.NewWatcher("tcp", client.addr, client.password, "")
		if err == nil {
			client.logger.Printf("listening to %s", client.addr)
		}
	}
	return
}

// Close the MPDClient
func (client *Client) Close() (err error) {
	client.logger.Println("closing connection")
	err = client.mpc.Close() // close client
	if err != nil {
		client.logger.Printf("failed to close client: %v", err)
	}
	return client.mpw.Close()
}

// Ping and try to re-connect if ping fails
// Panics if we can not re-connect FIXME: don't panic
func (client *Client) Ping() (err error) {
	if err = client.mpc.Ping(); err != nil {
		if err = client.reConnect(); err != nil {
			client.logger.Panic(err) // FIXME: no panic
		}
		if err = client.mpc.Ping(); err != nil {
			client.logger.Panic(err) // FIXME: no panic
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
		client.logger.Panic(err) // FIXME: no panic
	}
	return &status
}

// Play start playing
func (client *Client) Play(nr int) error {
	client.Ping()
	return client.mpc.Play(nr)
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
// func (client *Client) Search() *[]mpd.Attrs {
// 	client.Ping()
// 	attrs, err := client.mpc.Find(-1, -1)
// 	if err != nil {
// 		client.logger.Println("currentsong:", err)
// 	}
// 	return &attrs
// }

// RemovePlaylistEntry nr
func (client *Client) RemovePlaylistEntry(nr int) error {
	client.Ping()
	return client.mpc.Delete(nr, -1)
}

// Search for search string tokenized by space and searched in any
func (client *Client) Search(search string) *[]mpd.Attrs {
	var searchTokens []string
	for _, token := range strings.Split(search, " ") {
		if token != "" {
			searchTokens = append(searchTokens, "any")
			searchTokens = append(searchTokens, token)
		}
	}
	client.logger.Printf("tokens: %v", searchTokens)
	if len(searchTokens) > 0 {
		attrs, err := client.mpc.Search(searchTokens...)
		if err != nil {
			client.logger.Printf("search error: %v", err)
			return nil
		}
		return &attrs
	}
	return nil
}

// Add file to playlist
func (client *Client) Add(file string) error {
	return client.mpc.Add(file)
}

// EventLoop with a return channel for messages
func (client *Client) EventLoop(rc chan *Message) {
	defer client.logger.Println("stop eventloop")

	go func() { // event loop
		defer func() { // if you want to recover from any panic below, use this
			if r := recover(); r != nil {
				client.logger.Println("recovered", r)
			}
		}()
		rc <- NewStatus(client.Status())
		rc <- NewCurrentSong(client.CurrentSong())
		rc <- NewCurrentPlaylist(client.CurrentPlaylist())
		for subsystem := range client.mpw.Event {
			client.logger.Printf("MPD subsystem: %s", subsystem)
			switch subsystem {
			case "update":
				status := *client.Status()
				client.logger.Printf("Status: %v\n", status)
				if _, ok := status["updating_db"]; !ok { // if present, it's still in progress
					rc <- NewInfo("database updating")
				}
			case "player", "playlist":
				status := *client.Status()
				currentSong := *client.CurrentSong()
				rc <- NewStatus(client.Status())
				rc <- NewCurrentSong(client.CurrentSong())
				rc <- NewCurrentPlaylist(client.CurrentPlaylist())
				client.logger.Printf("Status: %v\n", status)
				client.logger.Printf("CurrentSong: %v\n", currentSong)
			}
		}
		client.logger.Printf("mpw loop exited")
	}()

	// error loop
	for err := range client.mpw.Error {
		// Seen so far:
		// mpd shutdown → write: broken pipe
		rc <- NewError(fmt.Sprintf("MPD watcher error: %v", err))
		break
	}
	return
}
