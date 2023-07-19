package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	apiV1 := app.Group("/api/v1")
	registerAuthRoutes(apiV1)
	registerUserRoutes(apiV1)
	registerProblemRoutes(apiV1)
}
