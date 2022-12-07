package util

import (
	"os"

	"github.com/rabbitmq/amqp091-go"
	amqp "github.com/rabbitmq/amqp091-go"
)

func DialRabbitMQ() *amqp091.Connection {
	rabbitmqURI := os.Getenv("RABBITMQ_URI")
	// rabbitmqURI := "amqp://root:root@localhost:5672/"
	conn, err := amqp.Dial(rabbitmqURI)
	FailOnError(err, "Failed to connect to RabbitMQ")

	return conn
}
