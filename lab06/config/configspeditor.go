package config

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func SetupSpeditorQueues(channel *amqp.Channel, name string) (<-chan amqp.Delivery, <-chan amqp.Delivery, <-chan amqp.Delivery, <-chan amqp.Delivery, error) {
	// speditor1 - listens on human, cargo, nil, adminSpedit
	// speditor2 - listens on nil, cargo, satellite, adminSpedit
	var consumerName string

	if name == "1" {
		consumerName = SpeditName1
	} else {
		consumerName = SpeditName2
	}

	queueHuman, err := SetupSingleConsumeQueue(
		channel,
		QueueNameHuman,
		fmt.Sprintf("%s_%s", consumerName, QueueNameHuman),
	)
	if err != nil {
		log.Printf("Not possible to setup %s queue due to an error", QueueNameHuman)
		return nil, nil, nil, nil, err
	}

	queueCargo, err := SetupSingleConsumeQueue(
		channel,
		QueueNameCargo,
		fmt.Sprintf("%s_%s", consumerName, QueueNameCargo),
	)
	if err != nil {
		log.Printf("Not possible to setup %s queue due to an error", QueueNameCargo)
		return nil, nil, nil, nil, err
	}

	queueSatellite, err := SetupSingleConsumeQueue(
		channel,
		QueueNameSatel,
		fmt.Sprintf("%s_%s", consumerName, QueueNameSatel),
	)
	if err != nil {
		log.Printf("Not possible to setup %s queue due to an error", QueueNameSatel)
		return nil, nil, nil, nil, err
	}

	queueAdminListen, err := SetupSingleConsumeQueue(
		channel,
		QueueNameAdminSpedit,
		fmt.Sprintf("%s_%s", consumerName, QueueNameAdminSpedit),
	)
	if err != nil {
		log.Printf("Not possible to setup %s queue due to an error", QueueNameAdminSpedit)
		return nil, nil, nil, nil, err
	}

	if consumerName == SpeditName1 {
		return queueHuman, queueCargo, nil, queueAdminListen, nil
	}
	return nil, queueCargo, queueSatellite, queueAdminListen, nil
}
