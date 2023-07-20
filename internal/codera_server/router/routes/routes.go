package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imtiaz246/codera_oj/internal/codera_server/handler"
	"github.com/imtiaz246/codera_oj/internal/middlewares"
)

func RegisterRoutes(app *fiber.App) {
	apiv1 := app.Group("/api/v1")

	{
		// Auth routes
		auth := apiv1.Group("/auth")
		auth.Post("/signup", handler.SignUp)
		auth.Post("/login", handler.Login)
		auth.Get("/renew-token", handler.RenewToken)
		auth.Get("/verify-email/:id/:token", handler.VerifyEmail)
	}

	{
		// User routes
		user := apiv1.Group("/users")
		user.Get("/:username", handler.GetUserByUsername)

		user.Use(middlewares.NewPasetoMiddleware())
		user.Put("/:username", handler.UpdateUser)
		user.Put("/:username/password", handler.UpdatePassword)
	}

	{
		// Problem Routes
		problems := apiv1.Group("/problems")
		problems.Get("/", handler.GetProblemSet)
		problems.Get("/:id", handler.GetProblemUsingID)
	}

	{
		// Author Routes
		author := apiv1.Group("/author")
		author.Use(
			middlewares.NewPasetoMiddleware(),
			middlewares.ReqUser(),
		)
		{
			pAuthor := author.Group("/problem")
			pAuthor.Post("/", handler.CreateProblem)
			pAuthor.Put("/:id", handler.UpdateProblem)
			pAuthor.Post("/:id/dataset", handler.AddDataset)
			pAuthor.Post("/:id/share-with", handler.ShareProblem)
			pAuthor.Post("/:id/tag", handler.AddProblemTag)
			pAuthor.Post("/:id/solution", handler.AddProblemSolution)
			pAuthor.Post("/:id/discussion-message", handler.AddDiscussionMessage)
		}
		{
			//cAuthor := author.Group("/contest")
			//cAuthor.Post("/") // FixMe
		}
	}
}
