package problem

import (
	"github.com/imtiaz246/codera_oj/internal/core/domain/models"
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	TagName string `gorm:"unique;index"`
	UserID  uint
	AddedBy *models.User `gorm:"foreignKey:UserID"`
}
