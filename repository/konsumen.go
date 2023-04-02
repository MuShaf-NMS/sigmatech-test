package repository

import (
	"github.com/MuShaf-NMS/sigmatech-test/entity"
	"gorm.io/gorm"
)

type KonsumenRepository interface {
	GetAll() ([]entity.Konsumen, error)
	Create(konsumen *entity.Konsumen) error
	GetOne(id uint) (entity.Konsumen, error)
	Update(konsumen *entity.Konsumen, id uint) error
	Delete(id uint) error
}

type konsumenRepository struct {
	db *gorm.DB
}

func (kr *konsumenRepository) GetAll() ([]entity.Konsumen, error) {
	var konsumens []entity.Konsumen
	err := kr.db.Find(&konsumens).Error
	return konsumens, err
}

func (kr *konsumenRepository) Create(konsumen *entity.Konsumen) error {
	err := kr.db.Create(konsumen).Error
	return err
}

func (kr *konsumenRepository) GetOne(id uint) (entity.Konsumen, error) {
	var konsumen entity.Konsumen
	err := kr.db.Where(&entity.Konsumen{Id: id}).First(&konsumen).Error
	return konsumen, err
}

func (kr *konsumenRepository) Update(konsumen *entity.Konsumen, id uint) error {
	err := kr.db.Where(&entity.Konsumen{Id: id}).Updates(konsumen).Error
	return err
}

func (kr *konsumenRepository) Delete(id uint) error {
	err := kr.db.Where(&entity.Konsumen{Id: id}).Delete(entity.Konsumen{}).Error
	return err
}

func NewKonsumenRepository(db *gorm.DB) KonsumenRepository {
	return &konsumenRepository{
		db: db,
	}
}
