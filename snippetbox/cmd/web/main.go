package main

import (
	"log"
	"net/http"
)

func main() {
	//Throughout this project so far we’ve been using the HandleFunc() method to register our
	//handler functions with the servemux. This is just some syntactic sugar that transforms a
	//function to a handler and registers it in one step, instead of having to do it manually. The
	//code above is functionality equivalent to this:
	//
	// mux.Handle("/", http.HandlerFunc(home))
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
