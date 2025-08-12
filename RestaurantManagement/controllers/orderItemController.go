package controllers	

import (
	"github.com/gin-gonic/gin"
)

func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to get all order items
	}
}

func GetOrderItemById() gin.HandlerFunc {

	return func(c *gin.Context) {
		// Logic to get order item by ID
	}
}

func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to create a new order item
	}
}

func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to update an existing order item
	}
}

func DeleteOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to delete an order item
	}
}