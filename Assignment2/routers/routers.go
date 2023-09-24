package routers

import (
	controllers "assignment2/controls"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.PUT("/order", controllers.Home)
	router.POST("/order", controllers.CreateOrder)
	router.GET("/order", controllers.GetOrders)
	router.GET("/order/:id", controllers.GetOrderById)
	router.PUT("/order/:id", controllers.UpdateOrder)

	return router
}
