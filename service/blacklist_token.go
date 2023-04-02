package service

import (
	"github.com/MuShaf-NMS/sigmatech-test/entity"
	"github.com/MuShaf-NMS/sigmatech-test/repository"
)

type BlacklistTokenService interface {
	CreateBlacklistToken(jti string)
	// return true if token has been blacklisted false other wise
	CheckBlacklistToken(jti string) bool
}

type blackListTokenService struct {
	repository repository.BlacklistTokenRepository
}

func (as *blackListTokenService) CreateBlacklistToken(jti string) {
	blackListTokenCreate := entity.BlacklistToken{
		JTI: jti,
	}
	as.repository.CreateBlacklistToken(&blackListTokenCreate)
}

func (as *blackListTokenService) CheckBlacklistToken(jti string) bool {
	res := as.repository.CheckBlacklistToken(jti)
	return res
}

func NewBlacklistTokenService(repository repository.BlacklistTokenRepository) BlacklistTokenService {
	return &blackListTokenService{
		repository: repository,
	}
}
