package v1service

import (
	v1dto "vidbox-api/internal/dto/v1"
	"vidbox-api/internal/models"

	"github.com/google/uuid"
)

type UserService interface {
	GetAllUser() ([]models.User, error)
	GetUserByUUID(uuid uuid.UUID) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(uuid uuid.UUID, user models.User) (models.User, error)
	DeleteUser(uuid uuid.UUID) error
}

type MediaService interface {
	GetMedia(params v1dto.MediaInput) (string, error)
}