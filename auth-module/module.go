package main

import (
	"context"

	"github.com/cesc1802/auth-module/feature"
	"github.com/cesc1802/share-module/system"
)

type Module struct{}

func (Module) Startup(ctx context.Context, mono system.Service) (err error) {
	return Root(ctx, mono)
}

func Root(ctx context.Context, mono system.Service) error {
	feature.RegisterHandlerV1(mono)
	return nil
}
