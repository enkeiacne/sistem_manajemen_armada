package main

import (
	"github.com/gin-gonic/gin"
	commonsValidator "transakarta_BE_test/internal/commons/validator"
	configEnviroment "transakarta_BE_test/internal/config/enviroment"
	"transakarta_BE_test/internal/config/mqtt"
	mqttSubscriber "transakarta_BE_test/internal/config/mqtt/subscriber"
	"transakarta_BE_test/internal/config/rabbitmq"
	"transakarta_BE_test/internal/database"
	"transakarta_BE_test/internal/routes"
)

func main() {
	configEnviroment.LoadEnv()

	// Connect to the database
	database.DatabaseConnect()

	// Connect to RabbitMQ
	_, err := rabbitmq.ConnectRabbitMQ()
	if err != nil {
		panic("Rabbitmq cant connected" + err.Error())
	}

	// Connect to MQTT
	err = mqtt.ConnectMQTT()
	if err != nil {
		panic("mqtt cant connected" + err.Error())
	}
	mqttSubscriber.SubscribeVehicleLocation()

	router := gin.Default()
	if configEnviroment.EnvironmentAppMode == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	commonsValidator.RegisterCustomValidators(router)

	// router
	routes.SetupRoutes(router)

	err = router.Run(":" + configEnviroment.EnvironmentAppPort)
	if err != nil {
		panic(err)
	}
}
