package mpc

import (
	"fmt"
	"log"
	"path"
	"strings"

	"github.com/fhs/gompd/mpd"
	"github.com/pscn/web-mpc/msg"
)

// Client with host, port & password & mpc reference
type Client struct {
	*mpd.Client
	*mpd.Watcher
	*log.Logger
	addr        string // host:port
	host        string
	port        int
	password    string
	queueLength int
	queuePage   int
}

// New with host, port, password and in & out channels
func New(host string, port int, password string, logger *log.Logger) (*Client, error) {
	mpc := &Client{
		addr:        fmt.Sprintf("%s:%d", host, port),
		host:        host,
		port:        port,
		password:    password,
		Logger:      logger,
		queueLength: -1,
		queuePage:   1,
	}
	return mpc, mpc.Connect()
}

// Connect the client
func (client *Client) Connect() (err error) {
	if client.password != "" {
		client.Printf("connecting to %s with %s", client.addr, client.password)
		client.Client, err = mpd.DialAuthenticated("tcp", client.addr, client.password)
	} else {
		client.Printf("connecting to %s", client.addr)
		client.Client, err = mpd.Dial("tcp", client.addr)
	}
	if err == nil {
		client.Printf("connected to %s (%s)", client.addr, client.Version())
		client.Watcher, err = mpd.NewWatcher("tcp", client.addr, client.password, "")
		if err == nil {
			client.Printf("listening to %s", client.addr)
		}
	}
	return
}

// Close the MPDClient
func (client *Client) Close() (err error) {
	client.Println("closing connection")
	err = client.Client.Close() // close client
	if err != nil {
		client.Println("failed to close watcher:", err)
	}
	return client.Watcher.Close()
}

// Version returns the protocol version in use
func (client *Client) Version() *msg.Message {
	return msg.VersionMsg(client.Client.Version())
}

// Stats for the MPD database
func (client *Client) Stats() error {
	attr, err := client.Client.Stats()
	client.Printf("Stats: %+v", attr)
	return err
}

// AddPrio sets prio for song in the queue at pos to prio
func (client *Client) AddPrio(prio int, file string) error {
	id, err := client.Client.AddID(strings.Replace(file, "%", "%%", -1), -1)
	if err != nil {
		client.Panic(err) // FIXME: no panic
	}
	return client.Client.SetPriorityID(prio, id)
}

// Update prepares an update message for the client containing:
// status, activeSong and queue
func (client *Client) Update(page int) *msg.Message {
	status, err := client.Status()
	if err != nil {
		client.Panic(err)
	}
	activeSong, err := client.CurrentSong()
	if err != nil {
		client.Panic(err)
	}
	queue, err := client.PlaylistInfo(-1, -1)
	if err != nil {
		client.Panic(err)
	}
	client.queueLength = len(queue)
	return msg.NewUpdate(&status, &activeSong, &queue, page, 10)
}

func escapeSearchToken(token string) string {
	// FIXME: workaround for broken gompd search (%) => fix it upstream when you
	// have an idead how to fix it :)
	return strings.Replace(token, "%", "%%%%", -1)
}

// Search for search string tokenized by space and searched in any
// FIXME: escape special characters.  e. g. % does not work. why?  MPD docu?
func (client *Client) Search(search string, page int) *msg.Message {
	var searchTokens []string
	for _, token := range strings.Split(search, " ") {
		if token != "" {
			if strings.HasPrefix(strings.ToLower(token), "t:") {
				searchTokens = append(searchTokens, "title")
				searchTokens = append(searchTokens, escapeSearchToken(token[2:]))
			} else if strings.HasPrefix(strings.ToLower(token), "a:") {
				searchTokens = append(searchTokens, "artist")
				searchTokens = append(searchTokens, escapeSearchToken(token[2:]))
			} else if strings.HasPrefix(strings.ToLower(token), "al:") {
				searchTokens = append(searchTokens, "album")
				searchTokens = append(searchTokens, escapeSearchToken(token[3:]))
			} else {
				searchTokens = append(searchTokens, "any")
				searchTokens = append(searchTokens, escapeSearchToken(token))
			}
		}
	}
	client.Printf("tokens: %v", searchTokens)
	if len(searchTokens) > 0 {
		attrs, err := client.Client.Search(searchTokens...)
		if err != nil { // this can happen if we would get too many results
			client.Println("search error:", err)
			return msg.NewError("Hrhr nice try... (Stephan mach kein Scheiß!)")
		}
		return msg.SearchResultMsg(&attrs, page, 10)
	}
	return nil
}

// Add file to playlist
func (client *Client) Add(file string) error {
	return client.Client.Add(strings.Replace(file, "%", "%%", -1))
}

// Clean the queue
func (client *Client) Clean() error {
	return client.Client.Delete(0, client.queueLength)
}

// ListDirectory lists the contents of directory
func (client *Client) ListDirectory(directory string) *msg.Message {
	attrs, err := client.ListInfo(directory)
	if err != nil {
		client.Println("directory list error:", err)
		return nil
	}
	previous, _ := path.Split(directory)
	if len(previous) > 1 {
		previous = previous[:len(previous)-1]
	}
	hasPrevious := true
	if directory == "" {
		hasPrevious = false
	}
	return msg.DirectoryList(directory, previous, hasPrevious, &attrs)
}

// ListPlaylists lists all playlists
func (client *Client) ListPlaylists(page int) *msg.Message {
	attrs, err := client.Client.ListPlaylists()
	if err != nil {
		client.Println("playlist list error:", err)
		return nil
	}
	return msg.PlaylistListMsg(&attrs, page, 10)
}

// eof
