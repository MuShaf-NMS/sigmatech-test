package service

import (
	"github.com/MuShaf-NMS/sigmatech-test/dto"
	"github.com/MuShaf-NMS/sigmatech-test/entity"
	"github.com/MuShaf-NMS/sigmatech-test/helper"
	"github.com/MuShaf-NMS/sigmatech-test/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	GetOne(konsumenId uint) (entity.Auth, error)
	Update(authDTO dto.UpdateAuth, konsumenId uint) error
	Login(login dto.Auth) (entity.Auth, error)
}

type authService struct {
	repository repository.AuthRepository
}

func (ks *authService) GetOne(konsumenId uint) (entity.Auth, error) {
	auth, err := ks.repository.GetOne(konsumenId)
	if err != nil {
		return auth, helper.NewError(404, "Not found")
	}
	return auth, err
}
func (ks *authService) Update(authDTO dto.UpdateAuth, konsumenId uint) error {
	auth, err := ks.GetOne(konsumenId)
	if err != nil {
		return err
	}
	if bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(authDTO.OldPassword)) != nil {
		return helper.NewError(400, "Old Password wrong")
	}
	password, err := bcrypt.GenerateFromPassword([]byte(authDTO.NewPassword), 10)
	if err != nil {
		return helper.NewError(400, "Failed to generate password")
	}
	auth = entity.Auth{
		Username: authDTO.Username,
		Password: string(password),
	}
	err = ks.repository.Update(&auth, konsumenId)
	return err
}

func (ks *authService) Login(login dto.Auth) (entity.Auth, error) {
	auth, err := ks.repository.GetOneByUsername(login.Username)
	if err != nil {
		return auth, helper.NewError(400, "Username or Passowrd wrong")
	}
	err = bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(login.Password))
	if err != nil {
		return auth, helper.NewError(400, "Username or Passowrd wrong")
	}
	return auth, nil
}

func NewAuthService(repository repository.AuthRepository) AuthService {
	return &authService{
		repository: repository,
	}
}
