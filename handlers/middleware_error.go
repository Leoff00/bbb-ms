package handlers

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

func ErrorHandler(c *fiber.Ctx) error {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	err := c.Next()
	if err != nil {
		logger.Error().Err(err).Msg("An error occurred during request handling")

		code := fiber.StatusInternalServerError
		message := "Internal Server Error"

		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
			message = e.Message
		}

		return c.Status(code).JSON(fiber.Map{
			"error":   true,
			"message": message,
		})
	}

	return nil
}
