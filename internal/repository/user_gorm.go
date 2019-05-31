package repository

import (
	"github.com/1612180/chat_stranger/internal/models"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
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

func (g *UserRepoGorm) Create(upload *models.UserUpload) (uint, []error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(upload.Password), bcrypt.DefaultCost)
	if err != nil {
		var errs []error
		errs = append(errs, err)
		return 0, errs
	}

	user := models.User{
		Credential: models.Credential{
			Name:           upload.Name,
			HashedPassword: string(hashedPassword),
		},
		Fullname:  upload.FullName,
		Gender:    upload.Gender,
		Birthyear: upload.BirthYear,
		Introduce: upload.Introduce,
	}

	if errs := g.db.Create(&user).GetErrors(); len(errs) != 0 {
		return 0, errs
	}

	return user.ID, nil
}

func (g *UserRepoGorm) UpdateInfo(id uint, upload *models.UserUpload) []error {
	var user models.User
	if errs := g.db.Where("id = ?", id).First(&user).Updates(
		map[string]interface{}{
			"full_name": upload.FullName,
		},
	).GetErrors(); len(errs) != 0 {
		return errs
	}
	return nil
}

func (g *UserRepoGorm) UpdatePassword(id uint, auth *models.Authentication) []error {
	var user models.User
	var cre models.Credential
	if errs := g.db.Where("id = ?", id).First(&user).GetErrors(); len(errs) != 0 {
		return errs
	}

	if errs := g.db.Model(&user).Related(&cre).GetErrors(); len(errs) != 0 {
		return errs
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(auth.Password), bcrypt.DefaultCost)
	if err != nil {
		var errs []error
		errs = append(errs, err)
		return errs
	}

	if errs := g.db.Model(&cre).Update("hashed_password", hashedPassword).GetErrors(); len(errs) != 0 {
		return errs
	}

	return nil
}

func (g *UserRepoGorm) Delete(id uint) []error {
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
