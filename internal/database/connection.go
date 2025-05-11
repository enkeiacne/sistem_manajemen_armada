package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	configEnviroment "transakarta_BE_test/internal/config/enviroment"
)

var DB *gorm.DB

func DatabaseConnect() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		configEnviroment.EnvironmentDatabaseHost,
		configEnviroment.EnvironmentDatabaseUser,
		configEnviroment.EnvironmentDatabasePassword,
		configEnviroment.EnvironmentDatabaseName,
		configEnviroment.EnvironmentDatabasePort,
		configEnviroment.EnvironmentDatabaseSSLMode,
	)
	var gormLogger logger.Interface = logger.Default.LogMode(logger.Silent)

	if os.Getenv("DATABASE_LOGGING") == "enable" {
		gormLogger = logger.Default.LogMode(logger.Info)
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}
	log.Println("Successfully connected to database")
	DB = db
}
