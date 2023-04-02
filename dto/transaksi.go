package dto

type Transaksi struct {
	NoKontrak     string `json:"no_kontrak" validate:"required"`
	OTR           uint   `json:"otr" validate:"required"`
	AdminFee      uint   `json:"admin_fee" validate:"required"`
	JumlahCicilan uint   `json:"jumlah_cicilan" validate:"required"`
	Bunga         uint   `json:"bunga" validate:"required"`
	NamaAset      string `json:"nama_aset" validate:"required"`
}
