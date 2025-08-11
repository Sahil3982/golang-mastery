package routes

import (
	"github.com/gin-gonic/gin"
	controller "restaurant-management/controllers"
)

func OrderRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orders", controller.GetOrders())
	incomingRoutes.GET("/orders/:order_id", controller.GetOrderById())
	incomingRoutes.POST("/orders", controller.CreateOrder())
	incomingRoutes.PATCH("/orders/:order_id", controller.UpdateOrder())
	incomingRoutes.DELETE("/orders/:order_id", controller.DeleteOrder())
}