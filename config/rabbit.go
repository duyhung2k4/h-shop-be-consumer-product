package config

import (
	"github.com/rabbitmq/amqp091-go"
)

func connectRabbitMQ() error {
	var err error = nil
	rabbitConnection, err = amqp091.Dial(urlRabbitMq)
	if err != nil {
		return err
	}

	return nil
}
