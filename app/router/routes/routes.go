package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imtiaz246/codera_oj/app/handler"
	"github.com/imtiaz246/codera_oj/services/middlewares"
)

func RegisterRoutes(app *fiber.App, handler *handler.Handler) {
	apiV1 := app.Group("/api/v1")

	/* -------------------- Auth Routes Begins -------------------- */
	auth := apiV1.Group("/auth")
	auth.Post("/signup", handler.SignUp)
	auth.Post("/login", handler.Login)
	auth.Get("renew-token", handler.RenewToken)
	auth.Get("/verify-email/:id/:token", handler.VerifyEmail)
	/* -------------------- Auth Routes Ends -------------------- */

	/* -------------------- User Routes Begins -------------------- */
	user := apiV1.Group("/users")
	user.Use(middlewares.New(middlewares.NewPasetoDefaultConfig()))
	user.Get("/:username", handler.GetUserByUsername)
	user.Put("/:username", handler.UpdateUser)
	user.Put("/:username/password", handler.UpdatePassword)
	/* -------------------- User Routes Ends -------------------- */
}
