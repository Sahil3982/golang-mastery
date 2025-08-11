package routes

import (
	"github.com/gin-gonic/gin"
	controller "restaurant-management/controllers"		
)


func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/order-items", controller.GetOrderItems())
	incomingRoutes.GET("/order-items/:order_item_id", controller.GetOrderItemById())
	incomingRoutes.POST("/order-items", controller.CreateOrderItem())
	incomingRoutes.PATCH("/order-items/:order_item_id", controller.UpdateOrderItem())
	incomingRoutes.DELETE("/order-items/:order_item_id", controller.DeleteOrderItem())
}