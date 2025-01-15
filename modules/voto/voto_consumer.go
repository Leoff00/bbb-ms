package voto

import (
	"context"
	"os"
	"strconv"

	"github.com/rs/zerolog"
)

type VotoConsumer struct {
	Vr *VotoRepository
}

func (vc *VotoConsumer) Consume() error {
	var logger = zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
	ctx := context.Background()

	channel, queue, err := GetChannel()
	if err != nil {
		return err
	}
	defer channel.Close()

	msgs, err := channel.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logger.Err(err).Msg("Failed to register a consumer")
		return err
	}

	go func() {
		for d := range msgs {
			logger.Info().Msgf("Receiving msg and in consumer -> %s", d.Body)
			toStr := string(d.Body)
			toInt, _ := strconv.Atoi(toStr)
			if err := vc.Vr.Save(toInt); err != nil {
				logger.Err(err).Msg("Failed to save to DB")
			}
		}
	}()

	logger.Info().Msg("Waiting for messages. Press Ctrl+C to exit.")
	<-ctx.Done()
	return nil
}
