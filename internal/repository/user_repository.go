package repository

import (
	v1dto "vidbox-api/internal/dto/v1"
	"vidbox-api/internal/models"

	"github.com/doug-martin/goqu/v9"
)

type SqlUserRepository struct {
	users []models.User
	db *goqu.Database
}

func NewSqlUserRepository(DB *goqu.Database) UserRepository {
	return &SqlUserRepository{
		users : make([]models.User, 0),
		db: DB,
	}
}

func (ur *SqlUserRepository) StoreCrawler(media []v1dto.MediaCrawler) error {
	if len(media) > 0 {
		_, err := ur.db.Insert(goqu.T("medias")).Rows(media).Executor().Exec()
		if err != nil {
			return err
		}
	}
	
	return nil
}

func (ur *SqlUserRepository) FindSlugExist(slug string) bool{

	var slugOphim string

	ds := ur.db.From(goqu.T("medias")).Where(
		goqu.C("ophim_slug").Eq(slug),
	)

	found, err := ds.ScanVal(&slugOphim)
	if found {
		return true
	}

	if err != nil {
		return true
	}

	return false
}

func (ur *SqlUserRepository) FindMediaByTmdb(tmdbId string, season *int, mediaType string) bool {
	var slug string

	ds := ur.db.From(goqu.T("medias"))
	switch mediaType {
	case "tv":
		ds = ds.Where(
			goqu.C("tmdb_id").Eq(tmdbId),
			goqu.C("season").Eq(season),
			goqu.C("media_type").Eq(mediaType),
		)
	default:
		ds = ds.Where(
			goqu.C("tmdb_id").Eq(tmdbId),
			goqu.C("season").IsNull(),
			goqu.C("media_type").Eq(mediaType),
		)
	}
	
	found, err := ds.ScanVal(&slug)
	if found {
		return true
	}

	if err != nil {
		return true
	}

	return false
}

func (ur *SqlUserRepository) FindSlugExistKKphim(slug string) bool{

	var slugKKphim string

	ds := ur.db.From(goqu.T("provider_kkphims")).Where(
		goqu.C("kkphim_slug").Eq(slug),
	)

	found, err := ds.ScanVal(&slugKKphim)
	if found {
		return true
	}

	if err != nil {
		return true
	}

	return false
}

func (ur *SqlUserRepository) FindSlugExistNguonC(slug string) bool{

	var slugNguonCphim string

	ds := ur.db.From(goqu.T("provider_nguoncs")).Where(
		goqu.C("nguonc_slug").Eq(slug),
	)

	found, err := ds.ScanVal(&slugNguonCphim)
	if found {
		return true
	}

	if err != nil {
		return true
	}

	return false
}

func (ur *SqlUserRepository) StoreCrawlerKKphim(media []v1dto.MediaCrawlerKKphim) error {

	if len(media) > 0 {
		_, err := ur.db.Insert(goqu.T("provider_kkphims")).Rows(media).Executor().Exec()
		if err != nil {
			return err
		}
	}
	
	return nil
}

func (ur *SqlUserRepository) StoreCrawlerNguonC(media []v1dto.MediaCrawlerNguonC) error {

	if len(media) > 0 {
		_, err := ur.db.Insert(goqu.T("provider_nguoncs")).Rows(media).Executor().Exec()
		if err != nil {
			return err
		}
	}
	
	return nil
}

func (ur *SqlUserRepository) UpdateOphimSlug(media v1dto.MediaCrawler) error {
	
	ds := ur.db.Update(goqu.T("medias")).Set(goqu.Record{
		"ophim_slug": media.Slug,
	})

	switch media.Type {
	case "tv":
		ds = ds.Where(
			goqu.C("tmdb_id").Eq(media.TMDBID),
			goqu.C("season").Eq(media.Season),
			goqu.C("media_type").Eq(media.Type),
		)
	default:
		ds = ds.Where(
			goqu.C("tmdb_id").Eq(media.TMDBID),
			goqu.C("season").IsNull(),
			goqu.C("media_type").Eq(media.Type),
		)
	}

	_, err := ds.Executor().Exec()
	if err != nil {
		return err
	}

	return nil
}