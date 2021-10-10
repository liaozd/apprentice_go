package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("hello wolrd")
	})
	http.ListenAndServe("127.0.0.1:9090", nil)
}
