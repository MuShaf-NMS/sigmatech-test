package repository

import (
	"github.com/MuShaf-NMS/sigmatech-test/entity"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Create(auth *entity.Auth) error
	GetOne(konsumenId uint) (entity.Auth, error)
	GetOneByUsername(username string) (entity.Auth, error)
	Update(auth *entity.Auth, konsumenId uint) error
	Delete(konsumenId uint) error
}

type authRepository struct {
	db *gorm.DB
}

func (kr *authRepository) Create(auth *entity.Auth) error {
	err := kr.db.Create(auth).Error
	return err
}

func (kr *authRepository) GetOne(konsumenId uint) (entity.Auth, error) {
	var auth entity.Auth
	err := kr.db.Where(&entity.Auth{KonsumenId: konsumenId}).First(&auth).Error
	return auth, err
}

func (kr *authRepository) GetOneByUsername(username string) (entity.Auth, error) {
	var auth entity.Auth
	err := kr.db.Where(&entity.Auth{Username: username}).First(&auth).Error
	return auth, err
}

func (kr *authRepository) Update(auth *entity.Auth, konsumenId uint) error {
	err := kr.db.Where(&entity.Auth{KonsumenId: konsumenId}).Updates(auth).Error
	return err
}

func (kr *authRepository) Delete(konsumenId uint) error {
	err := kr.db.Where(&entity.Auth{KonsumenId: konsumenId}).Delete(entity.Auth{}).Error
	return err
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}
