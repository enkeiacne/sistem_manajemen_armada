package configEnviroment

import (
	"github.com/joho/godotenv"
	"os"
)

var (
	EnvironmentAppPort                  = getEnv("APP_PORT")
	EnvironmentAppMode                  = getEnv("APP_MODE")
	EnvironmentDatabaseUser             = getEnv("DATABASE_USER")
	EnvironmentDatabasePassword         = getEnv("DATABASE_PASSWORD")
	EnvironmentDatabaseName             = getEnv("DATABASE_NAME")
	EnvironmentDatabaseHost             = getEnv("DATABASE_HOST")
	EnvironmentDatabasePort             = getEnv("DATABASE_PORT")
	EnvironmentDatabaseSSLMode          = getEnv("DATABASE_SSL_MODE")
	EnvironmentRabbitMQHost             = getEnv("RABBITMQ_HOST")
	EnvironmentRabbitMQPort             = getEnv("RABBITMQ_PORT")
	EnvironmentRabbitMQUser             = getEnv("RABBITMQ_USER")
	EnvironmentRabbitMQPassword         = getEnv("RABBITMQ_PASSWORD")
	EnvironmentMQTTHost                 = getEnv("MQTT_HOST")
	EnvironmentMQTTPort                 = getEnv("MQTT_PORT")
	EnvironmentMQTTClientID             = getEnv("MQTT_CLIENT_ID")
	EnvironmentGeofenceCurrentLatitude  = getEnv("GEOFENCE_CURRENT_LATITUDE")
	EnvironmentGeofenceCurrentLongitude = getEnv("GEOFENCE_CURRENT_LONGITUDE")
	EnvironmentGeofenceRadius           = getEnv("GEOFENCE_RADIUS")
	EnvironmentGeofenceTargetLatitude   = getEnv("GEOFENCE_TARGET_LATITUDE")
	EnvironmentGeofenceTargetLongitude  = getEnv("GEOFENCE_TARGET_LONGITUDE")
)

func getEnv(key string) string {
	LoadEnv()
	value := os.Getenv(key)
	if value == "" {
		panic("Environment variable '" + key + "' not set")
	}
	return value
}
func LoadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		panic("Warning: No .env file found")
	}
}
