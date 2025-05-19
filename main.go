// @title        Product API
// @version      1.0
// @description  REST API for managing products
// @host         localhost:8080
// @BasePath     /

package main

import (
	"github.com/gin-gonic/gin"
	"go-crud/config"
	_ "go-crud/docs"
	"go-crud/models"
	"go-crud/routes"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Product{})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.ProductRoute(r)

	r.Run(":8080")
}
