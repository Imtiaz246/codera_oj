package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imtiaz246/codera_oj/internal/codera_server/handler"
	"github.com/imtiaz246/codera_oj/middlewares"
)

func registerUserRoutes(apiV1 fiber.Router) {
	user := apiV1.Group("/users")
	user.Get("/:username", handler.GetUserByUsername)

	user.Use(middlewares.New(middlewares.NewPasetoDefaultConfig()))
	user.Put("/:username", handler.UpdateUser)
	user.Put("/:username/password", handler.UpdatePassword)
}
