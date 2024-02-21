package main

import (
	"log"
	"net/http"
)

func main() {

	// name := "Go"

	// var someName string = "Java"

	// fruits := []string{"apple","orange"}
	// fmt.Println("Hello World")

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println(r.Method)
	// })

	http.Handle("/", http.FileServer(http.Dir("templates")))

	// err := http.ListenAndServe(":8000", nil)
	// if err != nil {
	// 	log.Fatal("ListenAndServe:", err.Error())
	// }
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}
