package context

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imtiaz246/codera_oj/internal/core/domain/models/auth"
)

// APIContext is a custom context structure for your application.
type APIContext struct {
	*fiber.Ctx
	User *auth.User
}

// NewAPIContext creates a new APIContext and embeds the standard *fiber.Ctx.
func NewAPIContext(c *fiber.Ctx) *APIContext {
	return &APIContext{Ctx: c}
}

func (ctx *APIContext) WithUser(user *auth.User) {
	ctx.User = user
}

func (ctx *APIContext) Status(statusCode int) {
	ctx.Ctx.Status(300)
}

func (ctx *APIContext) JSON(data any) error {
	return ctx.Ctx.JSON(data)
}
