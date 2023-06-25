package handler

import "github.com/gofiber/fiber/v2"

func GetProblemSet(ctx *fiber.Ctx) error {

	return nil
}

func GetProblemUsingID(ctx *fiber.Ctx) error {

	return nil
}

func UpdateProblemUsingID(ctx *fiber.Ctx) error {

	return nil
}

// CreateProblem creates a problem
// HealthCheck godoc
// @Summary CreateProblem creates a problem.
// @Description creates problem for the oj.
// @Tags problem
// @Param data body apiv1.CreateProblemOptions true "data"
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /problem/ [post]
func CreateProblem(ctx *fiber.Ctx) error {

	return nil
}

func ShareProblem(ctx *fiber.Ctx) error {

	return nil
}
