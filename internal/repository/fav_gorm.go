package repository

import (
	"github.com/1612180/chat_stranger/internal/models"
	"github.com/jinzhu/gorm"
)

type FavoriteRepoGorm struct {
	db *gorm.DB
}

func NewFavoriteRepoGorm(db *gorm.DB) FavoriteRepo {
	db.DropTableIfExists(&models.Favorite{})
	db.AutoMigrate(&models.Favorite{})

	return &FavoriteRepoGorm{db: db}
}

func (g *FavoriteRepoGorm) FetchAll() ([]*models.Favorite, []error) {
	var favs []*models.Favorite
	if errs := g.db.Find(&favs).GetErrors(); len(errs) != 0 {
		return nil, errs
	}

	return favs, nil
}

func (g *FavoriteRepoGorm) Find(name string) (*models.Favorite, []error) {
	var fav models.Favorite
	if errs := g.db.Where("name = ?", name).First(&fav).GetErrors(); len(errs) != 0 {
		return nil, errs
	}

	return &fav, nil
}

func (g *FavoriteRepoGorm) Create(fav *models.Favorite) (int, []error) {
	if errs := g.db.Create(fav).GetErrors(); len(errs) != 0 {
		return 0, errs
	}

	return fav.ID, nil
}

func (g *FavoriteRepoGorm) Delete(id int) []error {
	tx := g.db.Begin()

	var fav models.Favorite

	if errs := tx.Where("id = ?", id).First(&fav).GetErrors(); len(errs) != 0 {
		tx.Rollback()
		return errs
	}

	if errs := tx.Delete(&fav).GetErrors(); len(errs) != 0 {
		tx.Rollback()
		return errs
	}

	tx.Commit()
	return nil
}
