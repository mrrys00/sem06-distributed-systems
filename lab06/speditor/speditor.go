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

	config.SetupSpeditorQueues(name)

	// declaring consumer with its properties over channel opened
	msgs, err := channel.Consume(
		"testing", // queue
		"",        // consumer
		true,      // auto ack
		false,     // exclusive
		false,     // no local
		false,     // no wait
		nil,       //args
	)
	if err != nil {
		panic(err)
	}

	// print consumed messages from queue
	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			fmt.Printf("Received Message: %s\n", msg.Body)
		}
	}()

	fmt.Println("Waiting for messages...")
	<-forever
}
