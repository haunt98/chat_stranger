package repository

import (
	"github.com/1612180/chat_stranger/internal/models"
	"github.com/jinzhu/gorm"
)

type UserRepoGorm struct {
	db *gorm.DB
}

func NewUserRepoGorm(db *gorm.DB) UserRepo {
	db.DropTableIfExists(&models.User{})
	db.AutoMigrate(&models.User{})

	return &UserRepoGorm{db: db}
}

func (g *UserRepoGorm) FetchAll() ([]*models.User, []error) {
	var users []*models.User
	errs := g.db.Find(&users).GetErrors()
	if len(errs) != 0 {
		return nil, errs
	}

	for _, user := range users {
		if errs := g.db.Model(user).Related(&user.Credential).GetErrors(); len(errs) != 0 {
			return nil, errs
		}
	}

	return users, nil
}

func (g *UserRepoGorm) Find(id int) (*models.User, []error) {
	var user models.User
	errs := g.db.Where("id = ?", id).First(&user).GetErrors()
	if len(errs) != 0 {
		return nil, errs
	}

	if errs = g.db.Model(&user).Related(&user.Credential).GetErrors(); len(errs) != 0 {
		return nil, errs
	}

	return &user, nil
}

func (g *UserRepoGorm) Create(user *models.User) (int, []error) {
	if errs := g.db.Create(user).GetErrors(); len(errs) != 0 {
		return 0, errs
	}

	return user.ID, nil
}

func (g *UserRepoGorm) UpdateInfo(id int, user *models.User) []error {
	if errs := g.db.Table("users").Where("id = ?", id).Updates(
		map[string]interface{}{
			"full_name":  user.FullName,
			"gender":     user.Gender,
			"birth_year": user.BirthYear,
			"introduce":  user.Introduce,
		},
	).GetErrors(); len(errs) != 0 {
		return errs
	}
	return nil
}

func (g *UserRepoGorm) Delete(id int) []error {
	tx := g.db.Begin()

	var user models.User
	var cre models.Credential

	if errs := tx.Where("id = ?", id).First(&user).GetErrors(); len(errs) != 0 {
		tx.Rollback()
		return errs
	}

	if errs := tx.Model(&user).Related(&cre).GetErrors(); len(errs) != 0 {
		tx.Rollback()
		return errs
	}

	if errs := tx.Delete(&user).GetErrors(); len(errs) != 0 {
		tx.Rollback()
		return errs
	}

	if errs := tx.Delete(&cre).GetErrors(); len(errs) != 0 {
		tx.Rollback()
		return errs
	}

	tx.Commit()
	return nil
}
