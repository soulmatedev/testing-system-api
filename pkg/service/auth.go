package service

import (
	"github.com/sirupsen/logrus"
	"testing-system-api/models"
	"testing-system-api/pkg/repository"
	"testing-system-api/pkg/utils"
)

type AuthService struct {
	repo repository.Auth
}

func NewAuthService(repo repository.Auth) *AuthService {
	return &AuthService{repo: repo}
}

func (a AuthService) SignIn(input *models.SignInInput, accountPassword string) error {
	if err := utils.ComparePasswords(accountPassword, input.Password); err != nil {
		logrus.Warning(err.Error())
		return err
	}
	return nil
}
