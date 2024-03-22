package controllers

import (
	"assigment-2/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Create Order
func (db *Database) CreateOrder(ctx *gin.Context) {
	var (
		order  model.Order
		result gin.H
	)

	if err := ctx.BindJSON((&order)); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "JSON binding failed",
		})
		fmt.Println(err)
		return
	}

	insertedOrder := db.DB.Create(&order)
	if insertedOrder.Error != nil {
		result = gin.H{
			"message": "Failed to insert order to database",
		}
		return
	}

	result = gin.H{
		"status": http.StatusCreated,
		"result": order,
	}
	ctx.JSON(http.StatusCreated, result)
}

// Get Order
func (db *Database) GetOrders(ctx *gin.Context) {
	var (
		orders []model.Order
		result gin.H
	)
	db.DB.Preload("Items").Find(&orders)
	if len(orders) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": orders,
			"count":  len(orders),
		}
	}
	ctx.JSON(http.StatusOK, result)
}

// Update Order
func (db *Database) UpdateOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	var (
		order    model.Order
		newOrder model.Order
		result   gin.H
	)
	if bindErr := ctx.BindJSON(&order); bindErr != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "JSON binding error",
		})
		return
	}
	dbErr := db.DB.First(&newOrder, "id = ?", id).Error
	if dbErr != nil {
		result = gin.H{
			"message": "data not found",
		}
	}
	newOrder.CustomerName = order.CustomerName
	newOrder.Items = order.Items
	fmt.Println(newOrder)
	dbErr = db.DB.Model(&newOrder).Preload("items").Where("id = ?", id).Session(&gorm.Session{FullSaveAssociations: true}).Updates(newOrder).Error
	// dbErr2 = db.DB.Model(&newOrder.Items).Where("id = ? AND order_id= ?", id, order.Items.Id).Updates(newOrder).Error
	fmt.Println(db.DB.Model(&newOrder).Where("id = ?", id).Updates(newOrder))
	if dbErr != nil {
		fmt.Println(dbErr)
		result = gin.H{
			"message": "updating order failed",
		}
	} else {
		result = gin.H{
			"message": "updating order successfully",
		}

	}
	ctx.JSON(http.StatusAccepted, result)

}

// Delete Order
func (db *Database) DeleteOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	var (
		order  model.Order
		items  []model.Item
		result gin.H
	)
	err := db.DB.First(&order, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	err2 := db.DB.Where("order_id = ?", id).Delete(&items).Error
	if err = db.DB.Delete(&order, id).Error; err != nil || err2 != nil {
		fmt.Println(err)
		result = gin.H{
			"result": "Delete failed",
		}
	} else {
		result = gin.H{
			"result": "Deleted Successfully",
		}
	}
	ctx.JSON(http.StatusOK, result)
}
