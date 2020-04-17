package db

import (
	"database/sql"
	"fmt"

	"github.com/ashwinspg/explore-golang/config"
	"github.com/ashwinspg/explore-golang/utils"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var db *sql.DB

func init() {
	var dbConnErr error
	l := utils.LogEntryWithRef()
	host := config.DB_HOST
	user := config.DB_USER
	password := config.DB_PASSWORD
	dbname := config.DB_NAME
	desc := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)

	db, dbConnErr = sql.Open("postgres", desc)
	if dbConnErr != nil {
		l.WithError(dbConnErr).Fatal("Failed to get DB connection")
	}

	db.SetMaxIdleConns(config.MAX_DB_CONNECTIONS)
	db.SetMaxOpenConns(config.MAX_DB_CONNECTIONS)
	logrus.Info("Successfully established database connection")
}

//GetDB - to get DB connection
func GetDB() *sql.DB {
	return db
}
