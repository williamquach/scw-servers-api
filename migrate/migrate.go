package main

import (
	"servers_api/initializers"
	"servers_api/models"

	"gorm.io/gorm"
)

func initialize() *gorm.DB {
	initializers.LoadEnvironmentVariables()
	db := initializers.ConnectDatabase()
	return db
}

func main() {
	db := initialize()
	db.AutoMigrate(&models.ServerModel{})
}
