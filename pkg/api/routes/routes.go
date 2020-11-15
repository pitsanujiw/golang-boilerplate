package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func NewHTTP(app *fiber.App) {

	api := app.Group("/api")
	r := api.Group("/v1")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello GOLANG")
	})

	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("api/v1/ Root")
	})

	r.Get("/:id", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(&fiber.Map{
			"name":    c.Params("id"),
			"face_id": "editor",
			"item":    "data",
		})
	})
}
