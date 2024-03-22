package config

import (
	"assigment-2/model"
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// variable
var (
	host     = "localhost"
	username = "postgres"
	password = "admin"
	dbPort   = "5432"
	dbName   = "orders_by"
	db       *gorm.DB
	err      error
)

// Database
func DbInit() *gorm.DB {
	config := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, username, password, dbName, dbPort)
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}
	db.Debug().AutoMigrate(model.Order{}, model.Item{})
	return db

}
