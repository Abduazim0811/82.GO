package rabbitmq

import (
	"github.com/rabbitmq/amqp091-go"
)

type Producer struct {
	channel *amqp091.Channel
}

func NewProducer() (*Producer, error) {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	err = ch.ExchangeDeclare(
		"message_topic",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &Producer{channel: ch}, nil
}

func (p *Producer) PublishMessage(routingKey string, body []byte) error {
	err := p.channel.Publish(
		"message_topic",
		routingKey,
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	)
	return err
}
