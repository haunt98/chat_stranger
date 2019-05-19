package repository

import (
	"github.com/1612180/chat_stranger/models"
	"github.com/jinzhu/gorm"
)

type TransactionRepoGorm struct {
	db *gorm.DB
}

func NewTransactionRepoGorm(db *gorm.DB) *TransactionRepoGorm {
	return &TransactionRepoGorm{
		db: db,
	}
}

func (g *TransactionRepoGorm) DeleteUserWithCredentialByUserID(id uint) []error {
	tx := g.db.Begin()

	var user models.User
	errs := g.db.Where("id = ?", id).First(&user).GetErrors()
	if len(errs) != 0 {
		tx.Rollback()
		return errs
	}

	var credential models.Credential
	errs = tx.Model(&user).Related(&credential).GetErrors()
	if len(errs) != 0 {
		tx.Rollback()
		return errs
	}

	errs = tx.Delete(&user).GetErrors()
	if len(errs) != 0 {
		tx.Rollback()
		return errs
	}

	errs = tx.Delete(&credential).GetErrors()
	if len(errs) != 0 {
		tx.Rollback()
		return errs
	}

	tx.Commit()
	return nil
}
