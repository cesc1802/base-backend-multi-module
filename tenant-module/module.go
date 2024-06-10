package tenant_module

import (
	"context"

	"github.com/cesc1802/share-module/system"
)

type Module struct{}

func (Module) Startup(ctx context.Context, mono system.Service) (err error) {
	return Root(ctx, mono)
}

func Root(ctx context.Context, mono system.Service) error {
	return nil
}
