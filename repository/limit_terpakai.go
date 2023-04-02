package repository

import (
	"github.com/MuShaf-NMS/sigmatech-test/entity"
	"gorm.io/gorm"
)

type LimitTerpakaiRepository interface {
	Create(limitTerpakai *entity.LimitTerpakai) error
	GetOne(konsumenId uint) (entity.LimitTerpakai, error)
	Update(limitTerpakai *entity.LimitTerpakai, konsumenId uint) error
	Delete(konsumenId uint) error
}

type limitTerpakaiRepository struct {
	db *gorm.DB
}

func (kr *limitTerpakaiRepository) Create(limitTerpakai *entity.LimitTerpakai) error {
	err := kr.db.Create(limitTerpakai).Error
	return err
}

func (kr *limitTerpakaiRepository) GetOne(konsumenId uint) (entity.LimitTerpakai, error) {
	var limitTerpakai entity.LimitTerpakai
	err := kr.db.Where(&entity.LimitTerpakai{KonsumenId: konsumenId}).First(&limitTerpakai).Error
	return limitTerpakai, err
}

func (kr *limitTerpakaiRepository) Update(limitTerpakai *entity.LimitTerpakai, konsumenId uint) error {
	err := kr.db.Where(&entity.LimitTerpakai{KonsumenId: konsumenId}).Updates(limitTerpakai).Error
	return err
}

func (kr *limitTerpakaiRepository) Delete(konsumenId uint) error {
	err := kr.db.Where(&entity.LimitTerpakai{KonsumenId: konsumenId}).Delete(entity.LimitTerpakai{}).Error
	return err
}

func NewLimitTerpakaiRepository(db *gorm.DB) LimitTerpakaiRepository {
	return &limitTerpakaiRepository{
		db: db,
	}
}
