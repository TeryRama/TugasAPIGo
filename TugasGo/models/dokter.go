package models

import "gorm.io/gorm"

type Dokter struct {
	gorm.Model          // Ini termasuk ID, CreatedAt, UpdatedAt, DeletedAt secara default
	Nama         string `json:"nama"`
	Spesialis    string `json:"spesialis"`
	NomorTelepon string `json:"nomor_telepon"`
}
