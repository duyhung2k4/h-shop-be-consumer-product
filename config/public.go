package config

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-chi/jwtauth/v5"
	"github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

func GetRDB() *redis.Client {
	return rdb
}

func GetAppPort() string {
	return appPort
}

func GetJWT() *jwtauth.JWTAuth {
	return jwt
}

func GetConnProductGRPC() *grpc.ClientConn {
	return clientProduct
}

func GetHost() string {
	return host
}

func GetRabbitConnection() *amqp091.Connection {
	return rabbitConnection
}

func GetElasticClient() *elasticsearch.TypedClient {
	return elasticClient
}
