package models

import (
	"github.com/imtiaz246/codera_oj/models/db"
	"gorm.io/gorm"
	"time"
)

type JudgerIdentity struct {
	gorm.Model
	DisplayName string
	Key         string `gorm:"unique,required"`
	Registered  bool   `gorm:"default:false"`
	NotBefore   time.Time
	NotAfter    time.Time
	OwnerID     uint
	Owner       *User `gorm:"required;foreignKey:OwnerID"`
}

func init() {
	if err := db.MigrateModelTables(JudgerIdentity{}); err != nil {
		panic(err)
	}
}
