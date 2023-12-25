package server_consumer

import (
	"context"
	"github.com/StasikLeyshin/grpc-kafka-services/internal/models"
	"github.com/StasikLeyshin/grpc-kafka-services/internal/repository"
	mock_repository "github.com/StasikLeyshin/grpc-kafka-services/internal/repository/mocks"
	"github.com/golang/mock/gomock"
	"testing"
)

func Test_CreateServer(t *testing.T) {
	ctx := context.Background()

	controller := gomock.NewController(t)

	type args struct {
		server *models.Server
		uuid   string
	}

	tests := []struct {
		name       string
		repository repository.ServerRepository
		input      args
	}{
		{
			name:       "nil response",
			repository: simulateRepository(controller, nil),
			input: args{
				server: nil,
				uuid:   "",
			},
		},
		{
			name:       "normal response",
			repository: simulateRepository(controller, nil),
			input: args{
				server: &models.Server{
					Host: "",
					Port: "",
					Name: "",
				},
				uuid: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewService(tt.repository)
			err := s.CreateServer(ctx, tt.input.server, tt.input.uuid)
			if err != nil {
				t.Errorf("Error: %v", err)
			}
		})
	}
}

func simulateRepository(controller *gomock.Controller, err error) repository.ServerRepository {
	mockService := mock_repository.NewMockServerRepository(controller)

	mockService.EXPECT().CreateServer(gomock.Any(), gomock.Any()).Return(err).AnyTimes()

	return mockService
}
