package routes

import (
	"depo-bangunan/controllers"
	"depo-bangunan/middleware"

	"github.com/gin-gonic/gin"
)

func ProductRoute(rg *gin.RouterGroup) {
	router := rg.Group("products")

	router.Use(middleware.Auth())
	router.Use(middleware.Admin())

	router.GET("/", controllers.GetProducts)
	router.POST("/", controllers.CreateProduct)
	router.PUT("/:id/update", controllers.UpdateProduct)
	router.DELETE("/:id/delete", controllers.DeleteProduct)
}
