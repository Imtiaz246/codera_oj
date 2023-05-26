package db

import (
	"fmt"
	models2 "github.com/imtiaz246/codera_oj/app/models"
	"github.com/imtiaz246/codera_oj/initializers/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitializeDB() error {
	dbConfig := config.GetDBConfig()
	var err error
	switch dbConfig.DbType {
	case "sqlite3":
		db, err = connSqliteDB()
		if err != nil {
			return err
		}
	case "postgres":
		db, err = connPostgresDB()
		if err != nil {
			return err
		}
	default:
		fmt.Println("wrong database")
	}
	doAutoMigrations(db)
	applySeeds()

	return nil
}

func GetDB() *gorm.DB {
	return db
}
func connPostgresDB() (*gorm.DB, error) {
	fmt.Println("postgres db")
	return nil, nil
}

func connSqliteDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("gorm.sqlite"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func doAutoMigrations(db *gorm.DB) {
	db.AutoMigrate(&models2.User{})
	db.AutoMigrate(&models2.Problem{})
	db.AutoMigrate(&models2.Contest{})
	db.AutoMigrate(&models2.VerifyEmail{})
}
