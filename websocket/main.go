package main

import (
	"github.com/QGitter/gostudy/websocket/ws"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func main() {
	http.HandleFunc("/ws", wsHandler)

	http.ListenAndServe("0.0.0.0:7777", nil)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		wsConn *websocket.Conn
		err    error
		conn   *ws.Connection
		data   []byte
	)
	if wsConn, err = upgrader.Upgrade(w, r, nil); err != nil {
		return
	}

	if conn, err = ws.InitConnection(wsConn); err != nil {
		goto ERR
	}

	go func() {
		for {
			if e := conn.WriteMessage([]byte("heartbeat")); e != nil {
				return
			}
			time.Sleep(time.Second)
		}
	}()
	for {
		if data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}

		if err = conn.WriteMessage(data); err != nil {
			goto ERR
		}
	}
ERR:
	conn.Close()
}
