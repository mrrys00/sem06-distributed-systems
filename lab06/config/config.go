package config

import (
	"fmt"
	"github.com/streadway/amqp"
)

var (
	password = "guest"
	username = "guest"
	host     = "localhost"
	port     = 5672

	QueueNameHuman = "queue_human"
	QueueNameCargo = "queue_cargo"
	QueueNameSatel = "queue_satellite"

	QueueNameAdminListen = "queue_admin_listen"
	QueueNameAdminCopy   = "queue_admin_copy"
	QueueNameAdminSpedit = "queue_admin_spedit"
	QueueNameAdminAgency = "queue_admin_agency"
	//QueueNameAdminSpeAge = "queue_admin_speditor_agency"

	AgencyName1 = "agency1"
	AgencyName2 = "agency2"
	SpeditName1 = "spedit1"
	SpeditName2 = "spedit2"

	AdminName = "admin"
)

func connectionString() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d/", username, password, host, port)
}

func SetupRabbitConnection() (*amqp.Connection, *amqp.Channel, error) {
	connection, err := amqp.Dial(connectionString())
	if err != nil {
		panic(err)
	}

	// opening a channel over the connection established to interact with RabbitMQ
	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}

	return connection, channel, nil
}

func SetupAdminQueues(channel *amqp.Channel) (*amqp.Queue, *amqp.Queue, <-chan amqp.Delivery, error) {
	// admin - sends on adminAgency, adminSpedit; listens on adminCopy
	queueSpedit, err := SetupSingleQueue(
		channel,
		QueueNameAdminSpedit,
	)
	if err != nil {
		return nil, nil, nil, err
	}

	queueAgency, err := SetupSingleQueue(
		channel,
		QueueNameCargo,
	)
	if err != nil {
		return nil, nil, nil, err
	}

	queueAdminCopy, err := SetupSingleConsumeQueue(
		channel,
		QueueNameAdminAgency,
		AdminName,
	)
	if err != nil {
		return nil, nil, nil, err
	}

	return &queueSpedit, &queueAgency, queueAdminCopy, nil
}

func SetupAgencyQueues(channel *amqp.Channel, name string) (*amqp.Queue, *amqp.Queue, *amqp.Queue, *amqp.Queue, <-chan amqp.Delivery, error) {
	// speditor1 - sends on human, cargo, satellite, adminCopy; listens on adminAgency
	// speditor2 - sends on human, cargo, satellite, adminCopy; listens on adminAgency
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
		consumerName,
	)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	queueCargo, err := SetupSingleConsumeQueue(
		channel,
		QueueNameCargo,
		consumerName,
	)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	queueSatellite, err := SetupSingleConsumeQueue(
		channel,
		QueueNameSatel,
		consumerName,
	)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	queueAdminListen, err := SetupSingleConsumeQueue(
		channel,
		QueueNameAdminSpedit,
		consumerName,
	)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	if consumerName == SpeditName1 {
		return queueHuman, queueCargo, nil, queueAdminListen, nil
	}
	return nil, queueCargo, queueSatellite, queueAdminListen, nil
}

func SetupSingleConsumeQueue(channel *amqp.Channel, queueName, consumerName string) (<-chan amqp.Delivery, error) {
	return channel.Consume(
		queueName,
		consumerName,
		true,
		false,
		false,
		false,
		nil,
	)
}

func SetupSingleQueue(channel *amqp.Channel, queueName string) (amqp.Queue, error) {
	return channel.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
}
