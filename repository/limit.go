package repository

import (
	"github.com/MuShaf-NMS/sigmatech-test/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type LimitRepository interface {
	GetAll(konsumenId uint) ([]entity.LimitPinjaman, error)
	Create(limit *entity.LimitPinjaman) error
	Delete(konsumenId uint) error
}

type limitRepository struct {
	db *gorm.DB
}

func (kr *limitRepository) GetAll(konsumenId uint) ([]entity.LimitPinjaman, error) {
	var konsumens []entity.LimitPinjaman
	err := kr.db.Order(clause.OrderByColumn{Column: clause.Column{Name: "tenor"}}).Where(&entity.LimitPinjaman{KonsumenId: konsumenId}).Find(&konsumens).Error
	return konsumens, err
}

func (kr *limitRepository) Create(limit *entity.LimitPinjaman) error {
	err := kr.db.Create(limit).Error
	return err
}

func (kr *limitRepository) Delete(konsumenId uint) error {
	err := kr.db.Where(&entity.LimitPinjaman{KonsumenId: konsumenId}).Delete(entity.LimitPinjaman{}).Error
	return err
}

func NewLimitRepository(db *gorm.DB) LimitRepository {
	return &limitRepository{
		db: db,
	}
}
