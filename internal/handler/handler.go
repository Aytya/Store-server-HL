package handler

import (
	"e-commerce/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go"
	"os"
)

type Handler struct {
	order   *OrderHandler
	user    *UserHandler
	product *ProductHandler
	payment *PaymentHandler
}

func NewHandler(order *repository.OrderRepository, payment *repository.PaymentRepository, user *repository.UserRepository, product *repository.ProductRepository) *Handler {
	return &Handler{
		order:   NewOrderHandler(order, user, product),
		user:    NewUserHandler(user),
		product: NewProductHandler(product),
		payment: NewPaymentHandler(payment),
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	stripe.Key = os.Getenv("STRIPE_KEY")

	user := router.Group("/user")
	{
		user.GET("/", h.user.GetAllUsers)
		user.POST("/", h.user.CreateUser)
		user.PUT("/:id", h.user.UpdateUser)
		user.DELETE("/:id", h.user.DeleteUser)
		user.GET("/:id", h.user.GetUserByID)
		user.GET("/search/:name", h.user.SearchUsersByName)
		user.GET("search/email/:email", h.user.SearchUsersByEmail)
	}

	product := router.Group("/products")
	{
		product.GET("/", h.product.GetAllProducts)
		product.POST("/", h.product.CreateProduct)
		product.PUT("/:id", h.product.UpdateProduct)
		product.DELETE("/:id", h.product.DeleteProduct)
		product.GET("/:id", h.product.GetProductByID)
		product.GET("/search/:name", h.product.SearchProductsByName)
		product.GET("/search/category/:category", h.product.SearchProductsByCategory)
	}

	order := router.Group("/orders")
	{
		order.GET("/", h.order.GetAllOrders)
		order.POST("/", h.order.CreateOrder)
		order.PUT("/:id", h.order.UpdateOrder)
		order.DELETE("/:id", h.order.DeleteOrder)
		order.GET("/:id", h.order.GetOrderByID)
		order.GET("/search", h.order.SearchOrdersByStatus)
		order.GET("/search/:user", h.order.SearchOrdersByUserID)
	}

	payment := router.Group("/payments")
	{
		payment.GET("/", h.payment.GetAllPayments)
		payment.POST("/", h.payment.CreatePayment)
		payment.PUT("/:id", h.payment.UpdatePayment)
		payment.DELETE("/:id", h.payment.DeletePayment)
		payment.GET("/:id", h.payment.GetPaymentByID)
		payment.GET("/search/user/:user_id", h.payment.SearchPaymentsByUserID)
		payment.GET("/search/:order_id", h.payment.SearchPaymentsByOrderID)
		payment.GET("/search", h.payment.SearchPaymentsByStatus)
	}

	return router
}
