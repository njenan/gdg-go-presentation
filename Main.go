package main

import (
	"net/http"
	"log"
	"encoding/json"
	"strconv"
	"math/rand"
)

var store map[int]map[string]interface{}

func init() {
	store = make(map[int]map[string]interface{})
}

func get(w http.ResponseWriter, r *http.Request) {
	stringId := r.URL.Path[1:]

	var id int
	var err error

	if id, err = strconv.Atoi(stringId); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	jsonObj := store[id]

	if jsonObj == nil {
		w.Write([]byte("Object does not exist"))
		return
	}

	var bytes []byte
	if bytes, err = json.Marshal(jsonObj); err != nil {
		w.Write([]byte("Could not marshal object"))
	}

	w.Write(bytes)
}

func post(w http.ResponseWriter, r *http.Request) {
	var jsonBlob map[string]interface{}

	if err := json.NewDecoder(r.Body).Decode(&jsonBlob); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	id := rand.Intn(1000000)
	jsonBlob["__id"] = id
	store[id] = jsonBlob

	if bytes, err := json.Marshal(jsonBlob); err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Write(bytes)
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "GET" {
			get(w, r)
		} else if r.Method == "POST" {
			post(w, r)
		} else {
			w.Write([]byte("Something went wrong"))
		}
	})

	log.Println("Server started at port 8080")

	http.ListenAndServe(":8080", nil)
}