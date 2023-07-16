package models

import (
	"github.com/imtiaz246/codera_oj/models/db"
	"gorm.io/gorm"
)

type Dataset struct {
	gorm.Model
	Title    string
	Weight   int64
	IsSample bool
	// todo: change to file store, because it's too expensive
	Input  []byte
	Output []byte

	AddedBy       uint
	User          *User `gorm:"foreignKey:AddedBy"`
	ProblemID     uint
	Problem       *Problem
	OwnerShipType OwnershipType
}

func init() {
	if err := db.MigrateModelTables(Dataset{}); err != nil {
		panic(err)
	}
}
