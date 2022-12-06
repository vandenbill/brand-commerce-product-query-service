package util

import "github.com/rabbitmq/amqp091-go"

func RegisterConsumer(ch *amqp091.Channel, q amqp091.Queue) <-chan amqp091.Delivery {
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	FailOnError(err, "Failed to register a consumer")

	return msgs
}
