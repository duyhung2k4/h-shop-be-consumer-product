package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return err
	}

	appPort = os.Getenv(APP_PORT)
	urlRedis = os.Getenv(URL_REDIS)
	host = os.Getenv(HOST)
	urlRabbitMq = os.Getenv(URL_RABBIT_MQ)
	urlElastic = os.Getenv(URL_ELASTIC)

	return nil
}
