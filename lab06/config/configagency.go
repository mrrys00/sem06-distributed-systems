package config

import "github.com/streadway/amqp"

func SetupAgencyQueues(channel *amqp.Channel, name string) (*amqp.Queue, *amqp.Queue, *amqp.Queue, *amqp.Queue, <-chan amqp.Delivery, error) {
	// agency1 - sends on human, cargo, satellite, adminCopy; listens on adminAgency
	// agency2 - sends on human, cargo, satellite, adminCopy; listens on adminAgency
	var speditorName string

	if name == "1" {
		speditorName = AgencyName1
	} else {
		speditorName = AgencyName2
	}

	queueHuman, err := SetupSingleQueue(
		channel,
		QueueNameHuman,
	)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	queueCargo, err := SetupSingleQueue(
		channel,
		QueueNameCargo,
	)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	queueSatellite, err := SetupSingleQueue(
		channel,
		QueueNameSatel,
	)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	queueAdminCopy, err := SetupSingleQueue(
		channel,
		QueueNameAdminCopy,
	)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	queueAdminAgencyListen, err := SetupSingleConsumeQueue(
		channel,
		QueueNameAdminAgency,
		speditorName,
	)

	return &queueHuman, &queueCargo, &queueSatellite, &queueAdminCopy, queueAdminAgencyListen, nil
}
