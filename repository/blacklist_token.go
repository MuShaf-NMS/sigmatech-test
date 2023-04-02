package repository

import (
	"github.com/MuShaf-NMS/sigmatech-test/entity"
	"gorm.io/gorm"
)

type BlacklistTokenRepository interface {
	CreateBlacklistToken(blacklistToken *entity.BlacklistToken) *gorm.DB
	// return true if token has been blacklisted false other wise
	CheckBlacklistToken(jti string) bool
}

type blacklistTokenRepository struct {
	connection *gorm.DB
}

func (btr *blacklistTokenRepository) CreateBlacklistToken(blacklistToken *entity.BlacklistToken) *gorm.DB {
	dbRes := btr.connection.Create(blacklistToken)
	return dbRes
}

func (btr *blacklistTokenRepository) CheckBlacklistToken(jti string) bool {
	var blacklists []entity.BlacklistToken
	btr.connection.Where(&entity.BlacklistToken{JTI: jti}).Find(&blacklists)
	return len(blacklists) != 0
}

func NewBlacklistTokenRepository(connection *gorm.DB) BlacklistTokenRepository {
	return &blacklistTokenRepository{
		connection: connection,
	}
}
