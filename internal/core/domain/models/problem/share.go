package problem

import (
	"github.com/imtiaz246/codera_oj/internal/core/domain/models"
	"gorm.io/gorm"
)

type Share struct {
	gorm.Model
	UserID         uint
	SharedWith     *models.User `gorm:"foreignKey:UserID"`
	ProblemID      uint
	Problem        *Problem
	PermissionType PermitType `gorm:"default:Viewer"`
}
