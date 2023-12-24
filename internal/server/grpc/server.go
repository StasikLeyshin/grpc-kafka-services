package grpc

import (
	"context"
	"fmt"
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
	port                 int
}

func NewServerGRPC(port int, implementationServer *server.Implementation) *ServerGRPC {
	//address := net.JoinHostPort(host, port)

	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	reflection.Register(grpcServer)

	desc.RegisterManagerServiceServer(grpcServer, implementationServer)

	return &ServerGRPC{
		grpcServer:           grpcServer,
		implementationServer: implementationServer,
		port:                 port,
	}
}

func (s *ServerGRPC) Start() error {
	fmt.Println("Start ServerGRPC", s.port)
	list, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return err
	}
	go func() {
		err = s.grpcServer.Serve(list)
		fmt.Println("Serve", err)
		if err != nil {
			//return err
		}
	}()
	fmt.Println("Exit")
	return nil
}

func (s *ServerGRPC) Stop(ctx context.Context) error {
	s.grpcServer.Stop()
	return nil
}
