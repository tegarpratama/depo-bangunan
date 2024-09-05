package routes

import (
	"depo-bangunan/controllers"
	"depo-bangunan/middleware"

	"github.com/gin-gonic/gin"
)

func OrderRoute(rg *gin.RouterGroup) {
	router := rg.Group("orders")

	router.Use(middleware.Auth())

	router.GET("/", controllers.GetOrders)
	router.POST("/", controllers.CreateOrder)
	router.GET("/:id/detail", controllers.DetailOrder)
	router.PUT("/:id/update", controllers.UpdateOrder)
	router.DELETE("/:id/delete", controllers.DeleteOrder)
}
