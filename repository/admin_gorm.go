package repository

import (
	"github.com/1612180/chat_stranger/models"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type AdminRepoGorm struct {
	db *gorm.DB
}

func NewAdminRepoGorm(db *gorm.DB) IAdminRepo {
	db.DropTableIfExists(&models.Admin{})
	db.AutoMigrate(&models.Admin{})

	return &AdminRepoGorm{db: db}
}

func (g *AdminRepoGorm) FetchAll() ([]*models.Admin, []error) {
	var admins []*models.Admin
	errs := g.db.Find(&admins).GetErrors()
	if len(errs) != 0 {
		return nil, errs
	}

	for _, admin := range admins {
		if errs := g.db.Model(admin).Related(&admin.Credential).GetErrors(); len(errs) != 0 {
			return nil, errs
		}
	}

	return admins, nil
}

func (g *AdminRepoGorm) Find(id uint) (*models.Admin, []error) {
	var admin models.Admin
	errs := g.db.Where("id = ?", id).First(&admin).GetErrors()
	if len(errs) != 0 {
		return nil, errs
	}

	if errs = g.db.Model(&admin).Related(&admin.Credential).GetErrors(); len(errs) != 0 {
		return nil, errs
	}

	return &admin, nil
}

func (g *AdminRepoGorm) Create(adminUpload *models.AdminUpload) (uint, []error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminUpload.Authentication.Password), bcrypt.DefaultCost)
	if err != nil {
		var errs []error
		errs = append(errs, err)
		return 0, errs
	}

	admin := models.Admin{
		Credential: models.Credential{
			Name:           adminUpload.Authentication.Name,
			HashedPassword: string(hashedPassword),
		},
		FullName: adminUpload.FullName,
	}

	if errs := g.db.Create(&admin).GetErrors(); len(errs) != 0 {
		return 0, errs
	}

	return admin.ID, nil
}

func (g *AdminRepoGorm) UpdateInfo(id uint, adminUpload *models.AdminUpload) []error {
	var admin models.Admin
	if errs := g.db.Where("id = ?", id).First(&admin).Updates(
		map[string]interface{}{
			"full_name": adminUpload.FullName,
		},
	).GetErrors(); len(errs) != 0 {
		return errs
	}

	return nil
}

func (g *AdminRepoGorm) UpdatePassword(id uint, authentication *models.Authentication) []error {
	var admin models.Admin
	var credential models.Credential
	errs := g.db.Where("id = ?", id).First(&admin).Related(&credential).GetErrors()
	if len(errs) != 0 {
		return errs
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(authentication.Password), bcrypt.DefaultCost)
	if err != nil {
		var errs []error
		errs = append(errs, err)
		return errs
	}

	if errs = g.db.Model(&credential).Update("hashed_password", hashedPassword).GetErrors(); len(errs) != 0 {
		return errs
	}

	return nil
}

func (g *AdminRepoGorm) Delete(id uint) []error {
	tx := g.db.Begin()

	var admin models.Admin
	var credential models.Credential

	if errs := g.db.Where("id = ?", id).First(&admin).GetErrors(); len(errs) != 0 {
		tx.Rollback()
		return errs
	}

	if errs := g.db.Model(&admin).Related(&credential).GetErrors(); len(errs) != 0 {
		tx.Rollback()
		return errs
	}

	if errs := g.db.Delete(&admin).GetErrors(); len(errs) != 0 {
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
