package util

import "github.com/rabbitmq/amqp091-go"

func DeclareQueue(ch *amqp091.Channel) amqp091.Queue {
	q, err := ch.QueueDeclare(
		"product-queue", // name
		false,           // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	FailOnError(err, "Failed to declare a queue")

	return q
}
