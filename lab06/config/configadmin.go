package config

import "github.com/streadway/amqp"

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
		QueueNameAdminAgency,
	)
	if err != nil {
		return nil, nil, nil, err
	}

	queueAdminCopy, err := SetupSingleConsumeQueue(
		channel,
		QueueNameAdminCopy,
		AdminName,
	)
	if err != nil {
		return nil, nil, nil, err
	}

	return &queueSpedit, &queueAgency, queueAdminCopy, nil
}
