package kafka

import (
	"context"
	"github.com/StasikLeyshin/grpc-kafka-services/internal/service"
	"github.com/StasikLeyshin/grpc-kafka-services/internal/service/server-consumer/converter"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

//type ConsumerConfig struct {
//	BufferSize uint `yaml:"buffer_size"`
//}

type Consumer struct {
	client *kafka.Reader
	logger *logrus.Logger

	cancel context.CancelFunc
	done   chan struct{}

	values chan *kafka.Message

	serverService service.ServerConsumerService
}

func NewConsumer(client *kafka.Reader, serverService service.ServerConsumerService, logger *logrus.Logger) *Consumer {
	return &Consumer{
		client:        client,
		serverService: serverService,
		logger:        logger,
	}
}

func (c *Consumer) Start() error {
	ctx, cancel := context.WithCancel(context.Background())

	c.cancel = cancel
	c.done = make(chan struct{})

	c.values = make(chan *kafka.Message, 100)

	go c.run(ctx)

	return nil
}

func (c *Consumer) Stop(ctx context.Context) error {
	c.cancel()

	select {
	case <-c.done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (c *Consumer) run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			close(c.values)
			close(c.done)

		default:
			msg, err := c.client.ReadMessage(ctx)
			if err != nil {
				c.logger.WithError(err).Error("failed to read message from kafka")
				continue
			}
			err = c.serverService.CreateServer(ctx, converter.ToServerFromKafka(msg.Value), string(msg.Key))
			c.logger.WithError(err).Error(err)
			if err != nil {
				continue
			}
			c.values <- &kafka.Message{
				Key:   msg.Key,
				Value: msg.Value,
			}
		}
	}
}

func (c *Consumer) GetValueChan() <-chan *kafka.Message {
	return c.values
}
