package controllers

import (
	"github.com/gin-gonic/gin"
)


func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to get all food items
	}
}

func GetFoodById() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to get food item by ID
	}
}
func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to create a new food item
	}
}
func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to update an existing food item
	}
}
func DeleteFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to delete a food item
	}
}
