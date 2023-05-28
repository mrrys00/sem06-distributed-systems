package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"spaceprogram/config"
)

func main() {
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

	config.SetupAdminQueues()

	// publikacja admina do wszystkich po parametrze UserId z publish

	//declaring queue with its properties over the channel opened
	msgs, err := channel.Consume(
		config.QueueNameAdminListen,
		"admin",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	// print consumed messages from queue
	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			fmt.Printf("[ADMIN] Received Message: %s\n", msg.Body)
		}
	}()

	fmt.Println("Waiting for messages...")
	<-forever
}
