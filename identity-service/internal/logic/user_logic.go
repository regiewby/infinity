package logic

import (
	"context"

	"github.com/gofiber/fiber"
	"github.com/google/uuid"
	"github.com/infinity/identity-service/internal/model"
	"github.com/infinity/identity-service/internal/model/converter"
	"github.com/infinity/identity-service/internal/repository"
	"github.com/infinity/identity-service/internal/storage"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const logTag = "user.logic"

type UserLogic struct {
	Logger         *logrus.Logger
	DB             *gorm.DB
	UserRepository *repository.UserRepository
}

func NewUserLogic(logger *logrus.Logger, db *gorm.DB, userRepository *repository.UserRepository) *UserLogic {
	return &UserLogic{
		Logger:         logger,
		DB:             db,
		UserRepository: userRepository,
	}
}

func (u *UserLogic) Create(ctx context.Context, request *model.UserCreateRequest) (*model.UserResponse, error) {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	isUserExist := true
	user := new(storage.User)
	if request.Email != "" {
		user.Email = request.Email
		if err := u.UserRepository.FindByEmail(tx, user, request.Email); err != nil {
			isUserExist = false
		}
	}
	if request.PhoneNumber != "" {
		user.PhoneNumber = request.PhoneNumber
		if err := u.UserRepository.FindByPhoneNumber(tx, user, request.PhoneNumber); err != nil {
			isUserExist = false
		}
	}

	if isUserExist {
		u.Logger.Info(logTag, "user is exist")
		return nil, fiber.ErrForbidden
	}

	user.SafeID = uuid.New().String()
	user.PIN = request.PIN

	if err := u.UserRepository.Create(tx, user); err != nil {
		u.Logger.WithError(err).Error(logTag, "error creating user")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		u.Logger.WithError(err).Error(logTag, "error creating user")
		return nil, fiber.ErrInternalServerError
	}

	return converter.UserToResponse(user), nil
}
