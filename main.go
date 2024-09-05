package main

import (
	"depo-bangunan/config"
	"depo-bangunan/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration & db connection
	config.LoadConfig()
	config.ConnectDB()

	r := gin.Default()
	router := r.Group("/api")

	routes.AuthRoute(router)
	routes.CustomerRoute(router)
	routes.ProductRoute(router)
	routes.OrderRoute(router)

	log.Fatal(r.Run(fmt.Sprintf("127.0.0.1:%v", config.ENV.PORT)))
}