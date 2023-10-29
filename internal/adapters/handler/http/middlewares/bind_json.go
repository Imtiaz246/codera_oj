package middlewares

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/imtiaz246/codera_oj/internal/utils"
	"net/http"
)

var v *validator.Validate

func init() {
	v = validator.New()
}

// validate validates a user defined structs
func validate(s interface{}) error {
	if err := v.Struct(s); err != nil {
		var errMsg string
		for _, e := range err.(validator.ValidationErrors) {
			errMsg += fmt.Sprintf("%s field validation failed on tag '%s', actual value is '%s'\n",
				e.Field(), e.Tag(), e.Value())
		}
		return fmt.Errorf(errMsg)
	}

	return nil
}

// bindJsonAndValidate binds request json payload and validates the
// requested payload.
func bindJsonAndValidate(ctx *fiber.Ctx, d any) error {
	if err := ctx.BodyParser(d); err != nil {
		return err
	}
	if err := validate(d); err != nil {
		return err
	}
	return nil
}

func BindJson(d any) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if err := bindJsonAndValidate(ctx, d); err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(utils.NewError(err))
		}
		ctx.Locals("body", d)

		return ctx.Next()
	}
}
