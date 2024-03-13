package database

import (
	"log"
	"time"
	"truckrud/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	dsn := "host=db user=postgres password=password dbname=myappdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	maxAttempts := 10

	for attempts := 1; attempts <= maxAttempts; attempts++ {
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("Successfully connected to the database.")
			db.AutoMigrate(&models.Driver{}, &models.Vehicle{}, &models.DriverVehicleLink{})
			DB = db
			return
		}

		log.Printf("Failed to connect to database (attempt %d/%d): %v", attempts, maxAttempts, err)
		time.Sleep(time.Duration(attempts) * time.Second)
	}

	log.Fatalf("Failed to connect to the database after %d attempts.", maxAttempts)
}
