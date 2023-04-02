package service

import (
	"github.com/MuShaf-NMS/sigmatech-test/dto"
	"github.com/MuShaf-NMS/sigmatech-test/entity"
	"github.com/MuShaf-NMS/sigmatech-test/helper"
	"github.com/MuShaf-NMS/sigmatech-test/repository"
)

type TransaksiService interface {
	GetAll(konsumenId uint) ([]entity.Transaksi, error)
	Create(transaksiDTO dto.Transaksi, konsumenId uint) error
	GetOne(konsumenId uint, id uint) (entity.Transaksi, error)
}

type transaksiService struct {
	repository              repository.TransaksiRepository
	limitRepository         repository.LimitRepository
	limitTerpakaiRepository repository.LimitTerpakaiRepository
}

func (ts *transaksiService) GetAll(konsumenId uint) ([]entity.Transaksi, error) {
	transaksis, err := ts.repository.GetAll(konsumenId)
	return transaksis, err
}
func (ts *transaksiService) Create(transaksiDTO dto.Transaksi, konsumenId uint) error {
	limitTerpakai, err := ts.limitTerpakaiRepository.GetOne(konsumenId)
	if err != nil {
		return helper.NewError(400, "Failed to get limit terpakai")
	}
	limit, err := ts.limitRepository.GetAll(konsumenId)
	if err != nil {
		return helper.NewError(400, "Failed to get limit")
	}
	outLimit := true
	total := limitTerpakai.Total + transaksiDTO.JumlahCicilan
	for _, l := range limit {
		if total <= l.Limit {
			outLimit = false
			lt := &limitTerpakai
			lt.Total = total
			lt.LimitId = l.Id
			ts.limitTerpakaiRepository.Update(lt, konsumenId)
			break
		}
	}
	if outLimit {
		return helper.NewError(400, "Limit tercapai")
	}
	transaksi := entity.Transaksi{
		KonsumenId:    konsumenId,
		NoKontrak:     transaksiDTO.NoKontrak,
		OTR:           transaksiDTO.OTR,
		AdminFee:      transaksiDTO.AdminFee,
		JumlahCicilan: transaksiDTO.JumlahCicilan,
		Bunga:         transaksiDTO.Bunga,
		NamaAset:      transaksiDTO.NamaAset,
	}
	err = ts.repository.Create(&transaksi)
	return err
}
func (ts *transaksiService) GetOne(konsumenId uint, id uint) (entity.Transaksi, error) {
	transaksi, err := ts.repository.GetOne(konsumenId, id)
	return transaksi, err
}

func NewTransaksiService(repository repository.TransaksiRepository, limitRepository repository.LimitRepository, limitTerpakaiRepository repository.LimitTerpakaiRepository) TransaksiService {
	return &transaksiService{
		repository:              repository,
		limitRepository:         limitRepository,
		limitTerpakaiRepository: limitTerpakaiRepository,
	}
}
