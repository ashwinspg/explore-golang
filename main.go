package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Server is running...")
	}

	http.HandleFunc("/ping", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
