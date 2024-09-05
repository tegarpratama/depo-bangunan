package main

import (
	"depo-bangunan/config"
	"depo-bangunan/docs"
	"depo-bangunan/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath /api
func main() {
	config.LoadConfig()
	config.ConnectDB()

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api"

	router := r.Group("/api")

	routes.AuthRoute(router)
	routes.CustomerRoute(router)
	routes.ProductRoute(router)
	routes.OrderRoute(router)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	log.Fatal(r.Run(fmt.Sprintf("127.0.0.1:%v", config.ENV.PORT)))
}