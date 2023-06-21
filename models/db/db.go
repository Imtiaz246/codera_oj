package db

import (
	"fmt"
	"github.com/imtiaz246/codera_oj/custom/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const (
	ErrUnsupportedDatabase = "unsupported database selected : %v"
)

var (
	// db instance holds the connection of the database.
	// Supported databases are : sqlite3, postgres, mysql
	db *gorm.DB
)

// init initializes the database connection
// and assign it to the db instance.
func init() {
	dbConfig := config.Cfg.Database

	switch dbConfig.DbType {
	case "sqlite3":
		if err := connSqliteDB(); err != nil {
			panic(err)
		}
	case "postgres":
		if err := connPostgresDB(); err != nil {
			panic(err)
		}
	case "mysql":
		if err := connMysqlDB(); err != nil {
			panic(err)
		}
	default:
		panic(ErrUnsupportedDatabase)
	}

	db.NamingStrategy = schema.NamingStrategy{
		TablePrefix: "codera",
	}
}

func GetEngine() *gorm.DB {
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

// MigrateModelTables apply auto migration for models
func MigrateModelTables(model interface{}) error {
	if err := db.AutoMigrate(model); err != nil {
		return err
	}
	return nil
}
