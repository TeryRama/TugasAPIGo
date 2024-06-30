package models

import "gorm.io/gorm"

type Pasien struct {
	gorm.Model          // Ini termasuk ID, CreatedAt, UpdatedAt, DeletedAt secara default
	Nama         string `json:"nama"`
	Alamat       string `json:"alamat"`
	TanggalLahir string `json:"tanggal_lahir"`
	NomorTelepon string `json:"nomor_telepon"`
}
