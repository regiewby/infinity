package converter

import (
	"github.com/infinity/identity-service/internal/model"
	"github.com/infinity/identity-service/internal/storage"
)

func UserToResponse(user *storage.User) *model.UserResponse {
	return &model.UserResponse{
		ID:          user.ID,
		SafeID:      user.SafeID,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
}
