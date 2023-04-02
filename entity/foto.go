package entity

import (
	"time"

	"github.com/MuShaf-NMS/sigmatech-test/config"
	"gorm.io/gorm"
)

// var conf = config.GetConfig()

type Foto struct {
	Id         uint      `gorm:"primaryKey" json:"id"`
	KonsumenId uint      `json:"konsumen_id"`
	Src        string    `gorm:"size(255)" json:"-"`
	Url        string    `gorm:"-" json:"url"`
	Ket        string    `json:"ket"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (foto *Foto) AfterFind(tx *gorm.DB) error {
	foto.Url = config.Conf.Assets_Link + "/" + foto.Src
	return nil
}

func (foto *Foto) AfterCreate(tx *gorm.DB) error {
	foto.Url = config.Conf.Assets_Link + "/" + foto.Src
	return nil
}

func (f *Foto) TableName() string {
	return "foto"
}
