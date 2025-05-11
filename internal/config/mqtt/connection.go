package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	configEnviroment "transakarta_BE_test/internal/config/enviroment"
)

var ClientMQTT mqtt.Client

func ConnectMQTT() error {
	opts := mqtt.NewClientOptions().AddBroker("tcp://" + configEnviroment.EnvironmentMQTTHost + ":" + configEnviroment.EnvironmentMQTTPort).
		SetClientID(configEnviroment.EnvironmentMQTTClientID)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	log.Println("Successfully Connected to MQTT Broker")
	ClientMQTT = client
	return nil
}
