package main

import (
	"govuets/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", handlers.GetHandler)

	app.Static("/", "./public")
	app.Listen(":3000")

}
