package config

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-chi/jwtauth/v5"
	"github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

const (
	APP_PORT    = "APP_PORT"
	DB_HOST     = "DB_HOST"
	DB_PORT     = "DB_PORT"
	DB_NAME     = "DB_NAME"
	DB_PASSWORD = "DB_PASSWORD"
	DB_USER     = "DB_USER"
	URL_REDIS   = "URL_REDIS"
	HOST        = "HOST"

	URL_RABBIT_MQ = "URL_RABBIT_MQ"
	URL_ELASTIC   = "URL_ELASTIC"
)

var (
	appPort     string
	urlRedis    string
	host        string
	urlRabbitMq string
	urlElastic  string

	rdb *redis.Client
	jwt *jwtauth.JWTAuth

	clientProduct    *grpc.ClientConn
	rabbitConnection *amqp091.Connection
	elasticClient    *elasticsearch.TypedClient
)
