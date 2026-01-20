package v1service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	v1dto "vidbox-api/internal/dto/v1"
	"vidbox-api/internal/repository"
	"vidbox-api/internal/utils"
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

type ResponseKKphim struct {
    Data struct {
        Items []struct {
			EpisodeCurrent 	string  `json:"episode_current"`
            OriginName 	string 	`json:"origin_name"`
			Slug 		string 	`json:"slug"`
			Year		int		`json:"year"`
        } `json:"items"`
    } `json:"data"`
}

type ResponseAllKKphim struct {
    Items []struct {
		EpisodeCurrent 	string  `json:"episode_current"`
        OriginName 		string 	`json:"origin_name"`
		Slug 			string 	`json:"slug"`
		Year			int		`json:"year"`
    } `json:"items"`
}

type ResponseTmdb struct {
	Results []struct {
		Id int	`json:"id"`
	} `json:"results"`
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
	for i := 1; i <= 5; i++ {
		
		resp, err := http.Get(fmt.Sprintf("https://ophim1.com/v1/api/danh-sach/phim-moi?page=%d", i))
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
			if item.TMDB.ID != "" && !us.repo.FindSlugExist(item.Slug) {

				if !us.repo.FindMediaByTmdb(item.TMDB.ID, item.TMDB.Season, item.TMDB.Type) {
					log.Printf("🎉 FOUND NEW SLUG:%s\n", utils.RegexOriginalName(item.Slug))
					results = append(results, v1dto.MediaCrawler{
						Type: item.TMDB.Type,
						TMDBID: item.TMDB.ID,
						Season: item.TMDB.Season,
						Slug: item.Slug,
					})
				} else {
					log.Printf("🎉 UPDATE SLUG OPHIM:%s\n", utils.RegexOriginalName(item.Slug))
					if err := us.repo.UpdateOphimSlug(v1dto.MediaCrawler{
						Type: item.TMDB.Type,
						TMDBID: item.TMDB.ID,
						Season: item.TMDB.Season,
						Slug: item.Slug,
					}); err != nil {
						return err
					}
				}
			}
		}

		if err := us.repo.StoreCrawler(results); err != nil {
			return err
		}
	}

	return nil
}

func (us *userService) CrawlerTvKkphim() error {
	for i := 1; i <= 120; i++ {
		log.Printf("CRAWLER PAGE:%d", i)
		resp, err := http.Get(fmt.Sprintf("https://phimapi.com/v1/api/danh-sach/phim-bo?limit=64&page=%d", i))
		if err != nil {
			log.Printf("CRAWLER ERR PAGE:%d\n", i)
		}
		defer resp.Body.Close()

		var raw ResponseKKphim
		if err := json.NewDecoder(resp.Body).Decode((&raw)); err != nil {
			log.Printf("DECODE ERR PAGE:%d\n", i)
		}

		var results []v1dto.MediaCrawlerKKphim

		for _, item := range raw.Data.Items {
			if us.repo.FindSlugExistKKphim(item.Slug) {
				log.Printf("🎯 SLUG EXTING:%s\n", item.Slug)
				continue
			}
			tmdbId := us.SearchTv(utils.RegexOriginalName(item.OriginName), item.Year)
			if tmdbId != 0 {
				//log.Printf("🎉 FOUND NEW SLUG:%s - %d\n", utils.RegexOriginalName(item.OriginName), tmdbId)
				results = append(results, v1dto.MediaCrawlerKKphim{
					Type: "tv",
					Season: utils.ExtractNumber(item.Slug),
					TMDBID: strconv.Itoa(tmdbId),
					Slug: item.Slug,
				})
				time.Sleep(500 * time.Millisecond)
			}
		}
		
		if err := us.repo.StoreCrawlerKKphim(results); err != nil {
			return err
		}
	}

	return nil
}

func (us *userService) CrawlerMovieKkphim() error {
	for i := 52; i <= 236; i++ {
		log.Printf("CRAWLER PAGE:%d", i)
		resp, err := http.Get(fmt.Sprintf("https://phimapi.com/v1/api/danh-sach/phim-le?limit=64&page=%d", i))
		if err != nil {
			log.Printf("CRAWLER ERR PAGE:%d\n", i)
		}
		defer resp.Body.Close()

		var raw ResponseKKphim
		if err := json.NewDecoder(resp.Body).Decode((&raw)); err != nil {
			log.Printf("DECODE ERR PAGE:%d\n", i)
		}

		var results []v1dto.MediaCrawlerKKphim

		for _, item := range raw.Data.Items {
			if us.repo.FindSlugExistKKphim(item.Slug) {
				log.Printf("🎯 SLUG EXTING:%s\n", item.Slug)
				continue
			}
			tmdbId := us.SearchMovie(utils.RegexOriginalName(item.OriginName), item.Year)
			if tmdbId != 0 {
				//log.Printf("🎉 FOUND NEW SLUG:%s - %d\n", item.Slug, tmdbId)
				results = append(results, v1dto.MediaCrawlerKKphim{
					Type: "movie",
					Season: nil,
					TMDBID: strconv.Itoa(tmdbId),
					Slug: item.Slug,
				})

				time.Sleep(500 * time.Millisecond)
			}
		}
		
		if err := us.repo.StoreCrawlerKKphim(results); err != nil {
			return err
		}
	}

	return nil
}

func (us *userService) CrawlerAllKKphim() error {
	for i := 1; i <= 5; i++ {
		log.Printf("CRAWLER PAGE:%d", i)
		resp, err := http.Get(fmt.Sprintf("https://phimapi.com/danh-sach/phim-moi-cap-nhat-v3?page=%d", i))
		if err != nil {
			log.Printf("CRAWLER ERR PAGE:%d\n", i)
		}
		defer resp.Body.Close()

		var raw ResponseAllKKphim
		
		if err := json.NewDecoder(resp.Body).Decode((&raw)); err != nil {
			log.Printf("DECODE ERR PAGE:%d\n", i)
		}

		var results []v1dto.MediaCrawlerKKphim

		for _, item := range raw.Items {

			if us.repo.FindSlugExistKKphim(item.Slug) {
				log.Printf("🎯 SLUG EXTING:%s\n", item.Slug)
				continue
			}

			if item.EpisodeCurrent == "Full" {
				tmdbId := us.SearchMovie(utils.RegexOriginalName(item.OriginName), item.Year)
				if tmdbId != 0 {
					log.Printf("🎉 FOUND NEW SLUG MOVIE:%s - %d\n", utils.RegexOriginalName(item.OriginName), tmdbId)
					results = append(results, v1dto.MediaCrawlerKKphim{
						Type: "movie",
						Season: nil,
						TMDBID: strconv.Itoa(tmdbId),
						Slug: item.Slug,
					})
					time.Sleep(300 * time.Millisecond)
				}
			} else {
				tmdbId := us.SearchTv(utils.RegexOriginalName(item.OriginName), item.Year)
				if tmdbId != 0 {
					log.Printf("🎉 FOUND NEW SLUG TV:%s - %d\n", utils.RegexOriginalName(item.OriginName), tmdbId)
					results = append(results, v1dto.MediaCrawlerKKphim{
						Type: "tv",
						Season: utils.ExtractNumber(item.Slug),
						TMDBID: strconv.Itoa(tmdbId),
						Slug: item.Slug,
					})
					time.Sleep(300 * time.Millisecond)
				}
			}
		}

		if err := us.repo.StoreCrawlerKKphim(results); err != nil {
			return err
		}
	}

	return nil
}

func (us *userService) SearchTv(name string, year int) int {
	resp, err := http.Get(fmt.Sprintf("https://api.themoviedb.org/3/search/tv?api_key=ef311eb0b9b07b9c73e9fb0a732cc150&query=%v&include_adult=false&language=en-US&page=1&year=%d", name, year))
	if err != nil {
		log.Printf("⛔ SEARCH TV ERR:%s\n", err)
		return 0
	}
	defer resp.Body.Close()

	var raw ResponseTmdb
	if err := json.NewDecoder(resp.Body).Decode((&raw)); err != nil {
		
		log.Printf("⛔ DECODE TV ERR:%s\n", err)
		log.Printf("⛔ Name Search ERR:%s\n", name)
		log.Printf("⛔ BODY ERR:%s\n", resp.Body)
		return 0
	}

	if len(raw.Results) > 0 {
		return raw.Results[0].Id	
	}
	log.Printf("⛔ Name Search Not Found:%s\n", name)
	return 0
}

func (us *userService) SearchMovie(name string, year int) int {
	resp, err := http.Get(fmt.Sprintf("https://api.themoviedb.org/3/search/movie?api_key=ef311eb0b9b07b9c73e9fb0a732cc150&query=%s&include_adult=false&language=en-US&page=1&year=%d", name, year))
	if err != nil {
		log.Printf("⛔ SEARCH Movie ERR:%s\n", err)
		return 0
	}
	defer resp.Body.Close()

	var raw ResponseTmdb
	if err := json.NewDecoder(resp.Body).Decode((&raw)); err != nil {
		
		log.Printf("⛔ DECODE TV ERR:%s\n", err)
		log.Printf("⛔ Name Search ERR:%s\n", name)
		log.Printf("⛔ BODY ERR:%s\n", resp.Body)
		return 0
	}

	if len(raw.Results) > 0 {
		return raw.Results[0].Id	
	}
	log.Printf("⛔ Name Search Not Found:%s\n", name)
	return 0
}