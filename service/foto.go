package service

import (
	"github.com/MuShaf-NMS/sigmatech-test/entity"
	"github.com/MuShaf-NMS/sigmatech-test/repository"
)

type FotoService interface {
	GetAll(konsumenId uint) ([]entity.Foto, error)
	Create(konsumenId uint, filename, ket string) (entity.Foto, error)
	GetOneByKet(konsumenId uint, ket string) (entity.Foto, error)
}

type fotoService struct {
	repository repository.FotoRepository
}

func (fs *fotoService) GetAll(konsumenId uint) ([]entity.Foto, error) {
	fotos, err := fs.repository.GetAll(konsumenId)
	return fotos, err
}

func (fs *fotoService) Create(konsumenId uint, filename, ket string) (entity.Foto, error) {
	foto := entity.Foto{
		KonsumenId: konsumenId,
		Src:        filename,
		Ket:        ket,
	}
	err := fs.repository.Create(&foto)
	return foto, err
}

func (fs *fotoService) GetOneByKet(konsumenId uint, ket string) (entity.Foto, error) {
	foto, err := fs.repository.GetOneByKet(konsumenId, ket)
	return foto, err
}

func NewFotoService(repository repository.FotoRepository) FotoService {
	return &fotoService{
		repository: repository,
	}
}
