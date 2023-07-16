package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	apiv1 "github.com/imtiaz246/codera_oj/internal/codera_server/api/v1"
	"github.com/imtiaz246/codera_oj/utils"
	"net/http"
)

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
	req := new(apiv1.CreateProblemOption)
	if err := BindAndValidate(ctx, req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewError(err))
	}

	pasetoPayload := ctx.Locals("payload")
	fmt.Println(pasetoPayload)

	return nil
}

func ShareProblem(ctx *fiber.Ctx) error {

	return nil
}
