package main

import (
	"github.com/gorilla/websocket"
)

type client struct {
	socket  *websocket.Conn
	recieve chan []byte
	room    *room
}

