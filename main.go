package main

import (
	"log"
	"net/http"
)

func main() {
	room := &room{}
	http.Handle("/", room)

	go room.run()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}

func NewRoom() *room {
	return &room{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
	}
}
