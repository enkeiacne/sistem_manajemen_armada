package databaseMigrations

import (
	"transakarta_BE_test/internal/database"
	databaseEntities "transakarta_BE_test/internal/database/entities"
)

func loadExtensions() error {
	if err := database.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error; err != nil {
		return err
	}
	return nil
}
func DatabaseMigration() error {
	err := loadExtensions()
	if err != nil {
		return err
	}
	err = database.DB.AutoMigrate(
		//databaseEntities.Vehicle{},
		databaseEntities.VehicleLocation{},
	)
	if err != nil {
		return err
	}
	return nil
}
