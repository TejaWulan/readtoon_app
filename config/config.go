package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	PGSQL_HOST     string
	API_VERSION    string
	PGSQL_PORT     string
	PGSQL_USER     string
	PGSQL_PASSWORD string
	PGSQL_DATABASE string
	SERVER_HOST    string
}

func LoadConfig() *Configuration {
	err := godotenv.Load("./config/config.env")
	if err != nil {
		log.Fatalf("Gagal baca file config.env: %v", err)
	}

	return &Configuration{
		PGSQL_HOST:     os.Getenv("PGSQL_HOST"),
		API_VERSION:    os.Getenv("API_VERSION"),
		PGSQL_PORT:     os.Getenv("PGSQL_PORT"),
		PGSQL_USER:     os.Getenv("PGSQL_USER"),
		PGSQL_PASSWORD: os.Getenv("PGSQL_PASSWORD"),
		PGSQL_DATABASE: os.Getenv("PGSQL_DATABASE"),
		SERVER_HOST:    os.Getenv("SERVER_HOST"),
	}
}
