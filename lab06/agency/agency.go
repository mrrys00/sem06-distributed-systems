package main

import (
	"fmt"
	"log"
	"os"
	"spaceprogram/config"

	"github.com/streadway/amqp"
)

func main() {
	name := os.Args[1]

	connection, channel, err := config.SetupRabbitConnection()
	if err != nil {
		log.Panicf("%+v\n", err)
	}
	defer func(connection *amqp.Connection) {
		err := connection.Close()
		if err != nil {
			log.Panicf("%+v\n", err)
		}
	}(connection)
	defer func(channel *amqp.Channel) {
		err := channel.Close()
		if err != nil {
			log.Panicf("%+v\n", err)
		}
	}(channel)

	config.SetupAgencyQueues(name)

	// declaring queue with its properties over the channel opened
	queue, err := channel.QueueDeclare(
		"testing", // name
		false,     // durable
		false,     // auto delete
		false,     // exclusive
		false,     // no wait
		nil,       // args
	)
	if err != nil {
		panic(err)
	}

	// publishing a message
	err = channel.Publish(
		"",        // exchange
		"testing", // key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Test Message"),
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Queue status:", queue)
	fmt.Println("Successfully published message")

}
