package main

import (
	"fmt"

	routes "github.com/AntonyIS/modart/api/http"
	"github.com/gofiber/fiber"
)

func main() {
	fmt.Println("Welcome to the Modular Art API ...")

	app := fiber.New()

	routes.SetupRoutes(app)

	app.Listen(5000)

}
