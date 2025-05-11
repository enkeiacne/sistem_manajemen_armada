package mqttSubscriber

import (
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"strconv"
	"strings"
	"time"
	commonsHelper "transakarta_BE_test/internal/commons/helper"
	configEnviroment "transakarta_BE_test/internal/config/enviroment"
	mqtt2 "transakarta_BE_test/internal/config/mqtt"
	rabbitmqProducer "transakarta_BE_test/internal/config/rabbitmq/producer"
	databaseEntities "transakarta_BE_test/internal/database/entities"
	moduleFrontOfficeVehicleLocationRepository "transakarta_BE_test/internal/modules/front-office/vehicle-location/repository"
)

type VehicleLocationPayload struct {
	VehicleID string  `json:"vehicle_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timestamp int64   `json:"timestamp"`
}

func SubscribeVehicleLocation() {
	mqtt2.ClientMQTT.Subscribe("/fleet/vehicle/+/location", 0, func(client mqtt.Client, msg mqtt.Message) {
		var payload VehicleLocationPayload
		if err := json.Unmarshal(msg.Payload(), &payload); err != nil {
			log.Println("Invalid payload:", err)
			return
		}

		// save to database
		data := databaseEntities.VehicleLocation{
			VehicleID: strings.Split(msg.Topic(), "/")[3],
			Latitude:  payload.Latitude,
			Longitude: payload.Longitude,
			Timestamp: time.Unix(payload.Timestamp, 0), // convert dari UNIX ke time.Time
		}

		repo := moduleFrontOfficeVehicleLocationRepository.NewVehicleLocationRepository()
		_, err := repo.Create(data)
		if err != nil {
			log.Println("Failed to save vehicle location:", err)
			return
		}

		// check distance
		latStr := configEnviroment.EnvironmentGeofenceTargetLatitude
		lonStr := configEnviroment.EnvironmentGeofenceTargetLongitude
		geofenceRadiusStr := configEnviroment.EnvironmentGeofenceRadius

		lat, err := strconv.ParseFloat(latStr, 64)
		if err != nil {
			log.Fatalf("invalid latitude: %v", err)
		}

		lon, err := strconv.ParseFloat(lonStr, 64)
		if err != nil {
			log.Fatalf("invalid longitude: %v", err)
		}
		geofenceRadius, err := strconv.ParseFloat(geofenceRadiusStr, 64)
		if err != nil {
			log.Fatalf("invalid geofence radius: %v", err)
		}
		var distance = commonsHelper.CalculateDistance(data.Latitude, data.Longitude, lat, lon)
		if distance <= geofenceRadius {
			event := map[string]interface{}{
				"vehicle_id": data.VehicleID,
				"event":      "geofence_entry",
				"location": map[string]float64{
					"latitude":  data.Latitude,
					"longitude": data.Longitude,
				},
				"timestamp": data.Timestamp.Unix(),
			}

			payload, _ := json.Marshal(event)
			err := rabbitmqProducer.PublishMessage("fleet.events", "geofence_alerts", payload)
			if err != nil {
				log.Println("Failed to publish geofence event:", err)
			}
			log.Println("Geofence event published:")
		} else {
			log.Println("distance is greater than geofence radius"+" distance: ", distance, " geofence radius: ", geofenceRadius)
		}
	})
}
