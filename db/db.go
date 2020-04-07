package db

import (
	"database/sql"
	"fmt"

	"github.com/ashwinspg/explore-golang/config"
	_ "github.com/lib/pq"
)

//GetPostgresDB - getting postgres connection
func GetPostgresDB() (*sql.DB, error) {
	host := config.POSTGRES_HOST
	user := config.POSTGRES_USER
	password := config.POSTGRES_PASSWORD
	dbname := config.POSTGRES_DB_NAME

	desc := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)

	db, err := createConnection(desc)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func createConnection(desc string) (*sql.DB, error) {
	db, err := sql.Open("postgres", desc)

	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	return db, nil
}
