package main

import (
	"depo-bangunan/config"
	"depo-bangunan/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	config.LoadConfig()

	connection := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true&loc=Asia%vJakarta", config.ENV.DB_USER, config.ENV.DB_PASSWORD, config.ENV.DB_HOST, config.ENV.DB_PORT, config.ENV.DB_DATABASE, "%2F")
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})
	log.Println("Migration Completed")
	
	seedUsers(db)
	log.Println("Seed Completed")
}

func seedUsers(db *gorm.DB) {
	users := []models.User{
		{Name: "admin", Email: "admin@gmail.com", Password: "$2a$10$Gw.6W2PD0nrUMNHNG29ZzOVYrcz64moGiJRV5t0qZmDhWJqf8zHvW", Role: "admin"},	
	}

	if err := db.Create(&users).Error; err != nil {
		log.Fatalf("failed to seed users: %v", err)
	}
	
	products := []models.Product{
		{Name: "Nike Air Jordan", Price: 1500000},	
		{Name: "PS5", Price: 5500000},	
		{Name: "Tumbler", Price: 100000},	
	}

	if err := db.Create(&products).Error; err != nil {
		log.Fatalf("failed to seed products: %v", err)
	}
}