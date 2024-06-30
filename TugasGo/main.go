package main

import (
	"project-root/config"
	"project-root/routes"
)

func main() {
	config.Connect() // Koneksi ke database

	// Migrasi tabel Pasien
	// models.MigratePasien(config.DB)

	// Setup router
	r := routes.SetupRouter()

	// Jalankan server
	r.Run(":8080")
}
