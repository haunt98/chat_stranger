package service

import (
	"github.com/1612180/chat_stranger/internal/model"
	"github.com/1612180/chat_stranger/internal/pkg/config"
	"github.com/1612180/chat_stranger/internal/pkg/token"
	"github.com/1612180/chat_stranger/internal/pkg/variable"
	"github.com/1612180/chat_stranger/internal/repository"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

//go:generate $GOPATH/bin/mockgen -destination=../mock/mock_service/mock_account.go -source=account.go

type AccountService interface {
	SignUp(showName, regName, password string) (string, error)
	LogIn(regName, password string) (string, error)
	Info(userID int) (model.User, error)
	UpdateInfo(userID int, showName, gender string, birthYear int) (model.User, error)
}

func NewAccountService(accountRepo repository.AccountRepo, cfg config.Config) AccountService {
	return &defaultAccountService{accountRepo: accountRepo, cfg: cfg}
}

// implement

type defaultAccountService struct {
	accountRepo repository.AccountRepo
	cfg         config.Config
}

func (s *defaultAccountService) SignUp(showName, regName, password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.Wrapf(err, "account service: sign up showName=%s regName=%s failed", regName, showName)
	}

	user, _, err := s.accountRepo.CreateUserCredential(showName, regName, string(hashedPass))
	if err != nil {
		return "", errors.Wrapf(err, "account service: sign up showName=%s regName=%s failed", regName, showName)
	}

	tkn, err := token.Create(token.AccountClaims{ID: user.ID, Role: "user"}, s.cfg.Get(variable.JWTSecret))
	if err != nil {
		return "", errors.Wrapf(err, "account service: sign up showName=%s regName=%s failed", regName, showName)
	}
	return tkn, nil
}

func (s *defaultAccountService) LogIn(regName, password string) (string, error) {
	user, cred, err := s.accountRepo.FindUserCredentialByRegName(regName)
	if err != nil {
		return "", errors.Wrapf(err, "account service: log in regName=%s failed", regName)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(cred.HashedPassword), []byte(password)); err != nil {
		return "", errors.Wrapf(err, "account service: log in regName=%s failed", regName)
	}

	tkn, err := token.Create(token.AccountClaims{ID: user.ID, Role: "user"}, s.cfg.Get(variable.JWTSecret))
	if err != nil {
		return "", errors.Wrapf(err, "account service: log in regName=%s failed", regName)
	}
	return tkn, nil
}

func (s *defaultAccountService) Info(userID int) (model.User, error) {
	user, _, err := s.accountRepo.FindUserCredential(userID)
	if err != nil {
		return user, errors.Wrapf(err, "account service: info user userID=%d", userID)
	}
	return user, nil
}

func (s *defaultAccountService) UpdateInfo(userID int, showName, gender string, birthYear int) (model.User, error) {
	user, err := s.accountRepo.UpdateUser(userID, showName, gender, birthYear)
	if err != nil {
		return user, errors.Wrapf(err, "account service: update info user userID=%d showName=%s gender=%s birthYear=%d",
			userID, showName, gender, birthYear)
	}
	return user, nil
}
