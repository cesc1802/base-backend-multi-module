package usecase

import (
	"auth-module/feature/authentication/domain/entity"
	"auth-module/feature/authentication/dto"

	"context"
	"errors"
)

type UserStore interface {
	Find(ctx context.Context, loginID string) (*entity.User, error)
	Create(ctx context.Context, cmd *entity.User) error
}

type RegisterUserUseCase struct {
	store UserStore
}

func NewRegisterUserUseCase(store UserStore) *RegisterUserUseCase {
	return &RegisterUserUseCase{
		store: store,
	}
}

func (uc *RegisterUserUseCase) Register(ctx context.Context, cmd dto.RegisterRequest) error {
	user, err := uc.store.Find(ctx, cmd.LoginID)

	if err != nil {
		return errors.New("cannot find user")
	}

	if user != nil {
		return errors.New("user existing")
	}

	if err := uc.store.Create(ctx, &entity.User{
		LoginID:  cmd.LoginID,
		Password: cmd.Password,
	}); err != nil {
		return errors.New("cannot create user")
	}

	return nil
}
