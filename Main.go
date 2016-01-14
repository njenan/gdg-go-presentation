package main

import (
	"net/http"
	"log"
	"encoding/json"
	"strconv"
	"math/rand"
	"errors"
)

var store map[int]map[string]interface{}

func init() {
	store = make(map[int]map[string]interface{})
}

type MethodHandler func(w http.ResponseWriter, r *http.Request) ([]byte, error)

func get(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	stringId := r.URL.Path[1:]

	if stringId == "" {
		return search(w, r)
	}

	var id int
	var err error

	if id, err = strconv.Atoi(stringId); err != nil {
		return nil, errors.New("Key was not a number")
	}

	jsonObj := store[id]

	if jsonObj == nil {
		return nil, errors.New("Document does not exist")
	}

	var bytes []byte
	if bytes, err = json.Marshal(jsonObj); err != nil {
		return nil, errors.New("Could not marshal object")
	}

	return bytes, nil
}

func search(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	size := len(store)
	allJsons := make([]map[string]interface{}, size)

	channel := make(chan map[string]interface{})

	for key, _ := range store {
		go func(key int, channel chan map[string]interface{}) {
			channel <- store[key]
		}(key, channel)
	}

	for i := 0; i < size; i ++ {
		allJsons = append(allJsons, <- channel)
	}

	if bytes, err := json.Marshal(allJsons); err != nil {
		return nil, errors.New("Could not marshal object")
	} else {
		return bytes, nil
	}
}

func post(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	var jsonBlob map[string]interface{}

	if err := json.NewDecoder(r.Body).Decode(&jsonBlob); err != nil {
		return nil, errors.New("Could not parse body as json")
	}

	id := rand.Intn(1000000)
	jsonBlob["__id"] = id
	store[id] = jsonBlob

	if bytes, err := json.Marshal(jsonBlob); err != nil {
		return nil, errors.New("Could not marshal json")
	} else {
		return bytes, nil
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var handler MethodHandler

		if r.Method == "GET" {
			handler = get
		} else if r.Method == "POST" {
			handler = post
		}

		w.Header().Set("Content-Type", "application/json")

		if bytes, err := handler(w, r); err != nil {
			w.Write([]byte(err.Error()))
		} else {
			w.Write(bytes)
		}
	})

	log.Println("Server started at port 8080")

	http.ListenAndServe(":8080", nil)
}