package controllers

import (
	"go-crud/Dtos"
	"go-crud/config"
	"go-crud/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetProducts godoc
// @Summary      Get all product
// @Description  add by json product
// @Tags         products
// @Accept       json
// @Produce      json
// @Success      200  {object} []models.Product
// @Router       /products [get]
func GetProducts(c *gin.Context) {
	var products []models.Product
	config.DB.Find(&products)
	c.JSON(http.StatusOK, Dtos.GenericResponseDto{Data: products})
}

// GetProduct godoc
// @Summary      Get a product
// @Description  add by json product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param   id     path    int     true  "Product Id"
// @Success      200  {object} models.Product
// @Router       /products/{id} [get]
func GetProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, Dtos.ErrorResponseDto{Error: "product not found"})
		return
	}
	c.JSON(http.StatusOK, Dtos.GenericResponseDto{Data: product})
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
		c.JSON(http.StatusBadRequest, Dtos.ErrorResponseDto{Error: err.Error()})
		return
	}
	config.DB.Create(&product)
	c.JSON(http.StatusCreated, Dtos.GenericResponseDto{Data: product})
}

// UpdateProduct godoc
// @Summary      Update a product
// @Description  add by json product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Product Id"
// @Param        product body models.Product true "Update product"
// @Success      200  {object} models.Product
// @Router       /products/{id} [put]
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	number64, err := strconv.ParseInt(id, 10, 64) // base 10, 64-bit
	if err != nil {
		c.JSON(http.StatusBadRequest, Dtos.ErrorResponseDto{Error: "Invalid Product Id"})
		return
	}
	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, Dtos.ErrorResponseDto{Error: "Product not found"})
		return
	}
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, Dtos.ErrorResponseDto{Error: err.Error()})
		return
	}

	product.Id = int(number64)
	config.DB.Save(&product)
	c.JSON(http.StatusOK, Dtos.GenericResponseDto{Data: product})
}

// DeleteProduct godoc
// @Summary      Delete a product
// @Description  add by json product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Success      200  {object}  Dtos.GenericResponseDto
// @Failure      400  {object}  Dtos.ErrorResponseDto
// @Router       /products/{id} [delete]
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Product{}, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, Dtos.ErrorResponseDto{Error: "Delete failed"})
		return
	}
	c.JSON(http.StatusOK, Dtos.GenericResponseDto{Message: "Product deleted", Data: true})
}
