package models

import "time"

type Kunjungan struct {
	ID               uint      `json:"id" gorm:"primary_key"`
	IDPasien         uint      `json:"id_pasien"`
	Pasien           Pasien    `json:"pasien" gorm:"foreignKey:IDPasien"`
	IDDokter         uint      `json:"id_dokter"`
	Dokter           Dokter    `json:"dokter" gorm:"foreignKey:IDDokter"`
	TanggalKunjungan time.Time `json:"tanggal_kunjungan"`
	Keluhan          string    `json:"keluhan"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
