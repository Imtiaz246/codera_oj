package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imtiaz246/codera_oj/internal/codera/auth"
	"github.com/imtiaz246/codera_oj/internal/codera/author"
	"github.com/imtiaz246/codera_oj/internal/codera/problem"
	"github.com/imtiaz246/codera_oj/internal/codera/user"
	"github.com/imtiaz246/codera_oj/internal/middlewares"
)

func RegisterRoutes(app *fiber.App) {
	apiv1 := app.Group("/api/v1")

	{
		// Auth routes
		authGrp := apiv1.Group("/auth")
		authGrp.Post("/signup", auth.SignUp)
		authGrp.Post("/login", auth.Login)
		authGrp.Get("/renew-token", auth.RenewToken)
		authGrp.Get("/verify-email/:id/:token", auth.VerifyEmail)
	}

	{
		// User routes
		userGrp := apiv1.Group("/users")
		userGrp.Get("/:username", user.GetUserByUsername)

		userGrp.Use(middlewares.NewPasetoMiddleware())
		userGrp.Put("/:username", user.UpdateUser)
		userGrp.Put("/:username/password", user.UpdatePassword)
	}

	{
		// Problem Routes
		problemGrp := apiv1.Group("/problems")
		problemGrp.Get("/", problem.GetProblemSet)
		problemGrp.Get("/:id", problem.GetProblemUsingID)
	}

	{
		// Author Routes
		authorGrp := apiv1.Group("/author")
		authorGrp.Use(
			middlewares.NewPasetoMiddleware(),
			middlewares.ReqUser(),
		)
		{
			pAuthorGrp := authorGrp.Group("/problems")
			pAuthorGrp.Post("/", author.CreateProblem)
			pAuthorGrp.Put("/:id", author.UpdateProblem)
			pAuthorGrp.Post("/:id/dataset", author.AddDataset)
			pAuthorGrp.Post("/:id/share", author.ShareProblem)
			pAuthorGrp.Post("/:id/tag", author.AddProblemTag)
			pAuthorGrp.Post("/:id/solutions", author.AddProblemSolution)
			pAuthorGrp.Put("/:id/solutions/:sid", author.UpdateProblemSolution)
			pAuthorGrp.Post("/:id/discussions", author.AddDiscussionMessage)
			pAuthorGrp.Delete("/:id/solutions/:sid", author.DeleteProblemSolution)
		}
		{
			//cAuthor := author.Group("/contest")
			//cAuthor.Post("/") // FixMe
		}
	}
}
