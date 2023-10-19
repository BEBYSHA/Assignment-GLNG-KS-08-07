package routers

import (
	"assignment-3/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Hello World")
	})

	router.GET("/weather", controllers.GetData)
	router.PUT("/weather", controllers.UpdateData)

	return router
}
