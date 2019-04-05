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
	host     string
	port     int
	password string
	logger   *log.Logger
	mpc      *mpd.Client
	mpw      *mpd.Watcher
	Event    *chan string
}

// New with host, port, password and in & out channels
func New(host string, port int, password string, logger *log.Logger) (*Client, error) {
	mpc := &Client{
		addr:     fmt.Sprintf("%s:%d", host, port),
		host:     host,
		port:     port,
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
			client.Event = &client.mpw.Event
		}
	}
	return
}

// Close the MPDClient
func (client *Client) Close() (err error) {
	client.logger.Println("closing connection")
	err = client.mpw.Close() // close client
	if err != nil {
		client.logger.Println("failed to close watcher:", err)
	}
	client.Event = nil
	return client.mpc.Close()
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

// Status returns mpd.Attrs
func (client *Client) Status() *Message {
	// we get EOF here sometimes.  why?
	client.Ping()
	status, err := client.mpc.Status()
	if err != nil {
		client.logger.Panic(err) // FIXME: no panic
	}
	return StatusMsg(&status)
}

// ActiveSong returns the currently active song
func (client *Client) ActiveSong() *Message {
	client.Ping()
	attrs, err := client.mpc.CurrentSong()
	if err != nil {
		client.logger.Println("ActiveSong:", err)
	}
	return ActiveSongMsg(&attrs)
}

// ActivePlaylist returns the currently active playlist / queue
func (client *Client) ActivePlaylist() *Message {
	client.Ping()
	attrs, err := client.mpc.PlaylistInfo(-1, -1)
	if err != nil {
		client.logger.Println("ActivePlaylist:", err)
	}
	return ActivePlaylistMsg(&attrs)
}

// RemovePlaylistEntry nr
func (client *Client) RemovePlaylistEntry(nr int) error {
	client.Ping()
	return client.mpc.Delete(nr, -1)
}

// Search for search string tokenized by space and searched in any
func (client *Client) Search(search string) *Message {
	var searchTokens []string
	for _, token := range strings.Split(search, " ") {
		if token != "" {
			if strings.HasPrefix(token, "t:") {
				searchTokens = append(searchTokens, "title")
				searchTokens = append(searchTokens, token[2:])
			} else if strings.HasPrefix(token, "a:") {
				searchTokens = append(searchTokens, "artist")
				searchTokens = append(searchTokens, token[2:])
			} else if strings.HasPrefix(token, "al:") {
				searchTokens = append(searchTokens, "album")
				searchTokens = append(searchTokens, token[3:])
			} else {
				searchTokens = append(searchTokens, "any")
				searchTokens = append(searchTokens, token)
			}
		}
	}
	client.logger.Printf("tokens: %v", searchTokens)
	if len(searchTokens) > 0 {
		attrs, err := client.mpc.Search(searchTokens...)
		if err != nil {
			client.logger.Println("search error:", err)
			return nil
		}
		return SearchResultMsg(&attrs)
	}
	return nil
}

// Add file to playlist
func (client *Client) Add(file string) error {
	return client.mpc.Add(file)
}

func (client *Client) ListDirectory(directory string) *Message {
	attrs, err := client.mpc.ListInfo(directory)
	if err != nil {
		client.logger.Println("directory list error:", err)
		return nil
	}
	return DirectoryListMsg(directory, &attrs)
}

// eof
