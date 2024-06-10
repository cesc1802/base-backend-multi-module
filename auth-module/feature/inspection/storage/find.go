package storage

import (
	"context"

	"github.com/cesc1802/auth-module/feature/inspection/domain/entity"
)

func (s *SQLStorage) Find(ctx context.Context, loginID string) (*entity.User, error) {
	return nil, nil
}
