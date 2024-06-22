//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/xince-fun/InstaGo/server/services/user/pkg/initialize"
)

func InitializeService() *UserServiceImpl {
	wire.Build(
		UserServiceImplSet,
		initialize.InitBlob,
		initialize.InitDB,
	)

	return new(UserServiceImpl)
}
