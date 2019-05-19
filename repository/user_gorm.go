package repository

import (
	"github.com/1612180/chat_stranger/models"
	"github.com/jinzhu/gorm"
)

type UserRepoGorm struct {
	db *gorm.DB
}

func NewUserRepoGorm(db *gorm.DB) IUserRepo {
	db.DropTableIfExists(&models.User{})
	db.CreateTable(&models.User{})

	return &UserRepoGorm{db: db}
}

func (g *UserRepoGorm) FetchAll() ([]*models.User, []error) {
	var users []*models.User
	errs := g.db.Find(&users).GetErrors()
	if len(errs) != 0 {
		return nil, errs
	} else {
		for _, user := range users {
			errs := g.db.Model(user).Related(&user.Credential).GetErrors()
			if len(errs) != 0 {
				return nil, errs
			}
		}
		return users, nil
	}
}

func (g *UserRepoGorm) FindByID(id uint) (*models.User, []error) {
	var user models.User
	errs := g.db.Where("id = ?", id).First(&user).GetErrors()
	if len(errs) != 0 {
		return nil, errs
	} else {
		errs := g.db.Model(&user).Related(&user.Credential).GetErrors()
		if len(errs) != 0 {
			return nil, errs
		}
		return &user, nil
	}
}

func (g *UserRepoGorm) Create(user *models.User) []error {
	errs := g.db.Create(user).GetErrors()
	if len(errs) != 0 {
		return errs
	}
	return nil
}

func (g *UserRepoGorm) UpdateInfo(userOld *models.User, userNew *models.User) []error {
	userOld.FullName = userNew.FullName

	errs := g.db.Save(userOld).GetErrors()
	if len(errs) != 0 {
		return errs
	}
	return nil
}

func (g *UserRepoGorm) Delete(user *models.User) []error {
	errs := g.db.Delete(user).GetErrors()
	if len(errs) != 0 {
		return errs
	}
	return nil
}
