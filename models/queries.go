package models

import (
	"github.com/imtiaz246/codera_oj/models/db"
)

type ST interface {
	*User |
		*Session |
		*Problem |
		*Contest |
		*Tag |
		*VerifyEmail |
		*ProblemTag |
		*Dataset |
		*ProblemShare |
		*ProblemChangeLog |
		*ProblemSolution |
		*ProblemDiscussion
}

func GetAllRecords[T ST]() ([]T, error) {
	t := make([]T, 0)
	if err := db.GetEngine().Find(&t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func GetRecordByID[T ST](ID string) (T, error) {
	var t T
	err := db.GetEngine().Where("id = ?", ID).First(&t).Error
	return t, err
}

func CreateRecord[T ST](t T) error {
	return db.GetEngine().Create(t).Error
}

func UpdateRecord[T ST](t T) error {
	return db.GetEngine().Save(t).Error
}

func DeleteRecord[T ST](t T) error {
	return db.GetEngine().Delete(t).Error
}

func DeleteRecordByID[T ST](ID uint) (T, error) {
	var t T
	if err := db.GetEngine().Where("id = ?", ID).Delete(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func GetRecordByExp[T ST](query any, args ...any) (T, error) {
	var t T
	if err := db.GetEngine().Where(query, args...).First(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}
