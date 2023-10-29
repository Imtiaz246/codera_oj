package draft

import (
	"context"
	"errors"
	"github.com/imtiaz246/codera_oj/internal/core/domain/dto"
	"github.com/imtiaz246/codera_oj/internal/core/domain/models/auth"
	"github.com/imtiaz246/codera_oj/internal/core/domain/models/problem"
	"github.com/imtiaz246/codera_oj/internal/core/ports"
)

type draftService struct {
	problemRepo ports.ProblemRepoInterface
	problemShareRepo ports.ShareRepoInterface
	datasetRepo ports.DatasetRepoInterface
}

var _ ports.DraftService = (*draftService)(nil)

func NewDraftService(pr ports.ProblemRepoInterface,
	ps ports.ShareRepoInterface,
	dr ports.DatasetRepoInterface) ports.DraftService {
	return &draftService{
		problemRepo: pr,
		problemShareRepo: ps,
		datasetRepo: dr,
	}
}

func (d *draftService) CreateProblem(ctx context.Context, opts *dto.CreateProblemOption) error {
	// FIXME: dummy user will get user from context
	user := &auth.User{}

	p := &problem.Problem{
		Author: user,
		Title:  opts.Title,
	}
	if err := d.problemRepo.CreateRecord(p); err != nil {
		return err
	}

	return nil
}

func (d *draftService) UpdateProblem(ctx context.Context, id uint, opts *dto.UpdateProblemOption) error {
	// FIXME: dummy user will get user from context
	user := &auth.User{}

	p, err := d.problemRepo.GetRecordByID(id)
	if err != nil {
		return err
	}
	if p.AuthorID != user.ID {
		return errors.New("statusForbidden")
	}

	p.TimeLimit = opts.TimeLimit
	p.MemoryLimit = opts.MemoryLimit
	p.Statement = opts.Statement
	p.InputStatement = opts.InputStatement
	p.OutputStatement = opts.OutputStatement
	p.NoteStatement = opts.NoteStatement
	p.StatementsVisibilityDuringContest = opts.StatementsVisibilityDuringContest
	p.CheckerType = opts.CheckerType

	if err = d.problemRepo.UpdateRecord(p); err != nil {
		return errors.New("internalError")
	}

	return nil
}

func (d *draftService) AddDataset(ctx context.Context, pid uint, opts *dto.DatasetOption) error {
	// FIXME: dummy user will get user from context
	user := &auth.User{}

	problemShare := &problem.Share{
		SharedWith: user,
		ProblemID:  pid,
	}
	if err := d.problemShareRepo.GetRecordByModel(problemShare); err != nil {
		return errors.New("internalError")
	}
	if !problemShare.CanAddDataset() {
		return errors.New("forbidden")
	}

	dataset := &problem.Dataset{
		Title: opts.Title,
		Weight: opts.Weight,
		IsSample: opts.IsSample,
		Input: opts.InputContent,
		Output: opts.OutputContent,
		UserID: user.ID,
		ProblemID: pid,
	}

	if err := d.datasetRepo.CreateRecord(dataset); err != nil {
		return errors.New("internalError")
	}

	return nil
}

func (d *draftService) ShareProblem(ctx context.Context, id uint, opts *dto.ShareProblemOption) error {

	return nil
}

func (d *draftService) AddProblemSolution(ctx context.Context, id uint, opts *dto.SolutionOption) error {

	return nil
}

func (d *draftService) UpdateProblemSolution(ctx context.Context) error {

	return nil
}

func (d *draftService) DeleteProblemSolution(ctx context.Context) error {

	return nil
}

func (d *draftService) AddDiscussionMessage(ctx context.Context) error {

	return nil
}

func (d *draftService) AddProblemTag(ctx context.Context) error {

	return nil
}
