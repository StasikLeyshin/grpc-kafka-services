package server

import (
	"context"
	"github.com/StasikLeyshin/grpc-kafka-services/internal/service/server/converter"
	desc "github.com/StasikLeyshin/grpc-kafka-services/pkg/server_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CreateServer(ctx context.Context, serverRequest *desc.CreateServerRequest) (*desc.CreateServerResponse, error) {
	if serverRequest == nil {
		return nil, status.Errorf(codes.Internal, "Internal error")
	}

	result, err := i.serverService.CreateServer(ctx, converter.ToCreateServerRequestFromGrpc(serverRequest))
	if err != nil {
		i.logger.WithError(err).Error(err)
		return nil, status.Errorf(codes.Internal, "Internal error")
	}

	return converter.ToCreateServerResponseToGrpc(result), nil
}
