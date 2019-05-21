package repository

import (
	"github.com/1612180/chat_stranger/models"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UserRepoGorm struct {
	db *gorm.DB
}

func NewUserRepoGorm(db *gorm.DB) IUserRepo {
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

func (g *UserRepoGorm) Find(id uint) (*models.User, []error) {
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

func (g *UserRepoGorm) Create(userUpload *models.UserUpload) (uint, []error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userUpload.Authentication.Password), bcrypt.DefaultCost)
	if err != nil {
		var errs []error
		errs = append(errs, err)
		return 0, errs
	}

	user := models.User{
		Credential: models.Credential{
			Name:           userUpload.Authentication.Name,
			HashedPassword: string(hashedPassword),
		},
		FullName:  userUpload.FullName,
		Gender:    userUpload.Gender,
		BirthYear: userUpload.BirthYear,
		Introduce: userUpload.Introduce,
	}

	if errs := g.db.Create(&user).GetErrors(); len(errs) != 0 {
		return 0, errs
	}

	return user.ID, nil
}

func (g *UserRepoGorm) UpdateInfo(id uint, userUpload *models.UserUpload) []error {
	var user models.User
	if errs := g.db.Where("id = ?", id).First(&user).Updates(
		map[string]interface{}{
			"full_name": userUpload.FullName,
		},
	).GetErrors(); len(errs) != 0 {
		return errs
	}
	return nil
}

func (g *UserRepoGorm) UpdatePassword(id uint, authentication *models.Authentication) []error {
	var user models.User
	var credential models.Credential
	if errs := g.db.Where("id = ?", id).First(&user).Related(&credential).GetErrors(); len(errs) != 0 {
		return errs
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(authentication.Password), bcrypt.DefaultCost)
	if err != nil {
		var errs []error
		errs = append(errs, err)
		return errs
	}

	if errs := g.db.Model(&credential).Update("hashed_password", hashedPassword).GetErrors(); len(errs) != 0 {
		return errs
	}

	return nil
}

func (g *UserRepoGorm) Delete(id uint) []error {
	tx := g.db.Begin()

	var user models.User
	var credential models.Credential

	if errs := g.db.Where("id = ?", id).First(&user).GetErrors(); len(errs) != 0 {
		tx.Rollback()
		return errs
	}

	if errs := g.db.Model(&user).Related(&credential).GetErrors(); len(errs) != 0 {
		tx.Rollback()
		return errs
	}

	if errs := g.db.Delete(&user).GetErrors(); len(errs) != 0 {
		tx.Rollback()
		return errs
	}

	if errs := g.db.Delete(&credential).GetErrors(); len(errs) != 0 {
		tx.Rollback()
		return errs
	}

	tx.Commit()
	return nil
}
