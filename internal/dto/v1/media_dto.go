package v1dto

type MediaInput struct {
	TMDBID 		int `uri:"tmdb_id" binding:"required"`
	Season 		int `uri:"season"`
	MediaType 	string `json:"media_type"`
}

type MediaCrawler struct {
	Type 	string 	`db:"media_type" json:"type"`
	TMDBID 	string 	`db:"tmdb_id" json:"id"`
	Season 	*int	`db:"season" json:"season"`
	Slug	string	`db:"ophim_slug" json:"slug"`
}