package routes

import (
	"github.com/gin-gonic/gin"
	controller "restaurant-management/controllers"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/menu", controller.GetMenu())
	incomingRoutes.GET("/menu/:menu_id", controller.GetMenuById())
	incomingRoutes.POST("/menu", controller.CreateMenu())
	incomingRoutes.PATCH("/menu/:menu_id", controller.UpdateMenu())
	incomingRoutes.DELETE("/menu/:menu_id", controller.DeleteMenu())
}