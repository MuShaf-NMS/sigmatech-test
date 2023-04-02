package dto

type KonsumenCreate struct {
	Username     string `json:"username" validate:"required"`
	Password     string `json:"password" validate:"required"`
	NIK          string `json:"nik" validate:"required"`
	FullName     string `json:"full_name" validate:"required"`
	LegalName    string `json:"legal_name" validate:"required"`
	TempatLahir  string `json:"tempat_lahir" validate:"required"`
	TanggalLahir string `json:"tanggal_lahir" validate:"required"`
	Gaji         uint   `json:"gaji" validate:"required"`
}

type KonsumenUpdate struct {
	NIK          string `json:"nik" validate:"required"`
	FullName     string `json:"full_name" validate:"required"`
	LegalName    string `json:"legal_name" validate:"required"`
	TempatLahir  string `json:"tempat_lahir" validate:"required"`
	TanggalLahir string `json:"tanggal_lahir" validate:"required"`
	Gaji         uint   `json:"gaji" validate:"required"`
}
