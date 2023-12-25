package main

import (
	"context"
	"github.com/StasikLeyshin/grpc-kafka-services/internal/app"
	"github.com/StasikLeyshin/grpc-kafka-services/internal/app/startup"
	"github.com/StasikLeyshin/grpc-kafka-services/internal/repository/server"
	"github.com/StasikLeyshin/grpc-kafka-services/internal/service/kafka"
	server_consumer "github.com/StasikLeyshin/grpc-kafka-services/internal/service/server-consumer"
	"os"
)

func main() {

	// Файл с конфигурацией проекта
	configPath := os.Getenv("CONFIG_PATH")

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
		logger.Fatalf("failed to connect to database: %v", err)
	}

	logger.Info("connecting to the database: success")

	defer func() {
		dbInstance, _ := dbClient.DB()
		if err = dbInstance.Close(); err != nil {
			logger.WithError(err).Warn("failed to close database")
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
		logger.Fatalf("failed to auto migrate: %v", err)
	}

	kafkaProducer := kafka.NewConsumer(kafkaClient, serviceClient, logger)

	// Запускаем компонент consumer kafka
	app.NewApp(logger, kafkaProducer).Run(context.Background())
}
