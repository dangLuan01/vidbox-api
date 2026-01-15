package repository

import (
	v1dto "vidbox-api/internal/dto/v1"
	"vidbox-api/internal/utils"

	"github.com/doug-martin/goqu/v9"
)

type SqlMediaRepository struct {
	db *goqu.Database
}

func NewSqlMediaRepository(DB *goqu.Database) MediaRepository {
	return &SqlMediaRepository{
		db: DB,
	}
}

func (mr *SqlMediaRepository) FindByTMDBID(params v1dto.MediaInput) (string, error) {
	var slug string

	ds := mr.db.From(goqu.T("medias")).Select(
		goqu.C("ophim_slug"),
	).Where(
		goqu.C("tmdb_id").Eq(params.TMDBID),
		goqu.C("season").Eq(params.Season),
		goqu.C("media_type").Eq(params.MediaType),
	)

	found, err := ds.ScanVal(&slug)
	if err != nil {
		return "", utils.WrapError(string(utils.ErrCodeInternal), "Error", err)
	}

	if !found {
		return "", utils.NewError(string(utils.ErrCodeNotFound), "Not found.")
	}

	return slug, nil
}