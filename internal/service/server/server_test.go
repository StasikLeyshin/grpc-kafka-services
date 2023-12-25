package server

import (
	"context"
	"github.com/StasikLeyshin/grpc-kafka-services/internal/service/kafka"
	mock_kafka "github.com/StasikLeyshin/grpc-kafka-services/internal/service/kafka/mocks"
	"github.com/StasikLeyshin/grpc-kafka-services/internal/service/server/model"
	"github.com/golang/mock/gomock"
	"testing"
)

func Test_CreateServer(t *testing.T) {
	ctx := context.Background()

	controller := gomock.NewController(t)

	type args struct {
		createServerRequest *model.CreateServerRequest
	}

	tests := []struct {
		name          string
		kafkaProducer kafka.KafkaProducer
		input         args
	}{
		{
			name:          "nil request",
			kafkaProducer: simulateKafkaProducer(ctx, controller, nil),
			input: args{
				createServerRequest: nil,
			},
		},
		{
			name:          "empty request",
			kafkaProducer: simulateKafkaProducer(ctx, controller, nil),
			input: args{
				createServerRequest: &model.CreateServerRequest{
					Host: "",
					Port: "",
					Name: "",
				},
			},
		},
		{
			name:          "normal request",
			kafkaProducer: simulateKafkaProducer(ctx, controller, nil),
			input: args{
				createServerRequest: &model.CreateServerRequest{
					Host: "localhost",
					Port: "3245",
					Name: "test",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewService(tt.kafkaProducer)
			_, err := s.CreateServer(ctx, tt.input.createServerRequest)
			if err != nil {
				t.Errorf("Error: %v", err)
			}
		})
	}
}

func simulateKafkaProducer(ctx context.Context, controller *gomock.Controller, err error) kafka.KafkaProducer {
	mockService := mock_kafka.NewMockKafkaProducer(controller)

	mockService.EXPECT().SendMessage(ctx, gomock.Any(), gomock.Any()).Return(err).AnyTimes()

	return mockService
}
