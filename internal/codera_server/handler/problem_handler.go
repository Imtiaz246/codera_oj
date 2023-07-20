package handler

import (
	"github.com/gofiber/fiber/v2"
)

func GetProblemSet(ctx *fiber.Ctx) error {

	return nil
}

func GetProblemUsingID(ctx *fiber.Ctx) error {

	ctx.Next()
	return nil
}
