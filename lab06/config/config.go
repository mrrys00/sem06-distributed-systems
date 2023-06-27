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

	QueueNameAdminCopy   = "queue_admin_copy"
	QueueNameAdminSpedit = "queue_admin_spedit"
	QueueNameAdminAgency = "queue_admin_agency"

	AgencyName1 = "agency1"
	AgencyName2 = "agency2"
	AgencyName  = "agency"
	userAgency  = "agency"
	passAgency  = "agency"

	SpeditName1  = "spedit1"
	SpeditName2  = "spedit2"
	SpeditorName = "speditor"
	userSpeditor = "speditor"
	passSpeditor = "speditor"

	AdminName = "admin"
	userAdmin = "admin"
	passAdmin = "admin"
)

func connectionString(name string) string {
	switch name {
	case AdminName:
		return fmt.Sprintf("amqp://%s:%s@%s:%d/", userAdmin, passAdmin, host, port)
	case AgencyName:
		return fmt.Sprintf("amqp://%s:%s@%s:%d/", userAgency, passAgency, host, port)
	case SpeditorName:
		return fmt.Sprintf("amqp://%s:%s@%s:%d/", userSpeditor, passSpeditor, host, port)
	}
	return fmt.Sprintf("amqp://%s:%s@%s:%d/", username, password, host, port)
}

func SetupRabbitConnection(name string) (*amqp.Connection, *amqp.Channel, error) {
	connection, err := amqp.Dial(connectionString(name))
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

func SetupSingleConsumeQueue(channel *amqp.Channel, queueName, clientName string) (<-chan amqp.Delivery, error) {
	return channel.Consume(
		queueName,
		clientName,
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
