package author

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/imtiaz246/codera_oj/internal/codera/structs"
	"github.com/imtiaz246/codera_oj/internal/utils"
	"github.com/imtiaz246/codera_oj/models"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// CreateProblem creates a problem
// HealthCheck godoc
// @Summary creates a problem.
// @Description creates problem for the oj.
// @Tags author
// @Param data body CreateProblemOption true "data"
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /author/problems/ [post]
func CreateProblem(ctx *fiber.Ctx) error {
	req := new(structs.CreateProblemOption)
	if err := utils.BindAndValidate(ctx, req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewError(err))
	}
	user := utils.GetUserFromCtx(ctx)

	problem := &models.Problem{
		Author: user,
		Title:  req.Title,
	}
	if err := models.CreateRecord[*models.Problem](problem); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	return ctx.Status(http.StatusOK).JSON(utils.NewSuccessResp(problem, "Problem created successfully"))
}

// UpdateProblem creates a problem
// HealthCheck godoc
// @Summary updates a problem.
// @Description updates problem with the new information.
// @Tags author
// @Param data body UpdateProblemOption true "data"
// @Param id path string true "problem id"
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /author/problems/{id} [put]
func UpdateProblem(ctx *fiber.Ctx) error {
	req := new(structs.UpdateProblemOption)
	if err := utils.BindAndValidate(ctx, req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewError(err))
	}
	user := utils.GetUserFromCtx(ctx)
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewError(err))
	}

	problem, err := models.GetRecordByID[*models.Problem](strconv.Itoa(id))
	if problem.AuthorID != user.ID {
		return ctx.Status(http.StatusNotAcceptable).JSON(utils.NewError(fmt.Errorf("access forbidden")))
	}
	problem = req.UpdateProblemModelFormat(user, uint(id))
	if err := models.UpdateRecord[*models.Problem](problem); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	return ctx.Status(http.StatusOK).JSON(utils.NewSuccessResp(nil, "Problem Updated Successfully"))
}

// ShareProblem shares a problem with other user
// HealthCheck godoc
// @Summary shares a problem with other user
// @Description shares problem so that other user can contribute to that problem.
// @Tags author
// @Param data body ShareProblemOption true "data"
// @Param id path string true "problem id"
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /author/problems/{id}/ [put]
func ShareProblem(ctx *fiber.Ctx) error {
	req := new(structs.ShareProblemOption)
	if err := utils.BindAndValidate(ctx, req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewError(err))
	}
	user := utils.GetUserFromCtx(ctx)
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewError(err))
	}
	if !req.PermitType.IsPermitTypeValid() {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewError(fmt.Errorf("permitType `%s` is not valid", req.PermitType)))
	}

	problem := &models.Problem{
		Model:  gorm.Model{ID: uint(id)},
		Author: user,
	}
	if err = models.GetRecordByModel[*models.Problem](problem); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewError(err))
	}

	// FIXME: will be created multiple record for a single user
	problemShare := &models.ProblemShare{
		SharedWith:     user,
		Problem:        problem,
		PermissionType: req.PermitType,
	}
	if err = models.CreateRecord[*models.ProblemShare](problemShare); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	return ctx.Status(http.StatusCreated).JSON(utils.NewSuccessResp(nil, fmt.Sprintf("Problem sharing with `%s` as `%s` successful", req.ShareWith, req.PermitType)))
}

// AddDataset shares a problem with other user
// HealthCheck godoc
// @Summary adds dataset for a problem
// @Description adds datasets(input & output file) for a problem
// @Tags author
// @Param data body ShareProblemOption true "data"
// @Param id path string true "problem id"
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /author/problems/{id}/dataset [post]
func AddDataset(ctx *fiber.Ctx) error {
	req := new(structs.DatasetOption)
	if err := utils.BindAndValidate(ctx, req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewError(err))
	}
	user := utils.GetUserFromCtx(ctx)
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.NewError(err))
	}

	problemShare := &models.ProblemShare{
		SharedWith: user,
		ProblemID:  uint(id),
	}
	if err := models.GetRecordByModel(problemShare); err != nil {
		return ctx.Status(http.StatusForbidden).JSON(utils.NewError(err))
	}
	if !problemShare.CanAddDataset() {
		return ctx.Status(http.StatusForbidden).JSON(utils.NewError(err))
	}

	//inputFile, err := ctx.FormFile("input")
	//if err != nil {
	//	return ctx.Status(http.StatusBadRequest).JSON(utils.NewError(err))
	//}
	//outputFile, err := ctx.FormFile("output")
	//if err != nil {
	//	return ctx.Status(http.StatusBadRequest).JSON(utils.NewError(err))
	//}

	return nil
}

// AddProblemSolution adds solution for a problem
// HealthCheck godoc
// @Summary adds solution for a problem
// @Description adds solutions for a problem. Only authorized people(to whom have to that problem) can add solution
// @Tags author
// @Param data body SolutionOption true "data"
// @Param id path string true "problem id"
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /author/problems/{id}/solutions [post]
func AddProblemSolution(ctx *fiber.Ctx) error {
	req := new(structs.SolutionOption)
	if err := utils.BindAndValidate(ctx, req); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	fmt.Println(req)
	return nil
}

func UpdateProblemSolution(ctx *fiber.Ctx) error {

	return nil
}

// AddDiscussionMessage adds discussion messages for a problem
// HealthCheck godoc
// @Summary adds discussion messages for a problem
// @Description adds discussion messages for a problem
// @Tags author
// @Param data body DiscussionOption true "data"
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
