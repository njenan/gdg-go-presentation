package main

import (
	"net/http"
	"io/ioutil"
	"log"
	"io"
	"errors"
)

func readBody(body io.ReadCloser) ([]byte, error) {
	out, err := ioutil.ReadAll(body)

	if err != nil {
		return nil, err
	}

	if len(out) == 0 {
		return nil, errors.New("No body to echo!")
	}

	return out, nil
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		out, err := readBody(r.Body)

		if err != nil {
			log.Println(err.Error())
			w.Write([]byte(err.Error()))
		}

		w.Write([]byte(out))
	})

	log.Println("Server started at port 8080")

	http.ListenAndServe(":8080", nil)
}