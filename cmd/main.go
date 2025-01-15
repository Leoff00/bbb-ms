package main

import (
	"os"

	"github.com/leoff00/picpay-ms/modules"
	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()

	err := modules.App().Listen(":3001")
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to start the server")
	}

}
