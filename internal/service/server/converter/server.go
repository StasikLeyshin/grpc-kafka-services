package converter

import (
	"github.com/StasikLeyshin/grpc-kafka-services/internal/service/server/model"
	desc "github.com/StasikLeyshin/grpc-kafka-services/pkg/server_v1"
)

func ToServerFromGrpc(server *desc.Server) *model.Server {
	return &model.Server{
		UUID:      server.Uuid,
		Name:      server.Name,
		Host:      server.Host,
		Port:      server.Port,
		CreatedAt: server.CreatedAt.AsTime(),
	}
}

func ToCreateServerRequestFromGrpc(server *desc.CreateServerRequest) *model.CreateServerRequest {
	return &model.CreateServerRequest{
		Name: server.Name,
		Host: server.Host,
		Port: server.Port,
	}
}

func ToCreateServerResponseToGrpc(server *model.CreateServerResponse) *desc.CreateServerResponse {
	return &desc.CreateServerResponse{
		Status: server.Status,
	}
}
