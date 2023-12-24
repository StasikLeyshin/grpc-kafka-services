package kafka

import (
	"context"
	"github.com/StasikLeyshin/grpc-kafka-services/internal/service"
	"github.com/StasikLeyshin/grpc-kafka-services/internal/service/server-consumer/converter"
	"github.com/segmentio/kafka-go"
)

//type ConsumerConfig struct {
//	BufferSize uint `yaml:"buffer_size"`
//}

type Consumer struct {
	client *kafka.Reader
	//config ConsumerConfig
	//logger log.Logger

	cancel context.CancelFunc
	done   chan struct{}

	values chan *kafka.Message

	serverService service.ServerConsumerService
}

func NewConsumer(client *kafka.Reader, serverService service.ServerConsumerService) *Consumer {
	return &Consumer{
		client:        client,
		serverService: serverService,
		//logger: logger.With(log.ComponentKey, "Kafka consumer"),
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
				continue
			}
			err = c.serverService.CreateServer(ctx, converter.ToServerFromKafka(msg.Value), string(msg.Key))
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
