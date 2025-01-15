package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

var RateLimit = limiter.New(limiter.Config{
	Max:        10,
	Expiration: 5 * time.Second,
	LimitReached: func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
			"message": "Rate limit exceeded, try again later",
		})
	},
})
