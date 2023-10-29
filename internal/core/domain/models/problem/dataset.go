package problem

import (
	"github.com/imtiaz246/codera_oj/internal/core/domain/models/auth"
	"gorm.io/gorm"
)

type Dataset struct {
	gorm.Model
	Title    string
	Weight   int64 `gorm:"default:100;index:,sort:desc"`
	IsSample bool
	// todo: change to file store, because it's too expensive
	Input  []byte
	Output []byte

	UserID    uint
	AddedBy   *auth.User `gorm:"foreignKey:UserID"`
	ProblemID uint
	Problem   *Problem
}
