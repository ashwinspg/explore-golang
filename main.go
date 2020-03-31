package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"

	"github.com/ashwinspg/explore-golang/config"
	"github.com/ashwinspg/explore-golang/handlers"
	"github.com/ashwinspg/explore-golang/middlewares"
)

func main() {
	router := httprouter.New()

	router.GET("/ping", middlewares.LoggerMiddleware(handlers.PingHandler))

	logrus.Infof("Listening on port %s", config.PORT)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.PORT), router); err != nil {
		logrus.Panic(err)
	}
}
