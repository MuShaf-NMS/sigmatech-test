package entity

import "time"

type LimitTerpakai struct {
	Id         uint          `gorm:"primaryKey" json:"id"`
	KonsumenId uint          `json:"konsumen_id"`
	LimitId    uint          `json:"limit_id"`
	Total      uint          `json:"total"`
	CreatedAt  time.Time     `json:"created_at"`
	UpdatedAt  time.Time     `json:"updated_at"`
	Limit      LimitPinjaman `gorm:"foreignKey:LimitId;references:Id"`
}

func (l *LimitTerpakai) TableName() string {
	return "limit_terpakai"
}
