package problem

import (
	"github.com/imtiaz246/codera_oj/internal/core/domain/models/auth"
	"gorm.io/gorm"
	"time"
)

type Discussion struct {
	gorm.Model
	Message string

	UserID    uint
	User      *auth.User
	ProblemID uint
	Problem   *Problem
	SentAt    time.Time
}
