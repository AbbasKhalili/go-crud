package controllers

import (
	"go-crud/config"
	"go-crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetProducts godoc
// @Summary      Get a product
// @Description  add by json product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product body models.Product true "Get a product"
// @Success      200  {object} models.Product
// @Router       /products [get]
func GetProducts(c *gin.Context) {
	var products []models.Product
	config.DB.Find(&products)
	c.JSON(http.StatusOK, products)
}

// GetProduct godoc
// @Summary      Get a product
// @Description  add by json product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product body models.Product true "Get a product"
// @Success      200  {object} models.Product
// @Router       /products [get]
func GetProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// CreateProduct godoc
// @Summary      Create a product
// @Description  add by json product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product body models.Product true "Add product"
// @Success      200  {object} models.Product
// @Router       /products [post]
func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&product)
	c.JSON(http.StatusCreated, product)
}

// UpdateProduct godoc
// @Summary      Update a product
// @Description  add by json product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product body models.Product true "Update product"
// @Success      200  {object} models.Product
// @Router       /products [put]
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&product)
	c.JSON(http.StatusOK, product)
}

// DeleteProduct godoc
// @Summary      Delete a product
// @Description  add by json product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product body models.Product true "Delete product"
// @Success      200  {object} models.Product
// @Router       /products [delete]
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Product{}, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Delete failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
