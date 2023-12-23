package grpc

import (
	"github.com/StasikLeyshin/grpc-kafka-services/internal/api/server"
	desc "github.com/StasikLeyshin/grpc-kafka-services/pkg/server_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"net"
)

type Config struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type ServerGRPC struct {
	grpcServer           *grpc.Server
	implementationServer *server.Implementation
	address              string
}

func NewServerGRPC(config Config, implementationServer *server.Implementation) *ServerGRPC {
	address := net.JoinHostPort(config.Host, config.Port)

	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	reflection.Register(grpcServer)

	desc.RegisterManagerServiceServer(grpcServer, implementationServer)

	return &ServerGRPC{
		grpcServer:           grpcServer,
		implementationServer: implementationServer,
		address:              address,
	}
}

func (s *ServerGRPC) Start() error {
	list, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}

	err = s.grpcServer.Serve(list)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServerGRPC) Stop() error {
	s.grpcServer.GracefulStop()
	return nil
}
