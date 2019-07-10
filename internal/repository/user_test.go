package repository

import (
	"testing"

	"github.com/1612180/chat_stranger/internal/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
)

func TestUserGorm_Find(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	userGorm := &userGorm{db: db}

	t.Run("false", func(t *testing.T) {
		migrate(db, t)

		_, _, ok := userGorm.Find(1)
		assert.Equal(t, false, ok)
	})

	t.Run("true", func(t *testing.T) {
		migrate(db, t)

		if err := db.Create(&model.Credential{}).Error; err != nil {
			t.Error(err)
		}
		if err := db.Create(&model.User{CredentialID: 1}).Error; err != nil {
			t.Error(err)
		}

		user, credential, ok := userGorm.Find(1)
		assert.Equal(t, true, ok)
		assert.Equal(t, 1, user.ID)
		assert.Equal(t, 1, user.CredentialID)
		assert.Equal(t, 1, credential.ID)
	})

}

func TestUserGorm_FindByRegisterName(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	userGorm := &userGorm{db: db}

	t.Run("false", func(t *testing.T) {
		migrate(db, t)

		_, _, ok := userGorm.FindByRegisterName("a")
		assert.Equal(t, false, ok)
	})

	t.Run("true", func(t *testing.T) {
		migrate(db, t)

		if err := db.Create(&model.Credential{RegisterName: "a"}).Error; err != nil {
			t.Error(err)
		}
		if err := db.Create(&model.User{CredentialID: 1}).Error; err != nil {
			t.Error(err)
		}

		user, credential, ok := userGorm.FindByRegisterName("a")
		assert.Equal(t, true, ok)
		assert.Equal(t, 1, user.ID)
		assert.Equal(t, 1, user.CredentialID)
		assert.Equal(t, 1, credential.ID)
		assert.Equal(t, "a", credential.RegisterName)
	})
}

func TestUserGorm_Create(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	migrate(db, t)

	userGorm := &userGorm{db: db}

	var user model.User
	var credential model.Credential
	ok := userGorm.Create(&user, &credential)
	assert.Equal(t, true, ok)
	assert.Equal(t, 1, user.ID)
	assert.Equal(t, user.CredentialID, credential.ID)
	assert.Equal(t, 1, credential.ID)
}

func TestUserGorm_Delete(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	userGorm := &userGorm{db: db}

	t.Run("false", func(t *testing.T) {
		migrate(db, t)

		ok := userGorm.Delete(1)
		assert.Equal(t, false, ok)
	})

	t.Run("true", func(t *testing.T) {
		migrate(db, t)

		if err := db.Create(&model.Credential{}).Error; err != nil {
			t.Error(err)
		}
		if err := db.Create(&model.User{CredentialID: 1}).Error; err != nil {
			t.Error(err)
		}

		ok := userGorm.Delete(1)
		assert.Equal(t, true, ok)
	})
}

func TestUserGorm_UpdateInfo(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	userGorm := &userGorm{db: db}

	t.Run("false", func(t *testing.T) {
		migrate(db, t)

		ok := userGorm.UpdateInfo(1, &model.User{})
		assert.Equal(t, false, ok)
	})

	t.Run("true", func(t *testing.T) {
		migrate(db, t)

		if err := db.Create(&model.User{}).Error; err != nil {
			t.Error(err)
		}

		ok := userGorm.UpdateInfo(1, &model.User{
			FullName:  "a",
			Gender:    "b",
			BirthYear: 1,
		})
		assert.Equal(t, true, ok)

		var user model.User
		if err := db.Find(&model.User{ID: 1}).First(&user).Error; err != nil {
			t.Error(err)
		}

		assert.Equal(t, "a", user.FullName)
		assert.Equal(t, "b", user.Gender)
		assert.Equal(t, 1, user.BirthYear)
	})
}

func TestUserGorm_UpdatePassword(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	userGorm := &userGorm{db: db}

	t.Run("false", func(t *testing.T) {
		migrate(db, t)

		ok := userGorm.UpdatePassword(1, &model.Credential{})
		assert.Equal(t, false, ok)
	})

	t.Run("true", func(t *testing.T) {
		migrate(db, t)

		if err := db.Create(&model.Credential{}).Error; err != nil {
			t.Error(err)
		}
		if err := db.Create(&model.User{CredentialID: 1}).Error; err != nil {
			t.Error(err)
		}

		ok := userGorm.UpdatePassword(1, &model.Credential{HashedPassword: "a"})
		assert.Equal(t, true, ok)

		var credential model.Credential
		if err := db.Where(&model.Credential{ID: 1}).First(&credential).Error; err != nil {
			t.Error(err)
		}

		assert.Equal(t, "a", credential.HashedPassword)
	})
}
