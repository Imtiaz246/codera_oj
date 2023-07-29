package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/imtiaz246/codera_oj/internal/codera/router/routes"
	"net/http"
)

// New creates a new fiber app with dependent handlers and routes and necessary configuration
func New() (*fiber.App, error) {
	app := fiber.New()
	app.Use(logger.New())
	app.Static("./public", "./public")
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, HEAD, PUT, PATCH, POST, DELETE",
	}))
	//isRunningOnProdMode := func() bool {
	//	appConfig := config.Settings.App
	//	if appConfig.RunMode == "dev" {
	//		return false
	//	} else {
	//		return false
	//	}
	//}
	//app.Use(csrf.New(csrf.Config{
	//	KeyLookup:      "cookie:csrf_",
	//	CookieName:     "csrf_",
	//	CookieHTTPOnly: isRunningOnProdMode(),
	//	CookieSecure:   isRunningOnProdMode(),
	//}))
	//app.Use(limiter.New())

	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON("pong")
	})
	routes.RegisterRoutes(app)

	return app, nil
}
