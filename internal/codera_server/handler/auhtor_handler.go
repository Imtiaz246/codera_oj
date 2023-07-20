package handler

import (
	"github.com/gofiber/fiber/v2"
	apiv1 "github.com/imtiaz246/codera_oj/internal/codera_server/api/v1"
	"github.com/imtiaz246/codera_oj/models"
	"github.com/imtiaz246/codera_oj/utils"
	"net/http"
	"strconv"
)

// UpdateProblem creates a problem
// HealthCheck godoc
// @Summary UpdateProblem updates a problem.
// @Description updates problem with the new information.
// @Tags author
// @Param data body apiv1.UpdateProblemOptions true "data"
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /author/problem/{id} [put]
func UpdateProblem(ctx *fiber.Ctx) error {
	req := new(apiv1.UpdateProblemOption)
	if err := BindAndValidate(ctx, req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewError(err))
	}
	user := GetUserFromCtx(ctx)
	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewError(err))
	}
	problem := req.UpdateProblemModelFormat(user, uint(id))
	if err := models.UpdateRecord[*models.Problem](problem); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	return ctx.Status(http.StatusOK).JSON(utils.NewSuccessResp(nil, "Problem Updated Successfully"))
}

// CreateProblem creates a problem
// HealthCheck godoc
// @Summary CreateProblem creates a problem.
// @Description creates problem for the oj.
// @Tags author
// @Param data body apiv1.CreateProblemOptions true "data"
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /author/problem/ [post]
func CreateProblem(ctx *fiber.Ctx) error {
	req := new(apiv1.CreateProblemOption)
	if err := BindAndValidate(ctx, req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewError(err))
	}
	user := GetUserFromCtx(ctx)

	problem := &models.Problem{
		Author: user,
		Title:  req.Title,
	}
	if err := models.CreateRecord[*models.Problem](problem); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	return ctx.Status(http.StatusOK).JSON(problem)
}

func ShareProblem(ctx *fiber.Ctx) error {

	return nil
}

func AddDataset(ctx *fiber.Ctx) error {

	return nil
}

func AddProblemTag(ctx *fiber.Ctx) error {

	return nil
}

func AddProblemSolution(ctx *fiber.Ctx) error {

	return nil
}

func AddDiscussionMessage(ctx *fiber.Ctx) error {

	return nil
}
