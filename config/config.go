package config

import (
	"os"
)

var (
	PORT                 string
	MOVIEBUFF_TOKEN      string
	MOVIEBUFF_URL        string
	DATABASE_URL         string
	MIGRATION_FILES_PATH string
)

func init() {
	PORT = os.Getenv("PORT")
	MOVIEBUFF_TOKEN = os.Getenv("MOVIEBUFF_TOKEN")
	MOVIEBUFF_URL = os.Getenv("MOVIEBUFF_URL")
	DATABASE_URL = os.Getenv("DATABASE_URL")
	MIGRATION_FILES_PATH = os.Getenv("MIGRATION_FILES_PATH")
}
