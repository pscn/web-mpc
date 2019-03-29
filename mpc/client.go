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
// Return false if connection is broken and could not be recovered
func (client *Client) Ping() error {
	client.logger.Println("ping")
	err := client.mpc.Ping()
	if err != nil {
		client.reConnect()
		return client.mpc.Ping()
	}
	return nil
}

// Status returns mpd.Attrs
func (client *Client) Status() *mpd.Attrs {
	client.logger.Println("status")
	// we get EOF here sometimes.  why?
	if err := client.Ping(); err != nil {
		panic(err)
	}
	status, err := client.mpc.Status()
	if err != nil {
		panic(err)
	}
	return &status
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
		rc <- ErrorEvent(errors.Wrapf(err, "failed to start watcher for MPD events"))
	}
	defer mpw.Close()
	rc <- StringEvent("listening for MPD events")

	go func() { // event loop
		status := *client.Status()
		rc <- StringEvent(fmt.Sprintf("MPD status: %s", status))
		for subsystem := range mpw.Event {
			rc <- StringEvent(fmt.Sprintf("MPD subsystem: %s", subsystem))
			switch subsystem {
			case "update":
				status := *client.Status()
				client.logger.Printf("Status: %v\n", status)
				if _, ok := status["updating_db"]; !ok { // if present, it's still in progress
					rc <- StringEvent("database updating")
				}
			case "player":

			}
		}
	}()

	// error loop
	for err := range mpw.Error {
		// Seen so far:
		// mpd shutdown → write: broken pipe
		rc <- ErrorEvent(errors.Wrapf(err, "MPDClient error loop"))
	}
	rc <- StringEvent("client shutdown")
	return
}
