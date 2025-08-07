package routes

import (
	"github.com/gin-gonic/gin"
	"restaurantManagement/controllers"
)

func UserRoutes(incommingRoutes *gin.Engine) {

	incommingRoutes.GET("/users",controller.GetUsers())
	incommingRoutes.GET("/users/user_id", controller.GetUserById())
	incommingRoutes.POST("/users/signup", controller.SignUp())
	incommingRoutes.POST("/users/signin", controller.SignIn())
}