package controllers

import (
	"go-crud/Dtos"
	"go-crud/models"
	"go-crud/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	service services.ProductService
}

func NewProductController(s services.ProductService) *ProductController {
	return &ProductController{service: s}
}

// GetProducts godoc
// @Summary      Get all product
// @Description  add by json product
// @Tags         products
// @Accept       json
// @Produce      json
// @Success      200 {object} Dtos.GenericResponseDto
// @Failure		 400 {object} Dtos.ErrorResponseDto
// @Router       /products [get]
func (pc *ProductController) GetProducts(c *gin.Context) {
	products, err := pc.service.GetAllProduct()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Dtos.ErrorResponseDto{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Dtos.GenericResponseDto{Data: products})
}

// GetProduct godoc
// @Summary      Get a product
// @Description  add by json product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param   id     path    int     true  "Product Id"
// @Success      200  {object} Dtos.GenericResponseDto
// @Failure      400  {object} Dtos.ErrorResponseDto
// @Router       /products/{id} [get]
func (pc *ProductController) GetProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := pc.service.GetProduct(id)
	if err != nil {
		c.JSON(http.StatusNotFound, Dtos.ErrorResponseDto{Error: err.Error()})
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
// @Success      200  {object} Dtos.GenericResponseDto
// @Failure      400  {object} Dtos.ErrorResponseDto
// @Router       /products [post]
func (pc *ProductController) CreateProduct(c *gin.Context) {
	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, Dtos.ErrorResponseDto{Error: err.Error()})
		return
	}
	product, err := pc.service.CreateProduct(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, Dtos.ErrorResponseDto{Error: err.Error()})
		return
	}
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
// @Success      200 {object} Dtos.GenericResponseDto
// @Failure      400 {object} Dtos.ErrorResponseDto
// @Router       /products/{id} [put]
func (pc *ProductController) UpdateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, Dtos.ErrorResponseDto{Error: "Invalid Product Id"})
		return
	}
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, Dtos.ErrorResponseDto{Error: err.Error()})
		return
	}
	result, err1 := pc.service.UpdateProduct(id, product)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, Dtos.ErrorResponseDto{Error: err1.Error()})
	}
	c.JSON(http.StatusOK, Dtos.GenericResponseDto{Data: result})
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
func (pc *ProductController) DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := pc.service.DeleteProduct(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, Dtos.ErrorResponseDto{Error: "Delete failed"})
		return
	}
	c.JSON(http.StatusOK, Dtos.GenericResponseDto{Message: "Product deleted", Data: true})
}
