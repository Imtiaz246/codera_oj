package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/imtiaz246/codera_oj/internal/adapters/handler/http/middlewares"
	"github.com/imtiaz246/codera_oj/internal/core/domain/dto"
	"net/http"
)

type Router struct {
	*fiber.App
}

func NewRouter(tc *dto.TokenConfig,
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
		v1.Route("/draft", func(r fiber.Router) {
			r.Use(middlewares.NewPasetoMiddleware(tc))
			r.Route("/problem", func(r fiber.Router) {
				r.Post("/", CreateProblem)
				r.Put("/:id", UpdateProblem)
				r.Post("/:id/dataset", AddDataset)
				r.Post("/:id/share", ShareProblem)
				r.Post("/:id/tag", AddProblemTag)
				r.Post("/:id/solutions", AddProblemSolution)
				r.Put("/:id/solutions/:sid", UpdateProblemSolution)
				r.Post("/:id/discussions", AddDiscussionMessage)
				r.Delete("/:id/solutions/:sid", DeleteProblemSolution)
			})
		})
	}

	return &Router{app}
}

// Serve starts the HTTP server
func (r *Router) Serve(listenAddr string) error {
	return r.Serve(listenAddr)
}
