// package main

// import (
// 	"log"
// 	"net/http"
// )

// func main() {
// 	room := &room{}
// 	http.Handle("/", room)

// 	go room.run()

// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		log.Fatal("ListenAndServe:", err.Error())
// 	}
// }

// func NewRoom() *room {
// 	return &room{
// 		forward: make(chan []byte),
// 		join:    make(chan *client),
// 		leave:   make(chan *client),
// 	}
// }

package main

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

type templateHandler struct {
	filename string
	templ    *template.Template
	once     sync.Once
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, r)
}

func main() {
	room := NewRoom()
	http.Handle("/chat", room)
	http.Handle("/", &templateHandler{filename: "chat.html"})

	go room.run()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}
