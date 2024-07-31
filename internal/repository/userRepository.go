package repository

import (
	"e-commerce/internal/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (repo *UserRepository) SaveUser(user *domain.User) error {
	return repo.DB.Create(&user).Error
}

func (repo *UserRepository) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := repo.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) GetAllUsers() ([]domain.User, error) {
	var users []domain.User
	if err := repo.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) GetUserByID(id string) (*domain.User, error) {
	var user domain.User
	if err := repo.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) UpdateUser(id string, updatedUser *domain.User) error {
	if err := repo.DB.Model(&domain.User{}).Where("id = ?", id).Updates(updatedUser).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) DeleteUser(id string) error {
	if err := repo.DB.Where("id = ?", id).Delete(&domain.User{}).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) SearchUsersByName(name string) ([]domain.User, error) {
	var users []domain.User
	if err := repo.DB.Where("name LIKE ?", "%"+name+"%").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) SearchUsersByEmail(email string) ([]domain.User, error) {
	var users []domain.User
	if err := repo.DB.Where("email = ?", email).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
