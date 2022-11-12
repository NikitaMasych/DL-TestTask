package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var (
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDBName   string
)

func init() {
	loadEnv()
	setupVariables()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err == nil {
		return
	}
	// In case we test from inner directories:
	max_directory_deep := 2
	for i := 1; i != max_directory_deep; i++ {
		escape_sequence := strings.Repeat("../", i)
		err = godotenv.Load("./" + escape_sequence + ".env")
		if err == nil {
			return
		}
	}
	log.Fatal(err)
}

func setupVariables() {
	PostgresHost = os.Getenv("POSTGRES_HOST")
	PostgresPort = os.Getenv("POSTGRES_PORT")
	PostgresUser = os.Getenv("POSTGRES_USER")
	PostgresPassword = os.Getenv("POSTGRES_PASSWORD")
	PostgresDBName = os.Getenv("POSTGRES_DB_NAME")
}
