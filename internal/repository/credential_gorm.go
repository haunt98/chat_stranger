package repository

import (
	"github.com/1612180/chat_stranger/internal/models"
	"github.com/jinzhu/gorm"
)

type CredentialRepoGorm struct {
	db *gorm.DB
}

func NewCredentialRepoGorm(db *gorm.DB) CredentialRepo {
	db.DropTableIfExists(&models.Credential{})
	db.AutoMigrate(&models.Credential{})

	return &CredentialRepoGorm{db: db}
}

func (g *CredentialRepoGorm) Find(name string) (*models.Credential, []error) {
	var credential models.Credential
	errs := g.db.Where("reg_name = ?", name).First(&credential).GetErrors()
	if len(errs) != 0 {
		return nil, errs
	}

	return &credential, nil
}

func (g *CredentialRepoGorm) TryAdmin(cre *models.Credential) (*models.Admin, []error) {
	var admin models.Admin
	errs := g.db.Model(cre).Related(&admin).GetErrors()
	if len(errs) != 0 {
		return nil, errs
	}

	return &admin, nil
}

func (g *CredentialRepoGorm) TryUser(cre *models.Credential) (*models.User, []error) {
	var user models.User
	errs := g.db.Model(cre).Related(&user).GetErrors()
	if len(errs) != 0 {
		return nil, errs
	}

	return &user, nil
}
