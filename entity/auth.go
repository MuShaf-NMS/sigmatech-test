package entity

import "time"

type Auth struct {
	Id         uint      `gorm:"primaryKey" json:"id"`
	KonsumenId uint      `json:"konsumen_id"`
	Username   string    `json:"username"`
	Password   string    `json:"-"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (a *Auth) TableName() string {
	return "auth"
}
