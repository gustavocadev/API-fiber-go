package main

import (
	"goFiber/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	// instance of fiber
	app := fiber.New()

	// middleware
	app.Use(logger.New())

	// Routes
	routes.UseRoute(app.Group("/api/users"))

	log.Fatal(app.Listen(":3000"))
}
