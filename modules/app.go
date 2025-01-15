package modules

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leoff00/picpay-ms/handlers"
	"github.com/leoff00/picpay-ms/modules/voto"
)

func App() *fiber.App {
	app := fiber.New()
	api := app.Group("/api")
	app.Use(handlers.QuotaLimit)
	app.Use(handlers.RateLimit)

	app.Use(handlers.ErrorHandler)

	api.Post("/voto", voto.VotoController)
	return app
}
