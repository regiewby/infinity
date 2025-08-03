package handlers

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/infinity/identity-service/internal/logic"
	"github.com/infinity/identity-service/internal/model"
	"github.com/sirupsen/logrus"
)

const logTag = "user.handler"

type UserHandler struct {
	Logger    *logrus.Logger
	Validator *validator.Validate
	Logic     *logic.UserLogic
}

func NewUserHandler(logger *logrus.Logger, validator *validator.Validate, logic *logic.UserLogic) *UserHandler {
	return &UserHandler{
		Logic:     logic,
		Logger:    logger,
		Validator: validator,
	}
}

func (u *UserHandler) Create(ctx *fiber.Ctx) error {
	request := new(model.UserCreateRequest)
	if err := ctx.BodyParser(request); err != nil {
		u.Logger.WithError(err).Error(logTag, "parsing body error: %s", err)
		return fiber.ErrBadRequest
	}

	if err := u.Validator.Struct(request); err != nil {
		u.Logger.WithError(err).Error(logTag, "validating error: %s", err)
		return fiber.ErrBadRequest
	}

	response, err := u.Logic.Create(ctx.UserContext(), request)
	if err != nil {
		u.Logger.WithError(err).Error(logTag, "creating user error: %s", err)
		return fiber.ErrInternalServerError
	}

	u.Logger.Info(logTag, "creating user success")
	return ctx.JSON(model.WebResponse[*model.UserResponse]{Data: response})
}
