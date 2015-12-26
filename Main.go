package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"math/rand"
	"strconv"
	"io/ioutil"
	"io"
	"strings"
	"os"
)

var (
	GET_HANDLER = "GET"
	POST_HANDLER = "POST"
	WORKING_DIR, _ = os.Getwd()
	JSON_DIR = WORKING_DIR + "/jsons/"
	store map[int]map[string]interface{} = make(map[int]map[string]interface{})
	requestHandlers map[string]handler = make(map[string]handler)
)

type handler func(w http.ResponseWriter, r *http.Request)

func handle(w http.ResponseWriter, r *http.Request) {
	handler := requestHandlers[r.Method]
	handler(w, r)
}

func readBodyToJson(body io.ReadCloser) (map[string]interface{}, error) {
	var out map[string]interface{}
	data, err := ioutil.ReadAll(body)

	if err == nil {
		err = json.Unmarshal(data, &out)

		if err == nil {
			id := rand.Int()
			out["__id"] = id

			store[id] = out
		}
	}

	return out, err
}

func writeJsonToDisk(jsonBlob map[string]interface{}) error {
	id := strconv.Itoa(jsonBlob["__id"].(int))
	file, err := os.Create(JSON_DIR + id)

	if err == nil {
		var jsonAsString []byte
		jsonAsString, err = json.Marshal(jsonBlob)

		if err == nil {
			_, err = file.Write(jsonAsString)
		}
	}

	return err
}

func getPacketByKey(key string) (map[string]interface{}, error) {
	var out map[string]interface{}
	filepath := JSON_DIR + key
	bytes, err := ioutil.ReadFile(filepath)

	if err == nil {
		err = json.Unmarshal(bytes, &out)
	}

	return out, err
}

func setupRoutes() {
	os.Mkdir(JSON_DIR, 0777)

	requestHandlers[GET_HANDLER] = func(w http.ResponseWriter, r *http.Request) {
		key := strings.TrimLeft(r.URL.Path, "/")
		var write string

		packet, err := getPacketByKey(key)

		if err == nil {
			var bytes []byte
			bytes, err = json.Marshal(packet)

			if err == nil {
				write = string(bytes)
			}
		}

		if err != nil {
			write = err.Error()
		}

		fmt.Fprintf(w, write)
	}

	requestHandlers[POST_HANDLER] = func(w http.ResponseWriter, r *http.Request) {
		var write string

		out, err := readBodyToJson(r.Body)

		if err == nil {
			err = writeJsonToDisk(out)

			if err == nil {
				var toret []byte
				toret, err = json.Marshal(out)
				write = string(toret)
			}
		}

		if err != nil {
			write = err.Error()
		}

		fmt.Fprintf(w, write)
	}
}

func main() {
	setupRoutes()
	http.HandleFunc("/", handle)

	port := "8080"
	fmt.Println("Webserver listening at port: " + port)
	http.ListenAndServe(":" + port, nil)
}
