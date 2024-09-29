package router

import (
	"nyctaxi_mapup/pkg/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/getTaxiData", controller.GetTaxiData)
	return router
}
