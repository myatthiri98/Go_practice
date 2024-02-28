package main

import "github.com/gorilla/websocket"

type client struct {
	to     chan []byte
	room   *room
	socket websocket.Conn
}
