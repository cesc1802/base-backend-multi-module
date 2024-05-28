package storage

import (
	"context"

	"auth-module/feature/authentication/domain/entity"
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

func (s *SQLStorage) Find(ctx context.Context, loginID string) (*entity.User, error) {
	return nil, nil
}

func (s *SQLStorage) Create(ctx context.Context, cmd *entity.User) error {
	return nil
}
