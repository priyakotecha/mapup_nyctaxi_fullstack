package router

import (
	"nyctaxi_mapup/pkg/controller"
	"nyctaxi_mapup/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()
	router.POST("/login", controller.Login)

	adminRoutes := router.Group("/admin")
	adminRoutes.Use(middleware.AuthMiddleware("admin"))
	{
		adminRoutes.GET("/getData", controller.GetTaxiData)
		adminRoutes.POST("/ingest", controller.IngestData)
	}

	managerRoutes := router.Group("/manager")
	managerRoutes.Use(middleware.AuthMiddleware("manager"))
	{
		managerRoutes.GET("/getdata", controller.GetTaxiData)
		managerRoutes.POST("/ingest", controller.IngestData) // Manager can ingest data
	}

	userRoutes := router.Group("/user")
	userRoutes.Use(middleware.AuthMiddleware("user"))
	{
		userRoutes.GET("/getdata", controller.GetTaxiData) // User can only read data
	}

	router.GET("/data-stream", controller.GetTaxiData)
	return router
}
