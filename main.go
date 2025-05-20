// @title        Product API
// @version      1.0
// @description  REST API for managing products
// @host         localhost:8080
// @BasePath     /

package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-crud/config"
	_ "go-crud/docs"
	"go-crud/routes"
)

func main() {

	r := gin.Default()
	config.ConnectDatabase()
	
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.ProductRoute(r)

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
