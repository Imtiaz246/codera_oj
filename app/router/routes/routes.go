package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/imtiaz246/codera_oj/app/handler"
	"github.com/imtiaz246/codera_oj/services/middlewares"
	"time"
)

func RegisterRoutes(app *fiber.App, handler *handler.Handler) {
	apiV1 := app.Group("/api/v1")
	app.Use(csrf.New(csrf.Config{
		KeyLookup:      "header:X-Csrf-Token",
		CookieName:     "csrf_",
		CookieSameSite: "Lax",
		Expiration:     1 * time.Hour,
		KeyGenerator:   utils.UUID,
		Storage:        storage,
		//CookieHTTPOnly: true,
		//CookieSecure:   true,
	}))
	app.Use(helmet.New())
	//app.Use(limiter.New())

	/* -------------------- Auth Routes Begins -------------------- */
	auth := apiV1.Group("/auth")
	auth.Use(csrf.New())
	auth.Post("/signup", handler.SignUp)
	auth.Post("/login", handler.Login)
	auth.Get("/renew-token", handler.RenewToken)
	auth.Get("/verify-email/:id/:token", handler.VerifyEmail)
	/* -------------------- Auth Routes Ends -------------------- */

	/* -------------------- User Routes Begins -------------------- */
	user := apiV1.Group("/users")
	user.Get("/:username", handler.GetUserByUsername)

	user.Use(middlewares.New(middlewares.NewPasetoDefaultConfig()))
	user.Put("/:username", handler.UpdateUser)
	user.Put("/:username/password", handler.UpdatePassword)
	/* -------------------- User Routes Ends -------------------- */
}
