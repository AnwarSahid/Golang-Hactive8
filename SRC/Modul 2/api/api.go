package main

import (
	"fmt"
	"net/http"
)

var port = ":8080"

func main() {
	// http.HandleFunc("/", greet)
	// http.HandleFunc("/cek", halo)

	// http.ListenAndServe(port, nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	msg := "halo dunia"
	fmt.Fprint(w, msg)
}
