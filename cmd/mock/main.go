package main

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"math"
	"strconv"
	"time"
	configEnviroment "transakarta_BE_test/internal/config/enviroment"
)

func main() {
	opts := mqtt.NewClientOptions().AddBroker("tcp://" + configEnviroment.EnvironmentMQTTHost + ":" + configEnviroment.EnvironmentMQTTPort)
	opts.SetClientID("mock-publisher")
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	vehicleID := "B1234XYZ"
	log.Println("mock publisher is running...")
	currentLatStr := configEnviroment.EnvironmentGeofenceCurrentLatitude
	currentLonStr := configEnviroment.EnvironmentGeofenceCurrentLongitude

	targetLatStr := configEnviroment.EnvironmentGeofenceTargetLatitude
	targetLonStr := configEnviroment.EnvironmentGeofenceTargetLongitude

	currentLat, err := strconv.ParseFloat(currentLatStr, 64)
	if err != nil {
		log.Fatalf("invalid current latitude: %v", err)
	}
	currentLon, err := strconv.ParseFloat(currentLonStr, 64)
	if err != nil {
		log.Fatalf("invalid current longitude: %v", err)
	}
	targetLat, err := strconv.ParseFloat(targetLatStr, 64)
	if err != nil {
		log.Fatalf("invalid target latitude: %v", err)
	}
	targetLon, err := strconv.ParseFloat(targetLonStr, 64)
	if err != nil {
		log.Fatalf("invalid target longitude: %v", err)
	}
	step := 0.005

	for {
		deltaLat := targetLat - currentLat
		deltaLon := targetLon - currentLon

		//if math.Abs(deltaLat) < 0.00001 && math.Abs(deltaLon) < 0.00001 {
		//	log.Println("Reached target.")
		//	break
		//}

		distance := math.Sqrt(deltaLat*deltaLat + deltaLon*deltaLon)
		moveLat := (deltaLat / distance) * step
		moveLon := (deltaLon / distance) * step
		currentLat += moveLat
		currentLon += moveLon

		// Publish data
		data := map[string]interface{}{
			"latitude":  currentLat,
			"longitude": currentLon,
			"timestamp": time.Now().Unix(),
		}
		payload, _ := json.Marshal(data)
		topic := fmt.Sprintf("/fleet/vehicle/%s/location", vehicleID)
		client.Publish(topic, 0, false, payload)

		log.Printf("Mock sent lat=%.6f, lon=%.6f", currentLat, currentLon)
		time.Sleep(2 * time.Second)
	}
}
