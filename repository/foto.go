package repository

import (
	"github.com/MuShaf-NMS/sigmatech-test/entity"
	"gorm.io/gorm"
)

type FotoRepository interface {
	GetAll(konsumenId uint) ([]entity.Foto, error)
	Create(foto *entity.Foto) error
	GetOneByKet(konsumenId uint, ket string) (entity.Foto, error)
}

type fotoRepository struct {
	db *gorm.DB
}

func (fr *fotoRepository) GetAll(konsumenId uint) ([]entity.Foto, error) {
	var transaksis []entity.Foto
	err := fr.db.Where(&entity.Foto{KonsumenId: konsumenId}).Find(&transaksis).Error
	return transaksis, err
}

func (fr *fotoRepository) Create(foto *entity.Foto) error {
	err := fr.db.Create(foto).Error
	return err
}

func (fr *fotoRepository) GetOneByKet(konsumenId uint, ket string) (entity.Foto, error) {
	var foto entity.Foto
	err := fr.db.Where(&entity.Foto{Ket: ket, KonsumenId: konsumenId}).First(&foto).Error
	return foto, err
}

func NewFotoRepository(db *gorm.DB) FotoRepository {
	return &fotoRepository{
		db: db,
	}
}
