package main

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	http_router "github.com/julienschmidt/httprouter"
)

func main() {
	port := "8080"
	router := http_router.New()

	router.GET("/ping", loggerMiddleware(ping))

	log.Infof("Listening on port %s", port)
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Panic(err)
	}
}

func ping(w http.ResponseWriter, r *http.Request, ps http_router.Params) {
	fmt.Fprintf(w, "Server is Running...")
}

func loggerMiddleware(next http_router.Handle) http_router.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps http_router.Params) {
		log.Info(fmt.Sprintf("-REQ INFO- %s %s", r.URL.String(), r.Method))
		begin := time.Now().UnixNano()
		next(w, r, ps)
		end := time.Now().UnixNano()
		log.Info(fmt.Sprintf("-RES INFO- PROCESSED IN %dns", (end - begin)))
	}
}
