package v1service

import (
	v1dto "vidbox-api/internal/dto/v1"
	"vidbox-api/internal/repository"
)

type mediaService struct {
	repo repository.MediaRepository
}

func NewMediaService(repo repository.MediaRepository) MediaService {
	return &mediaService{
		repo: repo,
	}
}

func (ms *mediaService) GetMedia(params v1dto.MediaInput) (string, error) {
	slug, err := ms.repo.FindByTMDBID(params)
	if err != nil {
		return "", err
	}

	return slug, nil
}