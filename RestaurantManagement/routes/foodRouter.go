package routes

import (
	"github.com/gin-gonic/gin"	
	controller "restaurant-management/controllers"
)

func FoodRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/food", controller.GetFood())
	incomingRoutes.GET("/food/:food_id", controller.GetFoodById())
	incomingRoutes.POST("/food", controller.CreateFood())
	incomingRoutes.PATCH("/food/:food_id", controller.UpdateFood())
	incomingRoutes.DELETE("/food/:food_id", controller.DeleteFood())
}