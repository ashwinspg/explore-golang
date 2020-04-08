package config

import (
	"fmt"
	"os"
	"strconv"
)

var (
	PORT                 string
	MOVIEBUFF_TOKEN      string
	MOVIEBUFF_URL        string
	MIGRATION_FILES_PATH string

	POSTGRES_HOST      string
	POSTGRES_PORT      string
	POSTGRES_USER      string
	POSTGRES_PASSWORD  string
	POSTGRES_DB_NAME   string
	DATABASE_URL       string
	MAX_DB_CONNECTIONS int
)

func init() {
	PORT = os.Getenv("PORT")
	MOVIEBUFF_TOKEN = os.Getenv("MOVIEBUFF_TOKEN")
	MOVIEBUFF_URL = os.Getenv("MOVIEBUFF_URL")
	MIGRATION_FILES_PATH = os.Getenv("MIGRATION_FILES_PATH")

	POSTGRES_HOST = os.Getenv("POSTGRES_HOST")
	POSTGRES_PORT = os.Getenv("POSTGRES_PORT")
	POSTGRES_USER = os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_DB_NAME = os.Getenv("POSTGRES_DB_NAME")
	DATABASE_URL = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_HOST, POSTGRES_PORT, POSTGRES_DB_NAME)
	MAX_DB_CONNECTIONS, _ = strconv.Atoi(os.Getenv("MAX_DB_CONNECTIONS"))
}
