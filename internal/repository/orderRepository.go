package repository

import (
	"e-commerce/internal/domain"
	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (or *OrderRepository) SaveOrder(order *domain.Order) error {
	return or.DB.Create(&order).Error
}

func (or *OrderRepository) GetOrderById(id uint) (*domain.Order, error) {
	var order domain.Order
	if err := or.DB.Where("id = ?", id).First(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (or *OrderRepository) GetAllOrders() ([]domain.Order, error) {
	var orders []domain.Order
	if err := or.DB.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (or *OrderRepository) UpdateOrder(id uint, updatedOrder *domain.Order) error {
	if err := or.DB.Model(&domain.Order{}).Where("id = ?", id).Updates(updatedOrder).Error; err != nil {
		return err
	}
	return nil
}

func (or *OrderRepository) DeleteOrder(id uint) error {
	if err := or.DB.Where("id = ?", id).Delete(&domain.Order{}).Error; err != nil {
		return err
	}
	return nil
}

func (or *OrderRepository) SearchOrdersByUserID(userID string) ([]domain.Order, error) {
	var orders []domain.Order
	if err := or.DB.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (or *OrderRepository) SearchOrdersByStatus(status string) ([]domain.Order, error) {
	var orders []domain.Order
	if err := or.DB.Where("status = ?", status).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}
