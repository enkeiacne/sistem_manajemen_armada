package rabbitmqProducer

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"transakarta_BE_test/internal/config/rabbitmq"
)

func PublishMessage(exchange, queue string, body []byte) error {
	ch, err := rabbitmq.ConnMq.Channel()
	if err != nil {
		return err
	}
	err = ch.ExchangeDeclare("fleet.events", "fanout", true, false, false, false, nil)
	if err != nil {
		log.Fatal("Exchange error:", err)
	}
	err = ch.Publish(
		exchange,
		queue,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		log.Printf("Failed to publish message: %s", err)
	}
	return err
}
