package repository

import (
	"e-commerce/internal/domain"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (pr *ProductRepository) SaveProduct(product *domain.Product) error {
	return pr.DB.Create(&product).Error
}

func (pr *ProductRepository) GetAllProducts() ([]domain.Product, error) {
	var products []domain.Product
	err := pr.DB.Find(&products).Error
	return products, err
}

func (pr *ProductRepository) GetProductByID(id string) (*domain.Product, error) {
	var product domain.Product
	err := pr.DB.First(&product, "id = ?", id).Error
	return &product, err
}

func (pr *ProductRepository) UpdateProduct(id string, updatedProduct *domain.Product) error {
	return pr.DB.Model(&domain.Product{}).Where("id = ?", id).Updates(updatedProduct).Error
}

func (pr *ProductRepository) DeleteProduct(id string) error {
	return pr.DB.Delete(&domain.Product{}, id).Error
}

func (pr *ProductRepository) SearchProductsByName(name string) ([]domain.Product, error) {
	var products []domain.Product
	err := pr.DB.Where("name ILIKE ?", "%"+name+"%").Find(&products).Error
	return products, err
}

func (pr *ProductRepository) SearchProductsByCategory(category string) ([]domain.Product, error) {
	var products []domain.Product
	err := pr.DB.Where("category ILIKE ?", "%"+category+"%").Find(&products).Error
	return products, err
}
