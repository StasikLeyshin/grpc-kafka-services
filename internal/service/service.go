package service

import (
	"context"
	"github.com/StasikLeyshin/grpc-kafka-services/internal/service/server/model"
)

type ServerService interface {
	CreateServer(ctx context.Context, server *model.CreateServerRequest) (*model.CreateServerResponse, error)
}
