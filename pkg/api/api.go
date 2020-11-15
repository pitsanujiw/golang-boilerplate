package api

import (
	"github.com/gofiber/fiber/v2"
)

func StartAPI() *fiber.App {
	app := fiber.New()
	return app
}
