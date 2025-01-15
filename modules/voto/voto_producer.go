package voto

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog"
)

type VotoProducer struct{}

func (vp *VotoProducer) VotoProducer(voto int) error {
	var logger = zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	channel, queue, err := GetChannel()
	if err != nil {
		return err
	}
	defer channel.Close()

	err = channel.PublishWithContext(ctx, "", queue.Name, false, false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        []byte(strconv.Itoa(voto)),
		})

	if err != nil {
		logger.Err(err).Msg("Message cannot be published")
		return err
	}

	logger.Info().Msg("Message published successfully")

	return nil
}
