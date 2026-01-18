package v1service

import (
	v1dto "vidbox-api/internal/dto/v1"
)

type UserService interface {
	// GetAllUser() ([]models.User, error)
	// GetUserByUUID(uuid uuid.UUID) (models.User, error)
	// CreateUser(user models.User) (models.User, error)
	// UpdateUser(uuid uuid.UUID, user models.User) (models.User, error)
	// DeleteUser(uuid uuid.UUID) error
	Crawler() error
}

type MediaService interface {
	GetMedia(params v1dto.MediaInput) (v1dto.MediaOutput, error)
}