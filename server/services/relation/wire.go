//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/xince-fun/InstaGo/server/services/relation/pkg/initialize"
)

func InitializeService() (*RelationServiceImpl, error) {
	wire.Build(
		RelationServiceImplSet,
		initialize.InitUser,
		initialize.InitDB,
	)

	return new(RelationServiceImpl), nil
}
