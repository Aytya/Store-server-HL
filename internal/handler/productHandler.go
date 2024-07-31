package handler

import (
	"e-commerce/internal/domain"
	"e-commerce/internal/repository"
	"e-commerce/internal/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type ProductHandler struct {
	ProductRepo *repository.ProductRepository
}

func NewProductHandler(pr *repository.ProductRepository) *ProductHandler {
	return &ProductHandler{ProductRepo: pr}
}

func (ph *ProductHandler) CreateProduct(c *gin.Context) {
	var product domain.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error decoding request body"})
		return
	}

	if err := validation.ValidateStruct(&product); err != nil {
		errorMessage := validation.HandleValidationErrors(err.(validator.ValidationErrors), domain.ProductBaseMessages)
		c.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		return
	}

	if err := ph.ProductRepo.SaveProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving product"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully!"})
}

func (ph *ProductHandler) GetAllProducts(c *gin.Context) {
	products, err := ph.ProductRepo.GetAllProducts()

	if len(products) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No products found"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching products"})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (ph *ProductHandler) GetProductByID(c *gin.Context) {
	id := c.Param("id")
	product, err := ph.ProductRepo.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (ph *ProductHandler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	var product domain.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error decoding request body"})
		return
	}

	if err := validation.ValidateStruct(&product); err != nil {
		errorMessage := validation.HandleValidationErrors(err.(validator.ValidationErrors), domain.ProductBaseMessages)
		c.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		return
	}

	existingProduct, err := ph.ProductRepo.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	product.CreatedAt = existingProduct.CreatedAt

	if err := ph.ProductRepo.UpdateProduct(id, &product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully!"})
}

func (ph *ProductHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	_, err := ph.ProductRepo.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if err := ph.ProductRepo.DeleteProduct(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Error deleting product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully!"})
}

func (ph *ProductHandler) SearchProductsByName(c *gin.Context) {
	name := c.Query("name")
	products, err := ph.ProductRepo.SearchProductsByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error searching products by name"})
		return
	}
	if len(products) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No products found"})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (ph *ProductHandler) SearchProductsByCategory(c *gin.Context) {
	category := c.Query("category")
	products, err := ph.ProductRepo.SearchProductsByCategory(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error searching products by category"})
		return
	}
	if len(products) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No products found"})
		return
	}
	c.JSON(http.StatusOK, products)
}
