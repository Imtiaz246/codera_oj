package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imtiaz246/codera_oj/app/handler"
	"github.com/imtiaz246/codera_oj/modules/middlewares"
)

func registerUserRoutes(apiV1 fiber.Router, handler *handler.Handler) {
	user := apiV1.Group("/users")
	user.Get("/:username", handler.GetUserByUsername)

	user.Use(middlewares.New(middlewares.NewPasetoDefaultConfig()))
	user.Put("/:username", handler.UpdateUser)
	user.Put("/:username/password", handler.UpdatePassword)
}
