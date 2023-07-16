package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imtiaz246/codera_oj/internal/codera_server/handler"
	"github.com/imtiaz246/codera_oj/middlewares"
)

func registerProblemRoutes(apiV1 fiber.Router) {
	{
		problemSet := apiV1.Group("/problemset")
		problemSet.Get("/", handler.GetProblemSet)
	}

	{
		problem := apiV1.Group("/problem")
		problem.Get("/:id", handler.GetProblemUsingID)

		problem.Use(middlewares.New(middlewares.NewPasetoDefaultConfig()))
		problem.Post("/", handler.CreateProblem)
		problem.Put("/:id", handler.UpdateProblemUsingID)
	}
}
