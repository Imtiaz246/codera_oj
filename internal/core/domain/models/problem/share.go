package problem

import (
	"github.com/imtiaz246/codera_oj/internal/core/domain/models/auth"
	"gorm.io/gorm"
)

type Share struct {
	gorm.Model
	UserID         uint
	SharedWith     *auth.User `gorm:"foreignKey:UserID"`
	ProblemID      uint
	Problem        *Problem
	PermissionType PermitType `gorm:"default:Viewer"`
}

func (ps *Share) CanAddDataset() bool {
	return ps.PermissionType == Author || ps.PermissionType == Editor
}

func (ps *Share) CanAddSolution() bool {
	return ps.PermissionType == Author || ps.PermissionType == Editor || ps.PermissionType == Tester
}