package routes

import (
	"depo-bangunan/controllers"
	"depo-bangunan/middleware"

	"github.com/gin-gonic/gin"
)

func CustomerRoute(rg *gin.RouterGroup) {
	router := rg.Group("customers")

	router.Use(middleware.Auth())

	router.GET("/", controllers.GetCustomers)
	router.POST("/", controllers.CreateCustomers)
	router.GET("/:id/detail", controllers.DetailCustomers)
	router.PUT("/:id/update", controllers.UpdateCustomer)
	router.DELETE("/:id/delete", controllers.DeleteCustomer)
}
