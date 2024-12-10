// main.go
package main

import (
	routes "GOMOCKER/routes/init"

	"github.com/gofiber/fiber/v2"
)

const port = ":8882"

func main() {
	app := fiber.New()
	routes.SetupRoutes(app)

	app.Listen(port)
}
