package mpc

import (
	"fmt"
	"log"
	"path"
	"strings"

	"github.com/fhs/gompd/mpd"
)

// Client with host, port & password & mpc reference
type Client struct {
	addr        string // host:port
	host        string
	port        int
	password    string
	logger      *log.Logger
	mpc         *mpd.Client
	mpw         *mpd.Watcher
	Event       *chan string
	queueLength int
}

// New with host, port, password and in & out channels
func New(host string, port int, password string, logger *log.Logger) (*Client, error) {
	mpc := &Client{
		addr:        fmt.Sprintf("%s:%d", host, port),
		host:        host,
		port:        port,
		password:    password,
		logger:      logger,
		queueLength: -1,
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
		client.logger.Printf("connected to %s (%s)", client.addr, client.mpc.Version())
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

// Version returns the protocol version in use
func (client *Client) Version() *Message {
	return VersionMsg(client.mpc.Version())
}

// Stats for the MPD database
func (client *Client) Stats() error {
	attr, err := client.mpc.Stats()
	client.logger.Printf("Stats: %+v", attr)
	return err
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

// Prio sets prio for song in the queue at pos to prio
func (client *Client) Prio(prio int, pos int) error {
	return client.mpc.SetPriority(prio, pos, -1)
}

func (client *Client) status() *mpd.Attrs {
	// we get EOF here sometimes.  why?
	client.Ping()
	status, err := client.mpc.Status()
	if err != nil {
		client.logger.Panic(err) // FIXME: no panic
	}
	client.logger.Printf("status: %+v", status)
	return &status
}

func (client *Client) activeSong() *mpd.Attrs {
	client.Ping()
	attrs, err := client.mpc.CurrentSong()
	if err != nil {
		client.logger.Println("ActiveSong:", err)
	}
	return &attrs
}

func (client *Client) queue() *[]mpd.Attrs {
	client.Ping()
	attrs, err := client.mpc.PlaylistInfo(-1, -1)
	if err != nil {
		client.logger.Println("ActivePlaylist:", err)
	}
	client.queueLength = len(attrs)
	return &attrs
}

// Update prepares an update message for the client containing:
// status, activeSong and queue
func (client *Client) Update() *Message {
	status := client.status()
	activeSong := client.activeSong()
	queue := client.queue()
	return UpdateDataMsg(status, activeSong, queue)
}

// RemovePlaylistEntry nr
func (client *Client) RemovePlaylistEntry(nr int) error {
	client.Ping()
	return client.mpc.Delete(nr, -1)
}

// Consume mode to enable
func (client *Client) Consume(enable bool) error {
	return client.mpc.Consume(enable)
}

// Repeat mode to enable
func (client *Client) Repeat(enable bool) error {
	return client.mpc.Repeat(enable)
}

// Random mode to enable
func (client *Client) Random(target bool) error {
	return client.mpc.Random(target)
}

// Single mode to enable
func (client *Client) Single(target bool) error {
	return client.mpc.Single(target)
}

func escapeSearchToken(token string) string {
	// FIXME: workaround for broken gompd search (%) => fix it upstream when you
	// have an idead how to fix it :)
	return strings.Replace(token, "%", "%%%%", -1)
}

// Search for search string tokenized by space and searched in any
// FIXME: escape special characters.  e. g. % does not work. why?  MPD docu?
func (client *Client) Search(search string) *Message {
	var searchTokens []string
	for _, token := range strings.Split(search, " ") {
		if token != "" {
			if strings.HasPrefix(token, "t:") {
				searchTokens = append(searchTokens, "title")
				searchTokens = append(searchTokens, escapeSearchToken(token[2:]))
			} else if strings.HasPrefix(token, "a:") {
				searchTokens = append(searchTokens, "artist")
				searchTokens = append(searchTokens, escapeSearchToken(token[2:]))
			} else if strings.HasPrefix(token, "al:") {
				searchTokens = append(searchTokens, "album")
				searchTokens = append(searchTokens, escapeSearchToken(token[3:]))
			} else {
				searchTokens = append(searchTokens, "any")
				searchTokens = append(searchTokens, escapeSearchToken(token))
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
	return client.mpc.Add(strings.Replace(file, "%", "%%", -1))
}

// Clean the queue
func (client *Client) Clean() error {
	return client.mpc.Delete(0, client.queueLength)
}

// ListDirectory lists the contents of directory
func (client *Client) ListDirectory(directory string) *Message {
	attrs, err := client.mpc.ListInfo(directory)
	if err != nil {
		client.logger.Println("directory list error:", err)
		return nil
	}
	previousDirectory, _ := path.Split(directory)
	if len(previousDirectory) > 1 {
		previousDirectory = previousDirectory[:len(previousDirectory)-1]
	}
	hasPreviousDirectory := true
	if directory == "" {
		hasPreviousDirectory = false
	}
	return DirectoryListMsg(previousDirectory, hasPreviousDirectory, &attrs)
}

// eof
