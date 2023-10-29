package models

import (
	"github.com/imtiaz246/codera_oj/internal/core/domain/models/auth"
	"github.com/imtiaz246/codera_oj/internal/core/domain/models/problem"
)

type ModelFactory interface {
	*auth.User |
		*auth.VerifyEmail |
		*auth.Session |
		*problem.Problem |
		*problem.ProblemTag |
		*problem.Discussion |
		*problem.Solution |
		*problem.Share |
		*problem.Dataset |
		*problem.Language
}
