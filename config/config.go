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

	DB_HOST            string
	DB_PORT            string
	DB_USER            string
	DB_PASSWORD        string
	DB_NAME            string
	DB_URL             string
	MAX_DB_CONNECTIONS int
)

func init() {
	PORT = os.Getenv("PORT")
	MOVIEBUFF_TOKEN = os.Getenv("MOVIEBUFF_TOKEN")
	MOVIEBUFF_URL = os.Getenv("MOVIEBUFF_URL")
	MIGRATION_FILES_PATH = os.Getenv("MIGRATION_FILES_PATH")

	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME = os.Getenv("DB_NAME")
	DB_URL = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	MAX_DB_CONNECTIONS, _ = strconv.Atoi(os.Getenv("MAX_DB_CONNECTIONS"))
}
