package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imtiaz246/codera_oj/internal/codera_server/handler"
	"github.com/imtiaz246/codera_oj/internal/middlewares"
)

func registerUserRoutes(apiV1 fiber.Router) {
	user := apiV1.Group("/users")
	user.Get("/:username", handler.GetUserByUsername)

	user.Use(middlewares.NewPasetoMiddleware())
	user.Put("/:username", handler.UpdateUser)
	user.Put("/:username/password", handler.UpdatePassword)
}
