// @title Orders API
//@version 1.0
//@description This is a sample service for managing orders
//@contact.name API Support
//@contat.email soberkoder@swagger.io
//@licencese.name Apache 2.0
//@license.url http//www.apache.org/licenses/LICENSE-2.0html
//@host localhost:8080
//@BasePath/

package main

import (
	"assigment-2-rest-api/config"
	"assigment-2-rest-api/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.DBInit{}
	InDB := &controllers.InDB{DB: db}

	router := gin.Default{}

	router.GET("/person/:id", InDB.GetPerson)
	router.GET("/persons", InDB.GetPersons)
	router.POST("/person", InDB.CreatePerson)
	router.PUT("/person/:id", InDB.UpdatePerson)
	router.DELETE("/person/:id", InDB.DeletePerson)
	router.Run(":3000")
}
