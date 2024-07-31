package repository

import (
	"e-commerce/internal/domain"
	"gorm.io/gorm"
)

type User interface {
	SaveUser(user *domain.User) error
	GetUserByEmail(email string) (*domain.User, error)
	GetAllUsers() ([]domain.User, error)
	GetUserByID(id string) (*domain.User, error)
	UpdateUser(id string, updatedUser *domain.User) error
	DeleteUser(id string) error
	SearchUsersByName(name string) ([]domain.User, error)
	SearchUsersByEmail(email string) ([]domain.User, error)
}

type Order interface {
	SaveOrder(order *domain.Order) error
	GetOrderById(id uint) (*domain.Order, error)
	GetAllOrders() ([]domain.Order, error)
	UpdateOrder(id uint, updatedOrder *domain.Order) error
	DeleteOrder(id uint) error
	SearchOrdersByUserID(userID string) ([]domain.Order, error)
	SearchOrdersByStatus(status string) ([]domain.Order, error)
}

type Product interface {
	SaveProduct(product *domain.Product) error
	GetAllProducts() ([]domain.Product, error)
	GetProductByID(id string) (*domain.Product, error)
	UpdateProduct(id string, updatedProduct *domain.Product) error
	DeleteProduct(id string) error
	SearchProductsByName(name string) ([]domain.Product, error)
	SearchProductsByCategory(category string) ([]domain.Product, error)
}

type Payment interface {
	GetAllPayments() ([]domain.Payment, error)
	CreatePayment(payment *domain.Payment) error
	GetPaymentByID(id string) (*domain.Payment, error)
	UpdatePayment(payment *domain.Payment) error
	DeletePayment(id string) error
	SearchPaymentsByUserID(userID string) ([]domain.Payment, error)
	SearchPaymentsByOrderID(orderID string) ([]domain.Payment, error)
	SearchPaymentsByStatus(status string) ([]domain.Payment, error)
}

type Repository struct {
	User
	Order
	Product
	Payment
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User:    NewUserRepository(db),
		Order:   NewOrderRepository(db),
		Product: NewProductRepository(db),
		Payment: NewPaymentRepository(db),
	}
}
