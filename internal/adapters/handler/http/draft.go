package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imtiaz246/codera_oj/internal/adapters/handler/http/utils"
	"github.com/imtiaz246/codera_oj/internal/core/domain/dto"
	"github.com/imtiaz246/codera_oj/internal/core/domain/models/auth"
	"github.com/imtiaz246/codera_oj/internal/core/ports"
	"net/http"
)

// DraftHandler represents the HTTP handler for problem drafts-related requests
type DraftHandler struct {
	svc ports.DraftService
}

// NewDraftHandler creates a new DraftHandler instance
func NewDraftHandler(svc ports.DraftService) *DraftHandler {
	return &DraftHandler{
		svc,
	}
}

// CreateProblem creates a problem
// HealthCheck godoc
// @Summary creates a problem.
// @Description creates a problem for the oj.
// @Tags author
// @Param data body structs.CreateProblemOption true "data"
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /author/problems/ [post]
func CreateProblem(ctx *fiber.Ctx) error {
	req := new(dto.CreateProblemOption)
	if err := utils.BindAndValidate(ctx, req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewError(err))
	}
	user := ctx.Locals("user").(*auth.User)

}

// UpdateProblem creates a problem
// HealthCheck godoc
// @Summary updates a problem.
// @Description update problem with the new information.
// @Tags author
// @Param data body structs.UpdateProblemOption true "data"
// @Param id path string true "problem id"
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /author/problems/{id} [put]
func UpdateProblem(ctx *fiber.Ctx) error {
	req := new(dto.UpdateProblemOption)
	if err := utils.BindAndValidate(ctx, req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewError(err))
	}
	user := ctx.Locals("user").(*auth.User)
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewError(err))
	}

}

// ShareProblem shares a problem with other user
// HealthCheck godoc
// @Summary shares a problem with other user
// @Description shares problem so that another user can contribute to that problem.
// @Tags author
// @Param data body structs.ShareProblemOption true "data"
// @Param id path string true "problem id"
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /author/problems/{id}/ [put]
func ShareProblem(ctx *fiber.Ctx) error {
	req := new(dto.ShareProblemOption)
	if err := utils.BindAndValidate(ctx, req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewError(err))
	}
	user := ctx.Locals("user").(*auth.User)
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewError(err))
	}

}

// AddDataset shares a problem with other user
// HealthCheck godoc
// @Summary adds dataset for a problem
// @Description adds datasets(input & output file) for a problem
// @Tags author
// @Param data body structs.ShareProblemOption true "data"
// @Param id path string true "problem id"
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /author/problems/{id}/dataset [post]
func AddDataset(ctx *fiber.Ctx) error {
	req := new(dto.DatasetOption)
	if err := utils.BindAndValidate(ctx, req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewError(err))
	}
	user := ctx.Locals("user").(*auth.User)
	pid, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewError(err))
	}

}

// AddProblemSolution adds a solution for a problem
// HealthCheck godoc
// @Summary adds a solution for a problem
// @Description adds solutions for a problem. Only authorized people (to whom have to that problem) can add solution
// @Tags author
// @Param data body structs.SolutionOption true "data"
// @Param id path string true "problem id"
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /author/problems/{id}/solutions [post]
func AddProblemSolution(ctx *fiber.Ctx) error {
	req := new(dto.SolutionOption)
	if err := utils.BindAndValidate(ctx, req); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	user := ctx.Locals("user").(*auth.User)
	pid, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewError(err))
	}

}

func UpdateProblemSolution(ctx *fiber.Ctx) error {

	return nil
}

func DeleteProblemSolution(ctx *fiber.Ctx) error {

	return nil
}

// AddDiscussionMessage adds discussion messages for a problem
// HealthCheck godoc
// @Summary adds discussion messages for a problem
// @Description adds discussion messages for a problem
// @Tags author
// @Param data body structs.DiscussionOption true "data"
// @Param id path string true "problem id"
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /author/problems/{id}/discussions [post]
func AddDiscussionMessage(ctx *fiber.Ctx) error {

	return nil
}

func AddProblemTag(ctx *fiber.Ctx) error {

	return nil
}
