package config

import (
	"os"
)

var (
	PORT            string
	MOVIEBUFF_TOKEN string
	MOVIEBUFF_URL   string
)

func init() {
	PORT = os.Getenv("PORT")
	MOVIEBUFF_TOKEN = os.Getenv("MOVIEBUFF_TOKEN")
	MOVIEBUFF_URL = os.Getenv("MOVIEBUFF_URL")
}
