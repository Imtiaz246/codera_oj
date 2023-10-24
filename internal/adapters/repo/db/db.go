package db

import (
	"fmt"
	"github.com/imtiaz246/codera_oj/internal/core/domain/dto"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Database struct {
	*gorm.DB
}

func NewDB(dbConfig *dto.DatabaseConfig) (*Database, error) {
	db := new(Database)
	var err error
	switch dbConfig.DbType {
	case "sqlite3":
		db.DB, err = connSqliteDB()
	case "postgres":
		db.DB, err = connPostgresDB()
	case "mysql":
		db.DB, err = connMysqlDB()
	default:
		return nil, fmt.Errorf("databast type `%v` is not supported", dbConfig.DbType)
	}

	if err != nil {
		return nil, err
	}

	db.NamingStrategy = schema.NamingStrategy{
		TablePrefix: "codera_",
	}

	return db, nil
}

func (d *Database) GetEngine() *gorm.DB {
	return d.DB
}

// connPostgresDB connects to postgres and assigns to db instance
func connPostgresDB() (*gorm.DB, error) {
	fmt.Println("postgres db")
	return nil, nil
}

// connMysqlDB connects to mysql and assigns to db instance
func connMysqlDB() (*gorm.DB, error) {
	fmt.Println("mysql db")
	return nil, nil
}

// connSqliteDB connects to sqlite3 and assigns to db instance
func connSqliteDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("gorm.sqlite"), &gorm.Config{})
	return db, err
}
