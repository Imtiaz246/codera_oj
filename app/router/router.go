package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/imtiaz246/codera_oj/app/router/routes"
)

// New creates a new fiber app with dependent handlers and routes and necessary configuration
func New() (*fiber.App, error) {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, HEAD, PUT, PATCH, POST, DELETE",
	}))
	app.Get("/swagger/*", swagger.HandlerDefault)

	routes.RegisterRoutes(app)

	return app, nil
}
