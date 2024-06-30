package models

type Dokter struct {
	ID           uint   `json:"id_dokter" gorm:"primaryKey"`
	Nama         string `json:"nama"`
	Spesialis    string `json:"spesialis"`
	NomorTelepon string `json:"nomor_telepon"`
}
