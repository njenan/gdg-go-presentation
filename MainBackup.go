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

type handler func(w http.ResponseWriter, r *http.Request) (map[string]interface{}, error)

func handle(w http.ResponseWriter, r *http.Request) {
	handler := requestHandlers[r.Method]
	jsonObj, err := handler(w, r)

	var toMarshal interface{}

	if err == nil {
		toMarshal = jsonObj
	} else {
		toMarshal = err
	}

	out, _ := json.Marshal(toMarshal)

	fmt.Fprintf(w, string(out));
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

	var file *os.File
	var err error

	if file, err = os.Create(JSON_DIR + id); err != nil {
		return err
	}

	var jsonAsString []byte

	if jsonAsString, err = json.Marshal(jsonBlob); err != nil {
		return err;
	}

	if _, err = file.Write(jsonAsString); err != nil {
		return err;
	}

	return nil
}

func getPacketByKey(key string) (map[string]interface{}, error) {
	var out map[string]interface{}
	var bytes []byte
	var err error
	filepath := JSON_DIR + key

	if bytes, err = ioutil.ReadFile(filepath); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(bytes, &out); err != nil {
		return nil, err;
	}

	return out, nil
}

func setupRoutes() {
	os.Mkdir(JSON_DIR, 0777)

	requestHandlers[GET_HANDLER] = func(w http.ResponseWriter, r *http.Request) (map[string]interface{}, error) {
		key := strings.TrimLeft(r.URL.Path, "/")
		packet, err := getPacketByKey(key)

		return packet, err
	}

	requestHandlers[POST_HANDLER] = func(w http.ResponseWriter, r *http.Request) (map[string]interface{}, error) {
		var out map[string]interface{}
		var err error

		if out, err = readBodyToJson(r.Body); err != nil {
			return nil, err
		}

		if err = writeJsonToDisk(out); err != nil {
			return nil, err
		}

		return out, nil
	}
}

func main() {
	setupRoutes()
	http.HandleFunc("/", handle)

	port := "8080"
	fmt.Println("Webserver listening at port: " + port)
	http.ListenAndServe(":" + port, nil)
}
