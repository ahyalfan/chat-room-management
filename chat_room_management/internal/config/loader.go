package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func NewConfig() *Config {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("env load failed", err.Error())
	}

	expInt, err := strconv.Atoi(os.Getenv("JWT_EXP"))
	if err != nil {
		log.Fatal("Error convert config env file: ", err.Error())
	}

	return &Config{
		Server: Server{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		Databases: Database{
			Host:     os.Getenv("DATABASE_HOST"),
			Port:     os.Getenv("DATABASE_PORT"),
			Username: os.Getenv("DATABASE_USER"),
			Password: os.Getenv("DATABASE_PASS"),
			Name:     os.Getenv("DATABASE_NAME"),
		},
		JWT: JWT{
			Key:     os.Getenv("JWT_KEY"),
			Expired: expInt,
		},
	}
}
