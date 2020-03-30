package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT string
)

func init() {
	godotenv.Load(".env")

	PORT = os.Getenv("PORT")
}
