package handlers

import (
	"fmt"
	"net/http"
)

//PingHandler - to ping the server
func PingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server is Running...")
}
