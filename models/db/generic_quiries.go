package db

import (
	"fmt"
	"reflect"
)

func GetAllRecords[T any]() ([]*T, error) {
	t := make([]*T, 0)
	if err := db.Find(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func GetRecordByID[T any](ID uint) (*T, error) {
	t := new(T)
	if err := db.Where("id", ID).First(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func CreateRecord(t any) error {
	if reflect.ValueOf(t).Kind() != reflect.Ptr {
		return fmt.Errorf("type has to be a pointer")
	}
	return db.Create(t).Error
}

func UpdateRecord(t any) error {
	if reflect.ValueOf(t).Kind() != reflect.Ptr {
		return fmt.Errorf("type has to be a pointer")
	}
	return db.Save(t).Error
}

func DeleteRecord(t any) error {
	if reflect.ValueOf(t).Kind() != reflect.Ptr {
		return fmt.Errorf("type has to be a pointer")
	}
	return db.Delete(t).Error
}

func DeleteRecordByID[T any](ID uint) (*T, error) {
	t := new(T)
	if err := db.Where("id = ?", ID).Delete(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func GetRecordByExp[T any](query any, args ...any) (*T, error) {
	t := new(T)
	if err := db.Where(query, args...).First(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}
