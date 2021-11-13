package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func readBodyOnce(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, "read body failed: %v", err)
		return
	}

	fmt.Fprintf(w, "read the data: %s \n", string(body))

	body, err = io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, "ddddd")
		return
	}
	fmt.Fprintf(w, "read the data one more time: [%s]", string(body))
}

func queryParams(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	fmt.Fprintf(w, "query is %v\n", values)
}

func wholeUrl(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(r.URL)
	fmt.Fprintf(w, string(data))

}
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/bodyonce", readBodyOnce)
	http.HandleFunc("/queryparams", queryParams)
	http.HandleFunc("/fullurl", wholeUrl)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
