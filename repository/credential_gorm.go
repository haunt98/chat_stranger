package repository

import (
	"github.com/1612180/chat_stranger/models"
	"github.com/jinzhu/gorm"
)

type CredentialRepoGorm struct {
	db *gorm.DB
}

func NewCredentialRepoGorm(db *gorm.DB) *CredentialRepoGorm {
	db.DropTableIfExists(&models.Credential{})
	db.CreateTable(&models.Credential{})

	return &CredentialRepoGorm{db: db}
}

func (g *CredentialRepoGorm) FetchAll() ([]*models.Credential, []error) {
	return nil, nil
}

func (g *CredentialRepoGorm) FindByID(id uint) (*models.Credential, []error) {
	return nil, nil
}

func (g *CredentialRepoGorm) FindByName(name string) (*models.Credential, []error) {
	var credential models.Credential
	errs := g.db.Where("name = ?", name).First(&credential).GetErrors()
	if len(errs) != 0 {
		return nil, errs
	}
	return &credential, nil
}

func (g *CredentialRepoGorm) FindByUser(user *models.User) (*models.Credential, []error) {
	var credential models.Credential
	errs := g.db.Model(user).Related(&credential).GetErrors()
	if len(errs) != 0 {
		return nil, errs
	}

	return &credential, nil
}

func (g *CredentialRepoGorm) Create(*models.Credential) (bool, []error) {
	return false, nil
}

func (g *CredentialRepoGorm) UpdatePassword(credentialOld *models.Credential, credentialNew *models.Credential) (bool, []error) {
	credentialOld.HashedPassword = credentialNew.HashedPassword
	errs := g.db.Save(credentialOld).GetErrors()
	if len(errs) != 0 {
		return false, errs
	}
	return true, nil
}

func (g *CredentialRepoGorm) Delete(credential *models.Credential) (bool, []error) {
	errs := g.db.Delete(credential).GetErrors()
	if len(errs) != 0 {
		return false, errs
	}
	return true, nil
}
