package main

import (
	"assigment-2/config"
	"assigment-2/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.DbInit()
	postgresDB := &controllers.Database{DB: db}
	router := gin.Default()
	router.GET("/orders", postgresDB.GetOrders)
	router.POST("/orders", postgresDB.CreateOrder)
	router.PUT("orders/:id", postgresDB.UpdateOrder)
	router.DELETE("orders/:id", postgresDB.DeleteOrder)
	router.Run(":3000")
}
