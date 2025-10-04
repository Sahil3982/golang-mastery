package controllers

import (
	"context"
	"net/http"
	"restaurant-management/database"
	"restaurant-management/models"
	"time"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		foodId := c.Param("food_id")

		var food models.Food

		err := foodCollection.FindOne(ctx, bson.M{"food_id": foodId}).Decode(&food)

		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching food item"})
		}
		c.JSON(http.StatusOK, gin.H{"food": food})
	}
}

func GetFoodById() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to get food item by ID
	}
}
func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
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
