package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/imtiaz246/codera_oj/custom/config"
)

func RegisterRoutes(app *fiber.App) {
	isRunningOnProdMode := func() bool {
		appConfig := config.Settings.App
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
	registerAuthRoutes(apiV1)
	registerUserRoutes(apiV1)
	registerProblemRoutes(apiV1)
}
