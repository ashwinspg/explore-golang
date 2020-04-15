package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func setPingRoutes(router chi.Router) {
	router.Get("/ping", PingHandler)
}

//PingHandler - to ping the server
func PingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Server is Running...")
}
