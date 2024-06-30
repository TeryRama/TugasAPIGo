package routes

import (
	"project-root/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Routes for Pasien
	router.GET("/pasien", controllers.GetPasien)
	router.GET("/pasien/:id", controllers.GetPasienByID)
	router.POST("/pasien", controllers.CreatePasien)
	router.PUT("/pasien/:id", controllers.UpdatePasien)
	router.DELETE("/pasien/:id", controllers.DeletePasien)

	// Routes for Dokter
	router.GET("/dokter", controllers.GetDokter)
	router.GET("/dokter/:id", controllers.GetDokterByID)
	router.POST("/dokter", controllers.CreateDokter)
	router.PUT("/dokter/:id", controllers.UpdateDokter)
	router.DELETE("/dokter/:id", controllers.DeleteDokter)

	// Routes for Kunjungan
	router.GET("/kunjungan", controllers.GetKunjungan)
	router.GET("/kunjungan/:id", controllers.GetKunjunganByID)
	router.POST("/kunjungan", controllers.CreateKunjungan)
	router.PUT("/kunjungan/:id", controllers.UpdateKunjungan)
	router.DELETE("/kunjungan/:id", controllers.DeleteKunjungan)

	return router
}
