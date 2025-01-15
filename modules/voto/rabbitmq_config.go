package voto

import (
	"os"

	"github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog"
)

type ConfigQueue struct {
	name       string
	durable    bool
	autoDelete bool
	exclusive  bool
	noWait     bool
	args       amqp091.Table
}

func queueDeclare(channel *amqp091.Channel, name string) (amqp091.Queue, error) {
	var logger = zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()

	c := ConfigQueue{
		name:       name,
		durable:    true,
		autoDelete: false,
		exclusive:  false,
		noWait:     false,
		args:       nil,
	}

	q, err := channel.QueueDeclare(c.name, c.durable, c.autoDelete, c.exclusive, c.noWait, c.args)

	if err != nil {
		logger.Err(err).Msg("Failed to declare the Queue")
		return amqp091.Queue{}, err
	}

	return q, nil
}

func GetChannel() (*amqp091.Channel, amqp091.Queue, error) {
	var logger = zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()

	conn, err := amqp091.Dial("amqp://user:pass@localhost:5672/")
	if err != nil {
		logger.Err(err).Msg("Failed to create a connection with amqp")
		return nil, amqp091.Queue{}, err
	}

	channel, err := conn.Channel()
	if err != nil {
		logger.Err(err).Msg("Failed to open a channel")
		return nil, amqp091.Queue{}, err
	}

	queue, err := queueDeclare(channel, "voto-topic")
	if err != nil {
		logger.Err(err).Msg("Failed to declare the queue")
		return nil, amqp091.Queue{}, err
	}

	return channel, queue, nil
}
