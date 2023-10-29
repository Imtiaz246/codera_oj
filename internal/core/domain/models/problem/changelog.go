package problem

import (
	"github.com/imtiaz246/codera_oj/internal/core/domain/models"
	"gorm.io/gorm"
	"time"
)

type ChangeLog struct {
	gorm.Model
	LogMessage string

	UserID    uint
	User      *models.User
	ProblemID uint
	Problem   *Problem

	OwnerShipType PermitType
	ChangedAt     time.Time
}
