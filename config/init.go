package config

import (
	"github.com/go-chi/jwtauth/v5"
)

func init() {

	loadEnv()
	jwt = jwtauth.New("HS256", []byte("h-shop"), nil)

	connectRedis()
	connectRabbitMQ()
	connectElastic()
}
