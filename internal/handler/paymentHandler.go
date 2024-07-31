package handler

import (
	"e-commerce/internal/domain"
	"e-commerce/internal/repository"
	"e-commerce/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

const (
	DefaultPaymentData = `{
 "hpan":"4405639704015096","expDate":"0125","cvc":"815","terminalId":"67e34d63-102f-4bd1-898e-370781d0074d"
}`
)

type PaymentHandler struct {
	repo *repository.PaymentRepository
}

func NewPaymentHandler(repository *repository.PaymentRepository) *PaymentHandler {
	return &PaymentHandler{
		repo: repository,
	}
}

func (h *PaymentHandler) GetAllPayments(c *gin.Context) {
	payments, err := h.repo.GetAllPayments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payments)
}

func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	var payment domain.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := service.GetPaymentToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	encryptedData, err := service.EncryptData(DefaultPaymentData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Encryption failed"})
		return
	}

	paymentResponse, err := service.MakePayment(token, encryptedData)
	if err != nil {
		log.Printf("Failed to make payment: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to make payment"})
		return
	}

	payment.PaymentStatus = paymentResponse.Status

	if err := h.repo.CreatePayment(&payment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, payment)
}

func (h *PaymentHandler) GetPaymentByID(c *gin.Context) {
	id := c.Param("id")
	payment, err := h.repo.GetPaymentByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}
	c.JSON(http.StatusOK, payment)
}

func (h *PaymentHandler) UpdatePayment(c *gin.Context) {
	id := c.Param("id")
	var payment domain.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	payment.ID = uint(idUint)

	if err := h.repo.UpdatePayment(&payment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payment)
}

func (h *PaymentHandler) DeletePayment(c *gin.Context) {
	id := c.Param("id")
	if err := h.repo.DeletePayment(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (h *PaymentHandler) SearchPaymentsByUserID(c *gin.Context) {
	userID := c.Query("user")
	payments, err := h.repo.SearchPaymentsByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payments)
}

func (h *PaymentHandler) SearchPaymentsByOrderID(c *gin.Context) {
	orderID := c.Query("order")
	payments, err := h.repo.SearchPaymentsByOrderID(orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payments)
}

func (h *PaymentHandler) SearchPaymentsByStatus(c *gin.Context) {
	status := c.Query("status")
	payments, err := h.repo.SearchPaymentsByStatus(status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payments)
}
