package problem

import (
	"github.com/imtiaz246/codera_oj/internal/core/domain/models"
	"gorm.io/gorm"
	"time"
)

type Discussion struct {
	gorm.Model
	Message string

	UserID    uint
	User      *models.User
	ProblemID uint
	Problem   *Problem
	SentAt    time.Time
}
