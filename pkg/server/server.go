package server

import (
	"fmt"
	api "go-boilerplate/pkg/api"
	ht "go-boilerplate/pkg/api/routes"
	"log"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func Start() error {
	app := api.StartAPI()
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
	}))
	ht.NewHTTP(app)
	fmt.Println("trying.....")
	log.Fatal(app.Listen(":3000"))

	return nil
}

