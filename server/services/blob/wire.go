//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/xince-fun/InstaGo/server/services/blob/pkg/initialize"
)

func InitializeService() (*BlobServiceImpl, error) {
	wire.Build(
		BlobServiceImplSet,
		initialize.InitMinio,
		initialize.InitMQ,
		initialize.InitDB,
	)

	return new(BlobServiceImpl), nil
}
