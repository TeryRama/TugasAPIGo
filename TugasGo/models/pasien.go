package models

type Pasien struct {
	ID           int    `json:"id" gorm:"primaryKey"`
	Nama         string `json:"nama"`
	Alamat       string `json:"alamat"`
	TanggalLahir string `json:"tanggal_lahir"`
	NomorTelepon string `json:"nomor_telepon"`
}
