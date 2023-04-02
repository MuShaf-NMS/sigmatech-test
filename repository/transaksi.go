package repository

import (
	"github.com/MuShaf-NMS/sigmatech-test/entity"
	"gorm.io/gorm"
)

type TransaksiRepository interface {
	GetAll(konsumenId uint) ([]entity.Transaksi, error)
	Create(transaksi *entity.Transaksi) error
	GetOne(konsumenId uint, id uint) (entity.Transaksi, error)
}

type transaksiRepository struct {
	db *gorm.DB
}

func (tr *transaksiRepository) GetAll(konsumenId uint) ([]entity.Transaksi, error) {
	var transaksis []entity.Transaksi
	err := tr.db.Where(&entity.Transaksi{KonsumenId: konsumenId}).Find(&transaksis).Error
	return transaksis, err
}

func (tr *transaksiRepository) Create(transaksi *entity.Transaksi) error {
	err := tr.db.Create(transaksi).Error
	return err
}

func (tr *transaksiRepository) GetOne(konsumenId uint, id uint) (entity.Transaksi, error) {
	var transaksi entity.Transaksi
	err := tr.db.Where(&entity.Transaksi{Id: id, KonsumenId: konsumenId}).First(&transaksi).Error
	return transaksi, err
}

func NewTransaksiRepository(db *gorm.DB) TransaksiRepository {
	return &transaksiRepository{
		db: db,
	}
}
