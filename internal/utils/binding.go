package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/imtiaz246/codera_oj/models"
	"reflect"
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

// GetUserFromCtx extracts requested user from ctx and returns it.
// The user will be assigned to context from middleware
func GetUserFromCtx(ctx *fiber.Ctx) *models.User {
	return ctx.Locals("user").(*models.User)
}

// GetFromReqBody reads data from ctx.Locals("body") which was set by the BindJson middleware earlier
func GetFromReqBody[T any](ctx *fiber.Ctx) (T, error) {
	var t T
	if reflect.TypeOf(t).Kind() != reflect.Pointer {
		return t, fmt.Errorf("type has to be a pointer while getting data from ctx")
	}
	body := ctx.Locals("body")
	if reflect.TypeOf(body) != reflect.TypeOf(t) {
		return t, fmt.Errorf("type mismatch while getting data from ctx")
	}
	t = body.(T)

	return t, nil
}
