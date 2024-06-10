package tokprovider

import (
	"context"

	"github.com/cesc1802/share-module/common"
)

type AppToken struct {
	AccessToken  string
	RefreshToken string
}

type AppPayload struct {
	UserID string
}

func (pl AppPayload) GetUserID() string {
	return pl.UserID
}

type TokenExtractor interface {
	Extract(token string) (common.Requester, error)
}

type TokenGenerator interface {
	Generate(ctx context.Context, payload *AppPayload) (*AppToken, error)
}
type TokenProvider interface {
	TokenExtractor
	TokenGenerator
}
