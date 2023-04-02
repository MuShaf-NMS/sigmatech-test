package entity

import "time"

type BlacklistToken struct {
	Id        uint      `gorm:"primaryKey;autoIncrement"`
	JTI       string    `gorm:"type:char(36);not null" json:"jti"`
	CreatedAt time.Time `json:"created_at"`
}

func (b *BlacklistToken) TableName() string {
	return "blacklist_token"
}
