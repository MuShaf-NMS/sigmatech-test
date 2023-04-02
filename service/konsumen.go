package service

import (
	"math/rand"
	"time"

	"github.com/MuShaf-NMS/sigmatech-test/custom_types"
	"github.com/MuShaf-NMS/sigmatech-test/dto"
	"github.com/MuShaf-NMS/sigmatech-test/entity"
	"github.com/MuShaf-NMS/sigmatech-test/helper"
	"github.com/MuShaf-NMS/sigmatech-test/repository"
	"golang.org/x/crypto/bcrypt"
)

type KonsumenService interface {
	GetAll() ([]entity.Konsumen, error)
	Create(konsumenDTO dto.KonsumenCreate) error
	GetOne(id uint) (entity.Konsumen, error)
	Update(konsumenDTO dto.KonsumenUpdate, id uint) error
	Delete(id uint) error
}

type konsumenService struct {
	repository              repository.KonsumenRepository
	authRepository          repository.AuthRepository
	limitRrepository        repository.LimitRepository
	limitTerpakaiRepository repository.LimitTerpakaiRepository
}

func (ks *konsumenService) GetAll() ([]entity.Konsumen, error) {
	konsumens, err := ks.repository.GetAll()
	return konsumens, err
}
func (ks *konsumenService) Create(konsumenDTO dto.KonsumenCreate) error {
	tglLahir, err := time.Parse("2006-01-02", konsumenDTO.TanggalLahir)
	if err != nil {
		customErr := helper.NewError(400, "Invalid tgl lahir format")
		return customErr
	}
	konsumen := entity.Konsumen{
		NIK:          konsumenDTO.NIK,
		FullName:     konsumenDTO.FullName,
		LegalName:    konsumenDTO.LegalName,
		TempatLahir:  konsumenDTO.TempatLahir,
		TanggalLahir: custom_types.Date(tglLahir),
		Gaji:         konsumenDTO.Gaji,
	}
	err = ks.repository.Create(&konsumen)
	if err != nil {
		return helper.NewError(400, "Failed to create konsumen")
	}
	password, err := bcrypt.GenerateFromPassword([]byte(konsumenDTO.Password), 10)
	if err != nil {
		return helper.NewError(400, "Failed to create auth")
	}
	auth := entity.Auth{
		KonsumenId: konsumen.Id,
		Username:   konsumenDTO.Username,
		Password:   string(password),
	}
	err = ks.authRepository.Create(&auth)
	if err != nil {
		return helper.NewError(400, "Failed to create konsumen")
	}
	if konsumenDTO.FullName == "Budi" {
		for i, l := range cl[0] {
			l.KonsumenId = konsumen.Id
			err = ks.limitRrepository.Create(l)
			if err != nil {
				return helper.NewError(400, "Failed to create limit")
			}
			if i == 0 {
				limitterpakai := entity.LimitTerpakai{
					KonsumenId: konsumen.Id,
					LimitId:    l.Id,
					Total:      0,
				}
				err = ks.limitTerpakaiRepository.Create(&limitterpakai)
			}
		}
	} else if konsumenDTO.FullName == "Annisa" {
		for i, l := range cl[1] {
			l.KonsumenId = konsumen.Id
			err = ks.limitRrepository.Create(l)
			if err != nil {
				return helper.NewError(400, "Failed to create limit")
			}
			if i == 0 {
				limitterpakai := entity.LimitTerpakai{
					KonsumenId: konsumen.Id,
					LimitId:    l.Id,
					Total:      0,
				}
				err = ks.limitTerpakaiRepository.Create(&limitterpakai)
			}
		}
	} else {
		for i, l := range cl[rand.Intn(2)] {
			l.KonsumenId = konsumen.Id
			err = ks.limitRrepository.Create(l)
			if err != nil {
				return helper.NewError(400, "Failed to create limit")
			}
			if i == 0 {
				limitterpakai := entity.LimitTerpakai{
					KonsumenId: konsumen.Id,
					LimitId:    l.Id,
					Total:      0,
				}
				err = ks.limitTerpakaiRepository.Create(&limitterpakai)
			}
		}
	}
	if err != nil {
		return helper.NewError(400, "Failed to create konsumen")
	}
	return nil
}
func (ks *konsumenService) GetOne(id uint) (entity.Konsumen, error) {
	konsumen, err := ks.repository.GetOne(id)
	return konsumen, err
}
func (ks *konsumenService) Update(konsumenDTO dto.KonsumenUpdate, id uint) error {
	tglLahir, err := time.Parse("2006-01-02", konsumenDTO.TanggalLahir)
	if err != nil {
		customErr := helper.NewError(400, "Invalid tgl lahir format")
		return customErr
	}
	konsumen := entity.Konsumen{
		NIK:          konsumenDTO.NIK,
		FullName:     konsumenDTO.FullName,
		LegalName:    konsumenDTO.LegalName,
		TempatLahir:  konsumenDTO.TempatLahir,
		TanggalLahir: custom_types.Date(tglLahir),
		Gaji:         konsumenDTO.Gaji,
	}
	err = ks.repository.Update(&konsumen, id)
	return err
}
func (ks *konsumenService) Delete(id uint) error {
	err := ks.repository.Delete(id)
	return err
}

func NewKonsumenService(repository repository.KonsumenRepository, authRepository repository.AuthRepository, limitRrepository repository.LimitRepository, limitTerpakaiRepository repository.LimitTerpakaiRepository) KonsumenService {
	return &konsumenService{
		repository:              repository,
		authRepository:          authRepository,
		limitRrepository:        limitRrepository,
		limitTerpakaiRepository: limitTerpakaiRepository,
	}
}
