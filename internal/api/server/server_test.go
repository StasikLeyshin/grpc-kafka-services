package server

import (
	"context"
	desc "github.com/StasikLeyshin/grpc-kafka-services/pkg/server_v1"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"testing"
)

func Test_CreateServer(t *testing.T) {

	mock := gomock.NewController(t)
	defer mock.Finish()

	const address = "localhost:8000"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := desc.NewManagerServiceClient(conn)

	tests := []struct {
		name     string
		request  *desc.CreateServerRequest
		response *desc.CreateServerResponse
		errCode  codes.Code
		errMsg   string
	}{
		{
			name: "nil response",
			request: &desc.CreateServerRequest{
				Host: "",
				Port: "",
				Name: "",
			},
			response: &desc.CreateServerResponse{
				Status: "ok",
			},
		},
		{
			name: "normal response",
			request: &desc.CreateServerRequest{
				Host: "localhost",
				Port: "3312",
				Name: "response",
			},
			response: &desc.CreateServerResponse{
				Status: "ok",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, err := c.CreateServer(context.Background(), tt.request)

			if err != nil {
				t.Fatalf("error creating the server: %v", err)
			}
			assert.Equal(t, response.Status, tt.response.Status)
		})
	}
}
