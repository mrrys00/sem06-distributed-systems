package main

import (
	"fmt"
	"log"
	"os"
	"spaceprogram/config"
	"spaceprogram/utils"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	name := os.Args[1]

	connection, channel, err := config.SetupRabbitConnection(config.AgencyName)
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

	queueHuman, queueCargo, queueSatellite, queueAdminCopy, queueAdminAgencyListen, err := config.SetupAgencyQueues(channel, name)

	forever := make(chan bool)
	go func() {
		log.Printf("Human routine starts ...\n")
		for {
			now := time.Now()
			msg := fmt.Sprintf("[Agency %s] [no %v] [HUMAN]", name, now)

			err := utils.PublishMessage(channel, queueHuman.Name, msg)
			if err != nil {
				log.Printf("Cannot send message to %s queue\n", queueHuman.Name)
			}
			err = utils.PublishMessage(channel, queueAdminCopy.Name, msg)
			if err != nil {
				log.Printf("Cannot send message to %s queue\n", queueAdminCopy.Name)
			}

			utils.PrintQueueParams(queueHuman)

			time.Sleep(10 * time.Second)
		}
	}()
	go func() {
		log.Printf("Cargo routine starts ...\n")
		for {
			now := time.Now()
			msg := fmt.Sprintf("[Agency %s] [no %v] [CARGO]", name, now)

			err := utils.PublishMessage(channel, queueCargo.Name, msg)
			if err != nil {
				log.Printf("Cannot send message to %s queue\n", queueCargo.Name)
			}
			err = utils.PublishMessage(channel, queueAdminCopy.Name, msg)
			if err != nil {
				log.Printf("Cannot send message to %s queue\n", queueAdminCopy.Name)
			}

			utils.PrintQueueParams(queueCargo)

			time.Sleep(7 * time.Second)
		}
	}()
	go func() {
		log.Printf("Satellite routine starts ...\n")
		for {
			now := time.Now()
			msg := fmt.Sprintf("[Agency %s] [no %v] [SATELLITE]", name, now)

			err := utils.PublishMessage(channel, queueSatellite.Name, msg)
			if err != nil {
				log.Printf("Cannot send message to %s queue\n", queueSatellite.Name)
			}
			err = utils.PublishMessage(channel, queueAdminCopy.Name, msg)
			if err != nil {
				log.Printf("Cannot send message to %s queue\n", queueAdminCopy.Name)
			}

			utils.PrintQueueParams(queueSatellite)

			time.Sleep(3 * time.Second)
		}
	}()

	if queueAdminAgencyListen != nil {
		log.Printf("Admin listening routine starts ...\n")
		go func() {
			for msg := range queueAdminAgencyListen {
				log.Printf("Admin message: %s\n", msg.Body)
			}
		}()
	}

	<-forever
}
