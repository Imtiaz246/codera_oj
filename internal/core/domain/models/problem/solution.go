package problem

import (
	"github.com/imtiaz246/codera_oj/internal/core/domain/models"
	"gorm.io/gorm"
	"time"
)

type Solution struct {
	gorm.Model
	Code         string
	LanguageID   uint
	Language     *Language
	LastExecuted *time.Time
	TimeTaken    *float64
	MemoryTaken  *float64

	UserID    uint
	User      *models.User
	ProblemId uint
	Problem   *Problem
}
