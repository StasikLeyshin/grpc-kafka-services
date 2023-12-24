package main

import (
	"context"
	api "github.com/StasikLeyshin/grpc-kafka-services/internal/api/server"
	"github.com/StasikLeyshin/grpc-kafka-services/internal/app"
	"github.com/StasikLeyshin/grpc-kafka-services/internal/app/startup"
	"github.com/StasikLeyshin/grpc-kafka-services/internal/server/grpc"
	"github.com/StasikLeyshin/grpc-kafka-services/internal/service/kafka"
	"github.com/StasikLeyshin/grpc-kafka-services/internal/service/server"
	"os"
)

func main() {

	// Файл с конфигурацией проекта
	configPath := os.Getenv("CONFIG_PATH")
	//configPath := "configs/producer-service.yaml"

	// Создаём логгер
	logger := startup.NewLogger()

	// Парсим файл конфигурации
	config, err := startup.NewConfig(configPath)
	if err != nil {
		logger.Fatalf("failed to Config: %v", err)
	}

	// Подключаемся к kafka в качестве Producer
	kafkaClient := startup.NewKafkaProducer(config.KafkaConfig)
	defer func() {
		if err := kafkaClient.Close(); err != nil {
			logger.WithError(err).Warn("failed to close kafka")
		}
	}()

	kafkaProducer := kafka.NewProducer(kafkaClient)

	// Клиент для реализации бизнес-логики
	serviceClient := server.NewService(kafkaProducer)

	// Релизация методов grpc
	implementationServer := api.NewImplementation(serviceClient)

	// Создаём экземпляр grpc сервера
	grpcClient := grpc.NewServerGRPC(config.GrpcConfig.Port, implementationServer)

	// Запускаем компонент grpc сервера
	app.NewApp(logger, grpcClient).Run(context.Background())

	//// Клиент для реализации бизнес-логики
	//client := service.NewService(logger, config.Store.FilePath)
	//err = client.Start()
	//if err != nil {
	//	logger.Fatalf("failed Start Service: %v", err)
	//}
	//
	//// Создаём экземпляр http сервера
	//httpRouter := http.NewHttpRouter(config.Http, client, logger)
	//
	//// Запускаем http сервер
	//app.NewApp(logger, httpRouter).Run(context.Background())
}
