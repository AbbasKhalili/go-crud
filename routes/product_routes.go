package routes

import (
	"github.com/gin-gonic/gin"
	"go-crud/config"
	"go-crud/controllers"
	"go-crud/repositories"
	"go-crud/services"
)

func ProductRoute(r *gin.Engine) {
	repo := repositories.NewProductRepository(config.DB)
	service := services.NewProductService(repo)
	controller := controllers.NewProductController(service)

	r.GET("/products", controller.GetProducts)
	r.GET("/products/:id", controller.GetProduct)
	r.POST("/products", controller.CreateProduct)
	r.PUT("/products/:id", controller.UpdateProduct)
	r.DELETE("/products/:id", controller.DeleteProduct)
}
