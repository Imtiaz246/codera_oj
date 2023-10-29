package repo

import (
	"github.com/imtiaz246/codera_oj/internal/adapters/repo/db"
	"github.com/imtiaz246/codera_oj/internal/core/domain/models/auth"
	problem_model "github.com/imtiaz246/codera_oj/internal/core/domain/models/problem"
	"github.com/imtiaz246/codera_oj/internal/core/ports"
	"os"
)

type problemRepo struct {
	ports.GenericInterface[*problem_model.Problem]
	*db.Database
}

var _ ports.ProblemRepoInterface = (*problemRepo)(nil)

func NewProblemRepo(d *db.Database) ports.ProblemRepoInterface {
	if err := d.DB.AutoMigrate(auth.User{}); err != nil {
		os.Exit(1)
	}
	return &problemRepo{
		Database:         d,
		GenericInterface: NewGenericRepo[*problem_model.Problem](d),
	}
}
