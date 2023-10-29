package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func ReqUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		//pasetoPayload := ctx.Locals("payload").(*paseto.JSONToken)
		//username := pasetoPayload.Get("username")
		//user := new(auth.User)
		//if err := models.GetUserByUsername(username, user); err != nil {
		//	return ctx.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
		//}
		//ctx.Locals("user", user)
		//
		return ctx.Next()
	}
}
