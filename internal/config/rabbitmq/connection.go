package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
	configEnviroment "transakarta_BE_test/internal/config/enviroment"
)

var ConnMq *amqp.Connection

func ConnectRabbitMQ() (*amqp.Connection, error) {
	var conn *amqp.Connection
	var err error

	url := "amqp://" +
		configEnviroment.EnvironmentRabbitMQUser + ":" +
		configEnviroment.EnvironmentRabbitMQPassword + "@" +
		configEnviroment.EnvironmentRabbitMQHost + ":" +
		configEnviroment.EnvironmentRabbitMQPort + "/"

	maxRetries := 3
	for i := 0; i < maxRetries; i++ {
		conn, err = amqp.Dial(url)
		if err == nil {
			log.Println("✅ Successfully connected to RabbitMQ")
			ConnMq = conn
			return conn, nil
		}

		log.Printf("🔁 Retry connecting to RabbitMQ (%d/%d): %v", i+1, maxRetries, err)
		time.Sleep(5 * time.Second)
	}

	log.Println("❌ Failed to connect to RabbitMQ after retries")
	return nil, err
}
