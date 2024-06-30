package models

import (
	"time"

	"gorm.io/gorm"
)

type Kunjungan struct {
	gorm.Model                 // Ini termasuk ID, CreatedAt, UpdatedAt, DeletedAt secara default
	IDPasien         uint      `json:"id_pasien"`
	Pasien           Pasien    `json:"pasien" gorm:"foreignKey:IDPasien"`
	IDDokter         uint      `json:"id_dokter"`
	Dokter           Dokter    `json:"dokter" gorm:"foreignKey:IDDokter"`
	TanggalKunjungan time.Time `json:"tanggal_kunjungan"`
	Keluhan          string    `json:"keluhan"`
}
