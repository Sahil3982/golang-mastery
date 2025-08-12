
package controllers
import (
	"github.com/gin-gonic/gin"
)

func GetUsers() gin.HandlerFunc{
	return func(c *gin.Context){

	}	
}

func GetUserById() gin.HandlerFunc{
	return func(c *gin.Context){
		// Logic to get user by ID
	}
}

func SignUp() gin.HandlerFunc{
	return func(c *gin.Context){
		// Logic for user signup
	}
}

func SignIn() gin.HandlerFunc{
	return func(c *gin.Context){
		// Logic for user signin
	}
}


