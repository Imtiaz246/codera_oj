package handler

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/imtiaz246/codera_oj/models"
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

// BindAndValidate binds request payload and validates the
// requested payload.
func BindAndValidate(ctx *fiber.Ctx, d any) error {
	if err := ctx.BodyParser(d); err != nil {
		return err
	}
	if err := validate(d); err != nil {
		return err
	}
	return nil
}

// GetUserFromCtx extracts requested user from ctx and returns it
func GetUserFromCtx(ctx *fiber.Ctx) (*models.User, error) {
	user := ctx.Locals("user").(*models.User)
	if user == nil {
		return nil, fmt.Errorf("user from context not found")
	}

	return user, nil
}
