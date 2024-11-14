package main

import (
	"fmt"
	"net/http"
)

// write a hello handler function
func helloHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello from the server"))
	if err != nil {
		panic(err)
	}
}

// view handler
func viewHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello from the Viewer"))
	if err != nil {
		panic(err)
	}
}

// create handler
func createHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello from the Create"))
	if err != nil {
		panic(err)
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/{$}", helloHandler)
	mux.HandleFunc("/aze/view", viewHandler)
	mux.HandleFunc("/aze/create", createHandler)

	fmt.Println("Listening on localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	fmt.Println(err)
}
