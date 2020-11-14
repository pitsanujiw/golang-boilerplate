package api

import (
	"fmt"
	ht "go-boilerplate/pkg/api/transports"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Start() error {
	app := fiber.New()

	ht.NewHTTP(app)
	fmt.Println("trying.....")
	log.Fatal(app.Listen(":3000"))

	defer Stop()

	return nil
}

func Stop() {
	fmt.Println("Hello")
}
