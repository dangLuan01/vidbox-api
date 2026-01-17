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