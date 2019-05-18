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

func (g *UserRepoGorm) FetchAll() ([]*models.User, error) {
	var users []*models.User
	err := g.db.Find(&users).Error
	if err != nil {
		return nil, err
	} else {
		for _, user := range users {
			err := g.db.Model(user).Related(&user.Credential).Error
			if err != nil {
				return nil, err
			}
		}
		return users, nil
	}
}

func (g *UserRepoGorm) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := g.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	} else {
		err := g.db.Model(&user).Related(&user.Credential).Error
		if err != nil {
			return nil, err
		}
		return &user, nil
	}
}

func (g *UserRepoGorm) Create(u *models.User) error {
	err := g.db.Create(u).Error
	return err
}

func (g *UserRepoGorm) Update(uOld *models.User, uNew *models.User) error {
	uOld.Credential = uNew.Credential
	err := g.db.Save(uOld).Error
	return err
}

func (g *UserRepoGorm) Delete(u *models.User) error {
	err := g.db.Delete(u).Error
	return err
}
