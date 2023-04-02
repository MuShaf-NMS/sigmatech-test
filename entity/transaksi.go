package entity

import "time"

type Transaksi struct {
	Id            uint      `gorm:"primaryKey" json:"id"`
	KonsumenId    uint      `json:"konsumen_id"`
	NoKontrak     string    `json:"no_kontrak"`
	OTR           uint      `json:"otr"`
	AdminFee      uint      `json:"admin_fee"`
	JumlahCicilan uint      `json:"jumlah_cicilan"`
	Bunga         uint      `json:"bungan"`
	NamaAset      string    `json:"nama_aset"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (t *Transaksi) TableName() string {
	return "transaksi"
}
