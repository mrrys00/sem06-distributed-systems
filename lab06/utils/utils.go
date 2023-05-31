package utils

import (
	"github.com/streadway/amqp"
	"log"
)

func PublishMessage(channel *amqp.Channel, name, body string) error {
	return channel.Publish(
		"",
		name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
}

func PrintQueueParams(queue *amqp.Queue) {
	log.Printf("Queue name %s\tconsumers %d\tmessages %d\n", queue.Name, queue.Consumers, queue.Messages)
	return
}
