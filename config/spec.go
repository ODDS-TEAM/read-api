package config

import (
	"os"

	"github.com/joho/godotenv"
)

type (
	// Specification holds environment variable name.
	Specification struct {
		DBHost    string
		DBName    string
		DBBookCol string
		DBTagCol  string
		ImgPath   string
		APIPort   string
	}
)

// Spec retrieves the value of the environment variable named by the key.
func Spec() *Specification {
	godotenv.Load()

	s := Specification{
		DBHost:    os.Getenv("DB_HOST"),
		DBName:    os.Getenv("DB_NAME"),
		DBBookCol: os.Getenv("DB_BOOK_COL"),
		DBTagCol:  os.Getenv("DB_TAG_COL"),
		ImgPath:   os.Getenv("IMG_PATH"),
		APIPort:   os.Getenv("API_PORT"),
	}
	return &s
}
