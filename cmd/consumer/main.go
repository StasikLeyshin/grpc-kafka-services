package main

import (
	"context"
	"github.com/StasikLeyshin/grpc-kafka-services/internal/app"
	"github.com/StasikLeyshin/grpc-kafka-services/internal/app/startup"
	"github.com/StasikLeyshin/grpc-kafka-services/internal/repository/server"
	"github.com/StasikLeyshin/grpc-kafka-services/internal/service/kafka"
	server_consumer "github.com/StasikLeyshin/grpc-kafka-services/internal/service/server-consumer"
	"log"
	"os"
)

func main() {

	// Файл с конфигурацией проекта
	configPath := os.Getenv("CONFIG_PATH")
	//configPath := "configs/config.yaml"

	// Создаём логгер
	logger := startup.NewLogger()

	// Парсим файл конфигурации
	config, err := startup.NewConfig(configPath)
	if err != nil {
		logger.Fatalf("failed to Config: %v", err)
	}

	// Открытие соединения с postgres
	dbClient, err := startup.DatabaseConnect(config.Database)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	log.Printf("connecting to the database: success")

	defer func() {
		dbInstance, _ := dbClient.DB()
		if err = dbInstance.Close(); err != nil {
			log.Fatalf("failed to close database: %v", err)
		}
	}()

	// Подключаемся к kafka в качестве Consumer
	kafkaClient := startup.NewKafkaConsumer(config.KafkaConfig)
	defer func() {
		if err := kafkaClient.Close(); err != nil {
			logger.WithError(err).Warn("failed to close kafka")
		}
	}()

	// Клиент для реализации бизнес-логики
	serviceClient := server_consumer.NewService(server.NewRepository(dbClient))

	err = serviceClient.Start()
	if err != nil {
		log.Fatalf("failed to auto migrate: %v", err)
	}

	kafkaProducer := kafka.NewConsumer(kafkaClient, serviceClient)

	//// Релизация методов grpc
	//implementationServer := api.NewImplementation(serviceClient)
	//
	//// Создаём экземпляр grpc сервера
	//grpcClient := grpc.NewServerGRPC(config.GrpcConfig, implementationServer)
	//
	//// Запускаем компонент kafka
	app.NewApp(logger, kafkaProducer).Run(context.Background())
}
