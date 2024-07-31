package repository

import (
	"e-commerce/internal/domain"
	"gorm.io/gorm"
)

type PaymentRepository struct {
	DB *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{
		DB: db,
	}
}

func (repo *PaymentRepository) GetAllPayments() ([]domain.Payment, error) {
	var payments []domain.Payment
	err := repo.DB.Find(&payments).Error
	return payments, err
}

func (repo *PaymentRepository) CreatePayment(payment *domain.Payment) error {
	return repo.DB.Create(payment).Error
}

func (repo *PaymentRepository) GetPaymentByID(id string) (*domain.Payment, error) {
	var payment domain.Payment
	err := repo.DB.First(&payment, "id = ?", id).Error
	return &payment, err
}

func (repo *PaymentRepository) UpdatePayment(payment *domain.Payment) error {
	return repo.DB.Save(payment).Error
}

func (repo *PaymentRepository) DeletePayment(id string) error {
	return repo.DB.Delete(&domain.Payment{}, "id = ?", id).Error
}

func (repo *PaymentRepository) SearchPaymentsByUserID(userID string) ([]domain.Payment, error) {
	var payments []domain.Payment
	err := repo.DB.Where("user_id = ?", userID).Find(&payments).Error
	return payments, err
}

func (repo *PaymentRepository) SearchPaymentsByOrderID(orderID string) ([]domain.Payment, error) {
	var payments []domain.Payment
	err := repo.DB.Where("order_id = ?", orderID).Find(&payments).Error
	return payments, err
}

func (repo *PaymentRepository) SearchPaymentsByStatus(status string) ([]domain.Payment, error) {
	var payments []domain.Payment
	err := repo.DB.Where("status = ?", status).Find(&payments).Error
	return payments, err
}
