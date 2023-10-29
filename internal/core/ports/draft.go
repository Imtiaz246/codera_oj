package ports

import (
	"context"
	"github.com/imtiaz246/codera_oj/internal/core/domain/dto"
)

// DraftService represents an interface for managing draft-related operations in the online judge system.
type DraftService interface {
	// CreateProblem creates a new problem.
	CreateProblem(ctx context.Context, opts *dto.CreateProblemOption) error

	// UpdateProblem updates an existing problem.
	UpdateProblem(ctx context.Context, id uint, opts *dto.UpdateProblemOption) error

	// AddDataset appends a dataset to a particular problem.
	AddDataset(ctx context.Context, id uint, opts *dto.DatasetOption) error

	// ShareProblem shares a problem with a designated user within the system.
	ShareProblem(ctx context.Context, id uint, opts *dto.ShareProblemOption) error

	// AddProblemSolution adds a solution to a specific problem.
	AddProblemSolution(ctx context.Context, id uint, opts *dto.SolutionOption) error

	// UpdateProblemSolution updates a solution for a particular problem.
	UpdateProblemSolution(ctx context.Context) error

	// DeleteProblemSolution removes a solution for a specific problem.
	DeleteProblemSolution(ctx context.Context) error

	// AddDiscussionMessage posts a discussion message related to a problem or contest.
	AddDiscussionMessage(ctx context.Context) error

	// AddProblemTag associates a tag with a particular problem.
	AddProblemTag(ctx context.Context) error
}
