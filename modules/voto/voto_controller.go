package voto

import (
	"github.com/gofiber/fiber/v2"
)

type VotoDTO struct {
	Voto int `json:"voto"`
}

func VotoController(c *fiber.Ctx) error {
	producer := VotoProducer{}
	usecase := NewVotoUseCase(&producer)

	var dto VotoDTO

	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Body Type",
		})

	}

	err := usecase.processVote(dto.Voto)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.SendStatus(fiber.StatusAccepted)
}
