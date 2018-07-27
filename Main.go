package main

import (
	"log"
	"net/http"
)

func main() {

	router := NewRouter()
	Init()
	log.Fatal(http.ListenAndServe(":8080", router))

	/*http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hello/goodmorning", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Good morning, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))*/

}
