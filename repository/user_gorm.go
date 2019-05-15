package repository

import (
	"github.com/1612180/chat_stranger/models"
	"github.com/jinzhu/gorm"
)

type gormUserRepo struct {
	db *gorm.DB
}

// Return implement of UserRepo interface
func NewGormUserRepo(db *gorm.DB) IUserRepo {
	db.DropTableIfExists(&models.User{})

	// Migrate
	db.AutoMigrate(&models.User{})

	// db.Create(&models.User{Name: "A"})
	// db.Create(&models.User{Name: "B"})

	return &gormUserRepo{db: db}
}

func (g *gormUserRepo) FetchAll() ([]*models.User, error) {
	var users []*models.User
	err := g.db.Find(&users).Error
	if err != nil {
		return nil, err
	} else {
		return users, nil
	}
}

func (g *gormUserRepo) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := g.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	} else {
		return &user, nil
	}
}

func (g *gormUserRepo) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := g.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	} else {
		return &user, nil
	}

}

func (g *gormUserRepo) Create(u *models.User) error {
	err := g.db.Create(u).Error
	return err
}

func (g *gormUserRepo) Update(uOld *models.User, uNew *models.User) error {
	uOld.Username = uNew.Username
	uOld.PasswordHash = uNew.PasswordHash
	err := g.db.Save(uOld).Error
	return err
}

func (g *gormUserRepo) Delete(u *models.User) error {
	err := g.db.Delete(u).Error
	return err
}
