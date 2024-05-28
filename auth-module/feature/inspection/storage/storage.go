package storage

import (
	"gorm.io/gorm"
)

type SQLStorage struct {
	db *gorm.DB
}

func NewSqlStorage(db *gorm.DB) *SQLStorage {
	return &SQLStorage{
		db: db,
	}
}
