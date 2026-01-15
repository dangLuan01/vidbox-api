package v1dto

type MediaInput struct {
	TMDBID 		int `uri:"tmdb_id" binding:"required"`
	Season 		int `uri:"season" binding:"required"`
	MediaType 	string `uri:"media_type" binding:"required"`
}