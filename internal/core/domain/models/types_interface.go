package models

import "github.com/imtiaz246/codera_oj/internal/core/domain/models/problem"

type ModelFactory interface {
	*User |
		*VerifyEmail |
		*Session |
		*problem.Problem |
		*problem.ProblemTag |
		*problem.Discussion |
		*problem.Solution |
		*problem.Share |
		*problem.Dataset |
		*problem.Language
}
