package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"restaurantManagement/database"
	"restaurantManagement/routes"
	"restaurantManagement/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRouters(router)
	routes.Use(middleware.Authentication())

	routes.foodRouters(router)
	routes.orderRouters(router)	
	routes.tableRouters(router)
	routes.reservationRouters(router)
	routes.menuRouters(router)
	routes.invoiceRouters(router)
	routes.orderItemRouters(router)


	router.Run(":" + port)


}
