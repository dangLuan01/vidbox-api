package repository

import (
	v1dto "vidbox-api/internal/dto/v1"
)

type UserRepository interface {
	StoreCrawler(media []v1dto.MediaCrawler) error
	FindSlugExist(slug string) bool
	FindSlugExistKKphim(slug string) bool
	StoreCrawlerKKphim(media []v1dto.MediaCrawlerKKphim) error
	StoreCrawlerNguonC(media []v1dto.MediaCrawlerNguonC) error
	FindMediaByTmdb(tmdbId string, season *int, mediaType string) bool
	UpdateOphimSlug(media v1dto.MediaCrawler) error
}

type MediaRepository interface {
	FindByTMDBID(params v1dto.MediaInput) (v1dto.MediaOutput, error)
}