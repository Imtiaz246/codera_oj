package routes

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/imtiaz246/codera_oj/app/handler"
	"github.com/imtiaz246/codera_oj/initializers/config"
)

func RegisterRoutes(app *fiber.App, handler *handler.Handler) {
	apiV1 := app.Group("/api/v1")

	/* -------------------- Auth Routes Begins -------------------- */
	auth := apiV1.Group("/auth")
	auth.Post("/signup", handler.SignUp)
	auth.Get("/verify-email/:id/:token", handler.VerifyEmail)
	auth.Post("/token", handler.RefreshToken)
	/* -------------------- Auth Routes Ends -------------------- */

	/* -------------------- User Routes Begins -------------------- */
	user := apiV1.Group("/users")
	user.Get("/:username", handler.GetUserByUsername)
	user.Put("/:username", handler.UpdateUser)
	user.Use(jwtware.New(jwtware.Config{
		SigningKey: config.GetAuthConfig().JWTSecret,
	}))
	user.Put("/:username/password", handler.UpdatePassword)
	/* -------------------- User Routes Ends -------------------- */
}
