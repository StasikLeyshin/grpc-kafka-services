package service

import (
	"context"
	"github.com/StasikLeyshin/grpc-kafka-services/internal/models"
	"github.com/StasikLeyshin/grpc-kafka-services/internal/service/server/model"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type ServerService interface {
	CreateServer(ctx context.Context, server *model.CreateServerRequest) (*model.CreateServerResponse, error)
}

type ServerConsumerService interface {
	CreateServer(ctx context.Context, server *models.Server, userUUID string) error
}
