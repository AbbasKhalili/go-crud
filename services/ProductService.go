package services

import (
	"errors"
	"go-crud/models"
	"go-crud/repositories"
)

type ProductService interface {
	GetProduct(id int) (*models.Product, error)
	GetAllProduct() (*[]models.Product, error)
	CreateProduct(input models.Product) (*models.Product, error)
	UpdateProduct(id int, input models.Product) (*models.Product, error)
	DeleteProduct(id int) error
}

type productService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) ProductService {
	return &productService{repo}
}

func (s *productService) GetAllProduct() (*[]models.Product, error) {
	return s.repo.GetAll()
}

func (s *productService) GetProduct(id int) (*models.Product, error) {
	return s.repo.GetById(id)
}
func (s *productService) CreateProduct(input models.Product) (*models.Product, error) {
	if input.CategoryId <= 0 {
		return nil, errors.New("CategoryId must be non-negative")
	}
	err := s.repo.Create(&input)
	return &input, err
}
func (s *productService) UpdateProduct(id int, input models.Product) (*models.Product, error) {
	product, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	// Apply business logic
	if input.CategoryId <= 0 {
		return nil, errors.New("CategoryId must be non-negative")
	}
	product.Id = id
	product.Name = input.Name
	product.CategoryId = input.CategoryId
	product.Ingredient = input.Ingredient
	err = s.repo.Update(product)
	return product, err
}

func (s *productService) DeleteProduct(id int) error {
	return s.repo.Delete(id)
}
