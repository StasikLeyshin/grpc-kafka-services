package server

import (
	"context"
	"github.com/StasikLeyshin/grpc-kafka-services/internal/service/server/converter"
	desc "github.com/StasikLeyshin/grpc-kafka-services/pkg/server_v1"
)

func (i *Implementation) CreateServer(ctx context.Context, serverRequest *desc.CreateServerRequest) (*desc.CreateServerResponse, error) {
	result, err := i.serverService.CreateServer(ctx, converter.ToCreateServerRequestFromGrpc(serverRequest))
	if err != nil {
		return nil, err
	}

	return converter.ToCreateServerResponseToGrpc(result), nil
}