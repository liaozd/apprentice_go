package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	// Check if the current request URL path exactly matches "/". If it doesn't, use
	// the http.NotFound() function to send a 404 response to the client.
	// Importantly, we then return from the handler. If we don't return the handler
	// would keep executing and also write the "Hello from SnippetBox" message.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id parameter from the query string and try to
	// convert it to an integer using the strconv.Atoi() function. If it can't
	// be converted to an integer, or the value is less than 1, we return a 404 page
	// not found response.
	// curl http://localhost:4000/snippet?id=123
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
	w.Write([]byte("Display a specific snippet..."))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	// curl -i -X POST http://localhost:4000/snippet/create
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		// If it's not, use the w.WriteHeader() method to send a 405 status
		// code and the w.Write() method to write a "Method Not Allowed"
		// response body. We then return from the function so that the
		// subsequent code is not executed.
		// curl -i -X PUT http://localhost:4000/snippet/create

		// Use the http.Error() function to send a 405 status code and "Method Not
		// Allowed" string as the response body.
		// instead of w.WriteHeader(405)
		http.Error(w, "Method Not Allowed", 405)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
