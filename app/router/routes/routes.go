package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/imtiaz246/codera_oj/app/handler"
	"github.com/imtiaz246/codera_oj/custom/config"
)

func RegisterRoutes(app *fiber.App, handler *handler.Handler) {
	isRunningOnProdMode := func() bool {
		appConfig := config.Cfg.App
		if appConfig.RunMode == "dev" {
			return false
		} else {
			return false
		}
	}
	app.Use(csrf.New(csrf.Config{
		KeyLookup:      "cookie:csrf_",
		CookieName:     "csrf_",
		CookieHTTPOnly: isRunningOnProdMode(),
		CookieSecure:   isRunningOnProdMode(),
	}))
	//app.Use(limiter.New())

	apiV1 := app.Group("/api/v1")
	registerAuthRoutes(apiV1, handler)
	registerUserRoutes(apiV1, handler)
	registerProblemRoutes(apiV1, handler)
}
