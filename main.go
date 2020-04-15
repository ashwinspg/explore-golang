package main

import (
	"fmt"
	"net/http"

	"github.com/ashwinspg/explore-golang/config"
	"github.com/ashwinspg/explore-golang/db"
	"github.com/ashwinspg/explore-golang/handlers"

	"github.com/sirupsen/logrus"
)

func main() {

	db.MigrateUp()
	db.InitiateConnection()

	logrus.Infof("Listening on port %s", config.PORT)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.PORT), handlers.GetRouter()); err != nil {
		logrus.Panic(err)
	}
}
