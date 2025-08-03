package repository

import (
	"github.com/infinity/identity-service/internal/storage"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository struct {
	Repository[storage.User]
	Log *logrus.Logger
}

func NewUserRepository(log *logrus.Logger) *UserRepository {
	return &UserRepository{
		Log: log,
	}
}

func (r *UserRepository) FindBySafeId(tx *gorm.DB, user *storage.User, safeID string) error {
	return tx.Where("safe_id = ?", safeID).First(user).Error
}

func (r *UserRepository) FindByPhoneNumber(tx *gorm.DB, user *storage.User, phoneNumber string) error {
	return tx.Where("phone_number = ?", phoneNumber).First(user).Error
}

func (r *UserRepository) FindByEmail(tx *gorm.DB, user *storage.User, email string) error {
	return tx.Where("email = ?", email).First(user).Error
}
