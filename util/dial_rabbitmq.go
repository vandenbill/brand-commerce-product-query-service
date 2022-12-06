package util

import (
	"github.com/rabbitmq/amqp091-go"
	amqp "github.com/rabbitmq/amqp091-go"
	"os"
)

func DialRabbitMQ() *amqp091.Connection {
	rabbitmqURI := os.Getenv("RABBITMQ_URI")
	conn, err := amqp.Dial(rabbitmqURI)
	FailOnError(err, "Failed to connect to RabbitMQ")

	return conn
}
