package http

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"net/http"
)

type Router struct {
	*fiber.App
}

func NewRouter(v *validator.Validate,
	ah *AuthHandler) *Router {
	app := fiber.New()
	app.Use(logger.New())
	app.Static("./static", "./public")
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, HEAD, PUT, PATCH, POST, DELETE",
	}))

	//r.Use(csrf.New(csrf.Config{
	//	KeyLookup:      "cookie:csrf_",
	//	CookieName:     "csrf_",
	//	CookieHTTPOnly: false,
	//	CookieSecure:   false,
	//}))
	//r.Use(limiter.New())

	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON("pong")
	})

	v1 := app.Group("/api/v1")
	{
		v1.Route("/auth", func(r fiber.Router) {
			r.Post("/signup", ah.SignUp)
			r.Post("/login", ah.Login)
			r.Get("/renew-token", ah.RenewToken)
			r.Get("/verify-email/:id/:token", ah.VerifyEmail)
		})
	}

	return &Router{app}
}

// Serve starts the HTTP server
func (r *Router) Serve(listenAddr string) error {
	return r.Serve(listenAddr)
}
