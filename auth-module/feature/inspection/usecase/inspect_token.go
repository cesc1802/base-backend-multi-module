package usecase

import (
	"context"
	"errors"

	"auth-module/feature/inspection/domain/entity"
	"share-module/common"
)

type UserStore interface {
	Find(ctx context.Context, loginID string) (*entity.User, error)
}

type TokenProvider interface {
	Extract(token string) (common.Requester, error)
}

type InspectTokenUseCase struct {
	store         UserStore
	tokenProvider TokenProvider
}

func NewInspectToken(store UserStore, tokenProvider TokenProvider) *InspectTokenUseCase {
	return &InspectTokenUseCase{
		store:         store,
		tokenProvider: tokenProvider,
	}
}

func (uc *InspectTokenUseCase) Register(ctx context.Context, token string) error {
	requester, err := uc.tokenProvider.Extract(token)
	if err != nil {
		return errors.New("extract token error")
	}

	user, err := uc.store.Find(ctx, requester.UserID())
	if err != nil {
		return errors.New("cannot find user")
	}

	if !user.IsActive() {
		return errors.New("user has band")
	}

	return nil
}
