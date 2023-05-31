package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
	"spaceprogram/config"
)

func main() {
	name := os.Args[1]

	connection, channel, err := config.SetupRabbitConnection(config.SpeditorName)
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

	queueHuman, queueCargo, queueSatellite, queueAdminListen, err := config.SetupSpeditorQueues(channel, name)
	if err != nil {
		log.Panicf("%+v\n", err)
	}

	forever := make(chan bool)
	if queueHuman != nil {
		go func() {
			log.Printf("Human queue starts ...\n")
			for msg := range queueHuman {
				log.Printf("Human request: %s\n", msg.Body)
			}
		}()
	}
	if queueCargo != nil {
		go func() {
			log.Printf("Cargo queue starts ...\n")
			for msg := range queueCargo {
				log.Printf("Cargo request: %s\n", msg.Body)
			}
		}()
	}
	if queueSatellite != nil {
		go func() {
			log.Printf("Satelitte queue starts ...\n")
			for msg := range queueSatellite {
				log.Printf("Satellite request: %s\n", msg.Body)
			}
		}()
	}
	if queueAdminListen != nil {
		go func() {
			log.Printf("Admin queue starts ...\n")
			for msg := range queueAdminListen {
				log.Printf("Admin message: %s\n", msg.Body)
			}
		}()
	}

	fmt.Println("Waiting for messages...")
	<-forever
}
