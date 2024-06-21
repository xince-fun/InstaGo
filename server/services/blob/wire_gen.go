// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/xince-fun/InstaGo/server/services/blob/app"
	"github.com/xince-fun/InstaGo/server/services/blob/infra/mq/amqp"
	"github.com/xince-fun/InstaGo/server/services/blob/infra/object/minio"
	"github.com/xince-fun/InstaGo/server/services/blob/infra/persistence"
	"github.com/xince-fun/InstaGo/server/services/blob/infra/persistence/dal"
	"github.com/xince-fun/InstaGo/server/services/blob/pkg/initialize"
)

// Injectors from wire.go:

func InitializeService() (*BlobServiceImpl, error) {
	db := initialize.InitDB()
	blobDal := dal.NewBlobDal(db)
	connection := initialize.InitMQ()
	pExchange := amqp.ProvidePExchange()
	publisher, err := amqp.NewPublisher(connection, pExchange)
	if err != nil {
		return nil, err
	}
	eventPublisher := amqp.NewEventPublisher(publisher)
	blobRepo := persistence.NewBlobRepo(blobDal, eventPublisher)
	client := initialize.InitMinio()
	minioBucketManager := minio.NewMinioBucketManager(client)
	blobApplicationService := app.NewBlobApplicationService(blobRepo, minioBucketManager)
	sExchange := amqp.ProvideSExchange()
	subscriber, err := amqp.NewSubscriber(connection, sExchange)
	if err != nil {
		return nil, err
	}
	eventSubscriber := amqp.NewEventSubscriber(subscriber)
	blobEventListener := app.NewBlobEventListener(blobRepo, eventSubscriber)
	blobServiceImpl := NewBlobServiceImpl(blobApplicationService, blobEventListener)
	return blobServiceImpl, nil
}
