
# Web MPC

A small web client for [MPD](https://www.musicpd.org/) written in [go](https://golang.org/) using websockets.

## Install

```sh
go get github.com/pscn/web-mpc
```

## Run

If your MPD is running on localhost on the standard port:

```sh
web-mpc
```

And go to [http://localhost:8666](http://localhost:8666).

To listen on a different port use (e. g. 8080):

```sh
web-mpc -addr :8080
```

To connect to a different MPD use:

```sh
web-mpc -mpd 192.168.0.1:6000 -password changeme
```

#### #eof
