package models

import (
	"github.com/imtiaz246/codera_oj/models/db"
	"gorm.io/gorm"
	"time"
)

type ProblemDiscussion struct {
	gorm.Model
	Message string

	UserID    uint
	User      *User
	ProblemID uint
	Problem   *Problem

	OwnerShipType OwnershipType
	SentAt        time.Time
}

func init() {
	if err := db.MigrateModelTables(ProblemChangeLog{}); err != nil {
		panic(err)
	}
}
