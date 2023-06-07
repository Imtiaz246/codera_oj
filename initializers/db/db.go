package db

import (
	"fmt"
	"github.com/imtiaz246/codera_oj/app/models"
	"github.com/imtiaz246/codera_oj/initializers/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ErrUnsupportedDatabase indicates error for unsupported database
const ErrUnsupportedDatabase = "unsupported database selected : %v"

var (
	// db instance holds the connection of the database.
	// Supported databases are : sqlite3, postgres, mysql
	db *gorm.DB
)

// InitializeDB initializes the database connection
// and assign it to the db instance.
func InitializeDB() error {
	dbConfig := config.GetDBConfig()

	switch dbConfig.DbType {
	case "sqlite3":
		if err := connSqliteDB(); err != nil {
			return err
		}
	case "postgres":
		if err := connPostgresDB(); err != nil {
			return err
		}
	case "mysql":
		if err := connMysqlDB(); err != nil {
			return err
		}
	default:
		return fmt.Errorf(ErrUnsupportedDatabase, dbConfig.DbType)
	}

	if err := doAutoMigrations(db); err != nil {
		return err
	}
	applySeeds()

	return nil
}

// GetDB returns the db instance
func GetDB() *gorm.DB {
	return db
}

// connPostgresDB connects to postgres and assigns to db instance
func connPostgresDB() (err error) {
	fmt.Println("postgres db")
	return
}

// connMysqlDB connects to mysql and assigns to db instance
func connMysqlDB() (err error) {
	fmt.Println("mysql db")
	return
}

// connSqliteDB connects to sqlite3 and assigns to db instance
func connSqliteDB() (err error) {
	db, err = gorm.Open(sqlite.Open("gorm.sqlite"), &gorm.Config{})
	return
}

// doAutoMigrations apply auto migration for models
func doAutoMigrations(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Problem{},
		&models.Contest{},
		&models.VerifyEmail{},
		&models.Sessions{})
}

// IsDBInitialized checks if the database is already initialized or not
func IsDBInitialized() bool {
	return db != nil
}
