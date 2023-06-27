package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"spaceprogram/config"
	"spaceprogram/utils"
	"time"
)

func main() {
	connection, channel, err := config.SetupRabbitConnection(config.AdminName)
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

	queueSpedit, queueAgency, queueAdminCopy, err := config.SetupAdminQueues(channel)
	if err != nil {
		log.Panicf("%+v\n", err)
	}

	forever := make(chan bool)
	go func() {
		log.Printf("Speditor routine starts ...\n")
		for {
			msg := fmt.Sprintf("[Admin to speditor]")

			err := utils.PublishMessage(channel, queueSpedit.Name, msg)
			if err != nil {
				log.Printf("Cannot send message to %s queue\n", queueSpedit.Name)
			}
			err = utils.PublishMessage(channel, queueSpedit.Name, msg)
			if err != nil {
				log.Printf("Cannot send message to %s queue\n", queueSpedit.Name)
			}
			utils.PrintQueueParams(queueSpedit)

			time.Sleep(5 * time.Second)
		}
	}()
	go func() {
		log.Printf("Agency routine starts ...\n")
		for {
			msg := fmt.Sprintf("[Admin to agency]")

			err := utils.PublishMessage(channel, queueAgency.Name, msg)
			if err != nil {
				log.Printf("Cannot send message to %s queue\n", queueAgency.Name)
			}
			err = utils.PublishMessage(channel, queueAgency.Name, msg)
			if err != nil {
				log.Printf("Cannot send message to %s queue\n", queueAgency.Name)
			}
			utils.PrintQueueParams(queueAgency)

			time.Sleep(3 * time.Second)
		}
	}()
	go func() {
		log.Printf("Agency and speditor routine starts ...\n")
		for {
			msg := fmt.Sprintf("[Admin to agency and speditor]")

			err := utils.PublishMessage(channel, queueAgency.Name, msg)
			if err != nil {
				log.Printf("Cannot send message to %s queue\n", queueAgency.Name)
			}
			err = utils.PublishMessage(channel, queueAgency.Name, msg)
			if err != nil {
				log.Printf("Cannot send message to %s queue\n", queueAgency.Name)
			}
			err = utils.PublishMessage(channel, queueSpedit.Name, msg)
			if err != nil {
				log.Printf("Cannot send message to %s queue\n", queueSpedit.Name)
			}
			err = utils.PublishMessage(channel, queueSpedit.Name, msg)
			if err != nil {
				log.Printf("Cannot send message to %s queue\n", queueSpedit.Name)
			}
			utils.PrintQueueParams(queueSpedit)
			utils.PrintQueueParams(queueAgency)

			time.Sleep(3 * time.Second)
		}
	}()

	if queueAdminCopy != nil {
		log.Printf("Copied messages routine starts ...\n")
		go func() {
			for msg := range queueAdminCopy {
				log.Printf("[COPY] message: %s\n", msg.Body)
			}
		}()
	}

	fmt.Println("Waiting for messages...")
	<-forever
}
