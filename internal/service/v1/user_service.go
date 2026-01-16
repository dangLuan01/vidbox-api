package v1service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	v1dto "vidbox-api/internal/dto/v1"
	"vidbox-api/internal/repository"
)

type Response struct {
    Data struct {
        Items []struct {
            TMDB struct {
                Type   string `json:"type"`
                ID     string `json:"id"`
                Season *int   `json:"season"`
            } `json:"tmdb"`
            Slug string `json:"slug"`
        } `json:"items"`
    } `json:"data"`
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (us *userService) Crawler() error {
	for i := 1; i <= 167; i++ {
		fmt.Printf("DANG CRAWLER PAGE:%d\n", i)
		resp, err := http.Get(fmt.Sprintf("https://ophim1.com/v1/api/danh-sach/hoat-hinh?page=%d", i))
		if err != nil {
			log.Printf("CRAWLER ERR PAGE:%d\n", i)
		}
		defer resp.Body.Close()

		var raw Response
		if err := json.NewDecoder(resp.Body).Decode((&raw)); err != nil {
			log.Printf("DECODE ERR PAGE:%d\n", i)
		}

		var results []v1dto.MediaCrawler

		for _, item := range raw.Data.Items {
			if item.TMDB.ID != "" {
				results = append(results, v1dto.MediaCrawler{
					Type: item.TMDB.Type,
					TMDBID: item.TMDB.ID,
					Season: item.TMDB.Season,
					Slug: item.Slug,
				})
			}
		}

		if err := us.repo.StoreCrawler(results); err != nil {
			return err
		}
	}

	return nil
}