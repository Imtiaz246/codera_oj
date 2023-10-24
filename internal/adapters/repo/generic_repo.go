package repo

import (
	"github.com/imtiaz246/codera_oj/internal/adapters/repo/db"
	"github.com/imtiaz246/codera_oj/internal/core/domain/models"
	"gorm.io/gorm"
)

type GenericRepo[T models.ModelFactory] struct {
	*db.Database
}

func NewGenericRepo[T models.ModelFactory](db *db.Database) GenericInterface[T] {
	return &GenericRepo[T]{
		db,
	}
}

func (g *GenericRepo[T]) GetAllRecords(t T) ([]T, error) {
	records := make([]T, 0)
	if err := g.DB.Find(&t).Error; err != nil {
		return nil, err
	}
	return records, nil
}

func (g *GenericRepo[T]) GetRecordByModel(t T, preloads ...string) error {
	return addPreloads(g.DB, preloads).First(&t).Error
}

func (g *GenericRepo[T]) GetRecordByID(id int64) (T, error) {
	var t T
	err := g.DB.Where("id = ?", id).First(&t).Error
	return t, err
}

func (g *GenericRepo[T]) CreateRecord(t T) error {
	return g.DB.Create(&t).Error
}

func (g *GenericRepo[T]) UpdateRecord(t T) error {
	return g.DB.Save(&t).Error
}

func (g *GenericRepo[T]) DeleteRecord(t T) error {
	return g.DB.Delete(&t).Error
}

func (g *GenericRepo[T]) DeleteRecordByID(id int64) (T, error) {
	var t T
	if err := g.DB.Where("id = ?", id).Delete(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (g *GenericRepo[T]) GetRecordByExpression(query any, args ...any) (T, error) {
	var t T
	if err := g.DB.Where(query, args...).First(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func addPreloads(db *gorm.DB, preloads []string) *gorm.DB {
	for _, t := range preloads {
		db = db.Preload(t)
	}
	return db
}
