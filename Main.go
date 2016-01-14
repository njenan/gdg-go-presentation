package main

import (
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Start handle request")

		w.Write([]byte("Hello world"))

		log.Println("End handle request")
	})

	log.Println("Server started at port 8080")

	http.ListenAndServe(":8080", nil)
}