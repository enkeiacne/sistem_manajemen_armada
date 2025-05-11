package main

import (
	"log"
	configEnviroment "transakarta_BE_test/internal/config/enviroment"
	"transakarta_BE_test/internal/database"
	databaseMigrations "transakarta_BE_test/internal/database/migrations"
)

func main() {
	configEnviroment.LoadEnv()
	database.DatabaseConnect()
	err := databaseMigrations.DatabaseMigration()
	if err != nil {
		panic(err)
	}
	log.Println("Successfully migrated database")
}
