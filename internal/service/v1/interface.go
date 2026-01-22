package v1service

import (
	v1dto "vidbox-api/internal/dto/v1"
)

type UserService interface {
	Crawler() error
	CrawlerTvKkphim() error
	CrawlerMovieKkphim() error
	CrawlerAllKKphim() error
	CrawlerTvNguonC() error
	CrawlerMovieNguonC() error
}

type MediaService interface {
	GetMedia(params v1dto.MediaInput) (v1dto.MediaOutput, error)
}