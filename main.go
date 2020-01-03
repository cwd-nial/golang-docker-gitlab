package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Go Ducks!"))
	})

	log.Println("Starting HTTP services at 8080")
	err := http.ListenAndServe(":8080", r) // Goroutine will block here

	if err != nil {
		log.Println("An error occurred starting HTTP listener at port 8080")
		log.Println("Error: " + err.Error())
	}
}
