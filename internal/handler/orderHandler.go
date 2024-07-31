package handler

import (
	"e-commerce/internal/domain"
	"e-commerce/internal/repository"
	"e-commerce/internal/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

type OrderHandler struct {
	OrderRepo   *repository.OrderRepository
	UserRepo    *repository.UserRepository
	ProductRepo *repository.ProductRepository
}

func NewOrderHandler(or *repository.OrderRepository, ur *repository.UserRepository, pr *repository.ProductRepository) *OrderHandler {
	return &OrderHandler{OrderRepo: or, UserRepo: ur, ProductRepo: pr}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var order domain.Order
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error decoding request body"})
		return
	}

	if err := validation.ValidateStruct(&order); err != nil {
		errorMessage := validation.HandleValidationErrors(err.(validator.ValidationErrors), domain.OrderBaseMessages)
		c.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		return
	}

	userIDStr := strconv.Itoa(int(order.UserID))
	if _, err := h.UserRepo.GetUserByID(userIDStr); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	for _, productID := range order.ProductIDs {
		if _, err := h.ProductRepo.GetProductByID(strconv.Itoa(int(productID))); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product with ID " + strconv.Itoa(int(productID)) + " not found"})
			return
		}
	}

	if err := h.OrderRepo.SaveOrder(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving order"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Order created successfully!"})
}

func (h *OrderHandler) GetOrderByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	order, err := h.OrderRepo.GetOrderById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) GetAllOrders(c *gin.Context) {
	orders, err := h.OrderRepo.GetAllOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving orders", "details": err.Error()})
		return
	}

	if len(orders) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No orders found"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) UpdateOrder(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	var updatedOrder domain.Order
	if err := c.BindJSON(&updatedOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error decoding request body"})
		return
	}

	if err := validation.ValidateStruct(&updatedOrder); err != nil {
		errorMessage := validation.HandleValidationErrors(err.(validator.ValidationErrors), domain.OrderBaseMessages)
		c.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		return
	}

	existingOrder, err := h.OrderRepo.GetOrderById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	updatedOrder.OrderDate = existingOrder.OrderDate

	if err := h.OrderRepo.UpdateOrder(uint(id), &updatedOrder); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order updated successfully!"})
}

func (h *OrderHandler) DeleteOrder(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	_, err = h.OrderRepo.GetOrderById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	if err := h.OrderRepo.DeleteOrder(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully!"})
}

func (h *OrderHandler) SearchOrdersByUserID(c *gin.Context) {
	userID := c.Param("user")
	orders, err := h.OrderRepo.SearchOrdersByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error searching orders by user ID"})
		return
	}

	if len(orders) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No orders found"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) SearchOrdersByStatus(c *gin.Context) {
	status := c.Query("status")
	orders, err := h.OrderRepo.SearchOrdersByStatus(status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error searching orders by status"})
		return
	}

	if len(orders) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No orders found for the given status"})
		return
	}

	c.JSON(http.StatusOK, orders)
}
