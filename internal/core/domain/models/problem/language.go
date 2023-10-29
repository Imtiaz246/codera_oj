package problem

import "gorm.io/gorm"

type Language struct {
	gorm.Model
	Name string `gorm:"unique;required"`
}
