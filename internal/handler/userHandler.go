package handler

import (
	"e-commerce/internal/domain"
	"e-commerce/internal/repository"
	"net/http"

	"e-commerce/internal/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	UserRepo *repository.UserRepository
}

func NewUserHandler(ur *repository.UserRepository) *UserHandler {
	return &UserHandler{UserRepo: ur}
}

func (uh *UserHandler) CreateUser(c *gin.Context) {
	var user domain.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error decoding request body"})
		return
	}

	if err := validation.ValidateStruct(&user); err != nil {
		errorMessage := validation.HandleValidationErrors(err.(validator.ValidationErrors), domain.UserBaseMessages)
		c.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		return
	}

	if _, err := uh.UserRepo.GetUserByEmail(user.Email); err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	if err := uh.UserRepo.SaveUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully!"})
}

func (uh *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := uh.UserRepo.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving users"})
		return
	}

	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Users not found"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (uh *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := uh.UserRepo.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (uh *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var updatedUser domain.User
	if err := c.BindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error decoding request body"})
		return
	}

	if err := validation.ValidateStruct(&updatedUser); err != nil {
		errorMessage := validation.HandleValidationErrors(err.(validator.ValidationErrors), domain.UserBaseMessages)
		c.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		return
	}

	existingUser, err := uh.UserRepo.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if updatedUser.Email != existingUser.Email {
		if _, err := uh.UserRepo.GetUserByEmail(updatedUser.Email); err == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
			return
		}
	}

	if err := uh.UserRepo.UpdateUser(id, &updatedUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully!"})
}

func (uh *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	_, err := uh.UserRepo.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := uh.UserRepo.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully!"})
}

func (uh *UserHandler) SearchUsersByName(c *gin.Context) {
	name := c.Query("name")
	users, err := uh.UserRepo.SearchUsersByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error searching users by name"})
		return
	}

	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Users not found"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (uh *UserHandler) SearchUsersByEmail(c *gin.Context) {
	email := c.Param("email")
	user, err := uh.UserRepo.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error searching users by email"})
		return
	}

	c.JSON(http.StatusOK, user)
}
