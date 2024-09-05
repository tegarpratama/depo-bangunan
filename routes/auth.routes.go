package routes

import (
	"depo-bangunan/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoute(rg *gin.RouterGroup) {
	router := rg.Group("auth")

	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
}
