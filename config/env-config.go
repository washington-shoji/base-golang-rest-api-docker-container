package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func EnvConfig(key string) string {
	path, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error in path %s", err)
		log.Fatal(err)
	}
	if err := godotenv.Load(filepath.Join(path, ".env")); err != nil {
		fmt.Printf("Error in connectiong %s", err)
		log.Fatal(err)
	}
	return os.Getenv(key)

}
