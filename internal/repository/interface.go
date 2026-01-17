package repository

import (
	v1dto "vidbox-api/internal/dto/v1"
)

type UserRepository interface {
	// FindAll() ([]models.User, error)
	// FindBYUUID(uuid uuid.UUID) (models.User, error)
	// Create(user models.User) error
	// Update(uuid uuid.UUID, user models.User) error
	// Delete(uuid uuid.UUID) error
	// FindByEmail(email string) (models.User, error)
	// UpdatePassword(uuid uuid.UUID, password string) error
	StoreCrawler(media []v1dto.MediaCrawler) error
	FindSlugExist(slug string) bool
}

type MediaRepository interface {
	FindByTMDBID(params v1dto.MediaInput) (string, error)
}