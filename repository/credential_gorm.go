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

func (credentialRepoGorm *CredentialRepoGorm) FetchAll() ([]*models.Credential, error) {
	return nil, nil
}

func (credentialRepoGorm *CredentialRepoGorm) FindByID(id uint) (*models.Credential, error) {
	return nil, nil
}

func (credentialRepoGorm *CredentialRepoGorm) FindByName(name string) (*models.Credential, error) {
	var credential models.Credential
	err := credentialRepoGorm.db.Where("name = ?", name).First(&credential).Error
	if err != nil {
		return nil, err
	}
	return &credential, nil
}

func (credentialRepoGorm *CredentialRepoGorm) Create(*models.Credential) error {
	return nil
}

func (credentialRepoGorm *CredentialRepoGorm) Update(*models.Credential, *models.Credential) error {
	return nil
}

func (credentialRepoGorm *CredentialRepoGorm) Delete(*models.Credential) error {
	return nil
}
