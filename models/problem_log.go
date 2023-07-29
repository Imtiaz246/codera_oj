package models

import (
	"github.com/imtiaz246/codera_oj/models/db"
	"gorm.io/gorm"
	"time"
)

type ProblemChangeLog struct {
	gorm.Model
	LogMessage string

	UserID    uint
	User      *User
	ProblemID uint
	Problem   *Problem

	OwnerShipType PermitType
	ChangedAt     time.Time
}

func init() {
	if err := db.MigrateModelTables(ProblemChangeLog{}); err != nil {
		panic(err)
	}
}
