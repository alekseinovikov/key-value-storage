package main

import (
	"errors"
	"log"
	"net/http"
)

var ErrorNoSuchKey = errors.New("no such key")
var store = make(map[string]string)

func helloGoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello net/http!\n"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloGoHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func Put(key string, value string) error {
	store[key] = value
	return nil
}

func Get(key string) (string, error) {
	value, ok := store[key]
	if !ok {
		return "", ErrorNoSuchKey
	}

	return value, nil
}

func Delete(key string) error {
	delete(store, key)
	return nil
}
