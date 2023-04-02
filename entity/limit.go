package entity

import "time"

type LimitPinjaman struct {
	Id         uint      `gorm:"primaryKey" json:"id"`
	KonsumenId uint      `json:"konsumen_id"`
	Tenor      uint      `json:"tenor"`
	Limit      uint      `json:"limit"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (l *LimitPinjaman) TableName() string {
	return "limit"
}
