package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/leoff00/picpay-ms/db"
	"github.com/leoff00/picpay-ms/modules/voto"
	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
	db, err := db.Conn()

	if err != nil {
		logger.Err(err).Msg("Failed to connect with DB")

	}

	repository := voto.NewVotoRepository(db)
	consumer := voto.VotoConsumer{Vr: repository}

	go func() {
		if err := consumer.Consume(); err != nil {
			logger.Err(err).Msg("Cannot consume messages")
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan

	defer func() {
		db.Close()
		logger.Info().Msg("Database connection closed successfully")
	}()

	logger.Info().Msg("Received shutdown signal, exiting gracefully...")
}
