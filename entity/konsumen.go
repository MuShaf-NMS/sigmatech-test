package entity

import (
	"time"

	"github.com/MuShaf-NMS/sigmatech-test/custom_types"
)

type Konsumen struct {
	Id            uint              `gorm:"primaryKey" json:"id"`
	NIK           string            `gorm:"type:char(18)" json:"nik"`
	FullName      string            `gorm:"type:varchar(200)" json:"full_name"`
	LegalName     string            `gorm:"type:varchar(200)" json:"legal_name"`
	TempatLahir   string            `gorm:"type:varchar(200)" json:"tempat_lahir"`
	TanggalLahir  custom_types.Date `gorm:"type:date" json:"tanggal_lahir"`
	Gaji          uint              `gorm:"primaryKey" json:"gaji"`
	FotoKTP       uint              `gorm:"primaryKey" json:"foto_ktp"`
	FotoSelfie    uint              `gorm:"primaryKey" json:"foto_selfie"`
	CreatedAt     time.Time         `json:"created_at"`
	UpdatedAt     time.Time         `json:"updated_at"`
	Foto          []Foto            `gorm:"foreignKey:KonsumenId;references:Id" json:"foto"`
	Auth          Auth              `gorm:"foreignKey:KonsumenId;references:Id" json:"auth"`
	Limit         []LimitPinjaman   `gorm:"foreignKey:KonsumenId;references:Id" json:"limit"`
	LimitTerpakai LimitTerpakai     `gorm:"foreignKey:KonsumenId;references:Id" json:"limit_terpakai"`
	Transaksi     []Transaksi       `gorm:"foreignKey:KonsumenId;references:Id" json:"transaksi"`
}

func (k *Konsumen) TableName() string {
	return "konsumen"
}
