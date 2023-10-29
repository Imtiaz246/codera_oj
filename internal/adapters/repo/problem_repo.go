package repo

import (
	problemmodel "github.com/imtiaz246/codera_oj/internal/core/domain/models/problem"
	"github.com/imtiaz246/codera_oj/internal/core/ports"
)

type ProblemRepo struct {
	ports.GenericInterface[*problemmodel.Problem]
}
