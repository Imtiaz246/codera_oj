package handler

import (
	"github.com/gofiber/fiber/v2"
	apiv1 "github.com/imtiaz246/codera_oj/internal/codera_server/api/v1"
	"github.com/imtiaz246/codera_oj/models"
	"github.com/imtiaz246/codera_oj/utils"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetProblemSet(ctx *fiber.Ctx) error {

	return nil
}

func GetProblemUsingID(ctx *fiber.Ctx) error {

	return nil
}

// UpdateProblem creates a problem
// HealthCheck godoc
// @Summary UpdateProblem updates a problem.
// @Description updates problem with the new information.
// @Tags problem
// @Param data body apiv1.UpdateProblemOptions true "data"
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /problem/{id} [put]
func UpdateProblem(ctx *fiber.Ctx) error {
	req := new(apiv1.UpdateProblemOption)
	if err := BindAndValidate(ctx, req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewError(err))
	}
	user, err := GetUserFromCtx(ctx)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewError(err))
	}

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewError(err))
	}
	problem := &models.Problem{
		Model:                             gorm.Model{ID: uint(id)},
		Author:                            user,
		TimeLimit:                         req.TimeLimit,
		MemoryLimit:                       req.MemoryLimit,
		Statement:                         req.Statement,
		InputStatement:                    req.InputStatement,
		OutputStatement:                   req.OutputStatement,
		NoteStatement:                     req.NoteStatement,
		StatementsVisibilityDuringContest: req.StatementsVisibilityDuringContest,
		CheckerType:                       req.CheckerType,
	}
	if err := models.UpdateRecord[*models.Problem](problem); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	return ctx.Status(http.StatusOK).JSON("Problem updated successfully")
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
	user, err := GetUserFromCtx(ctx)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewError(err))
	}

	problem := &models.Problem{
		Author: user,
		Title:  req.Title,
	}
	if err = models.CreateRecord[*models.Problem](problem); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	return ctx.Status(http.StatusOK).JSON(problem)
}

func ShareProblem(ctx *fiber.Ctx) error {

	return nil
}

func CreateDataset(ctx *fiber.Ctx) error {

	return nil
}
