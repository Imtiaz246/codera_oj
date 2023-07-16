package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imtiaz246/codera_oj/internal/codera_server/handler"
)

func registerAuthRoutes(apiV1 fiber.Router) {
	auth := apiV1.Group("/auth")
	auth.Post("/signup", handler.SignUp)
	auth.Post("/login", handler.Login)
	auth.Get("/renew-token", handler.RenewToken)
	auth.Get("/verify-email/:id/:token", handler.VerifyEmail)
}
