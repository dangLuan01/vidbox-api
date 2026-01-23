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

func (mr *SqlMediaRepository) FindByTMDBID(params v1dto.MediaInput) (v1dto.MediaOutput, error) {
	var slug v1dto.MediaOutput

	ds := mr.db.From(goqu.T("medias")).Select(
		goqu.C("ophim_slug"),
		goqu.C("kkphim_slug"),
		goqu.C("nguonc_slug"),
	)

	if params.MediaType == "tv" {
		ds = ds.Where(
			goqu.C("tmdb_id").Eq(params.TMDBID),
			goqu.C("season").Eq(params.Season),
		)	
	} else {
		ds = ds.Where(
			goqu.C("tmdb_id").Eq(params.TMDBID),
			goqu.C("season").IsNull(),
		)
	}

	found, err := ds.ScanStruct(&slug)
	if err != nil {
		return v1dto.MediaOutput{}, utils.WrapError(string(utils.ErrCodeInternal), "Error", err)
	}

	if !found {
		return v1dto.MediaOutput{}, utils.NewError(string(utils.ErrCodeNotFound), "Not found.")
	}

	return slug, nil
}