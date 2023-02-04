package intrface

import (
	"masjiddotexe/models"

	"github.com/gin-gonic/gin"
)

type JadwalRepository interface {
	GetAll(c *gin.Context) ([]models.JadwalSholat, error)
	Create(data models.JadwalSholat) error
	Update(data models.JadwalSholat, param string) error
	Delete(param string) error
}

type JadwalUsecase interface {
	GetAll(c *gin.Context) ([]models.JadwalSholat, error)
	Create(c *gin.Context) error
	Update(c *gin.Context) error
	Delete(c *gin.Context) error
}

//////////////////////////////////////////////////////////////////////////////////////////

type InvRepository interface {
	GetAllInv(c *gin.Context) ([]models.Inventaris, error)
	CreateInv(data models.Inventaris) error
	UpdateInv(data models.Inventaris, param string) error
	DeleteInv(param string) error
}

type InvUsecase interface {
	GetAllInv(c *gin.Context) ([]models.Inventaris, error)
	CreateInv(c *gin.Context) error
	UpdateInv(c *gin.Context) error
	DeleteInv(c *gin.Context) error
}

//////////////////////////////////////////////////////////////////////////////////////////

type KtgRepository interface {
	GetAllKtg(c *gin.Context) ([]models.Ktg_Inventaris, error)
	CreateKtg(data models.Ktg_Inventaris) error
	UpdateKtg(data models.Ktg_Inventaris, param string) error
	DeleteKtg(param string) error
}

type KtgUsecase interface {
	GetAllKtg(c *gin.Context) ([]models.Ktg_Inventaris, error)
	CreateKtg(c *gin.Context) error
	UpdateKtg(c *gin.Context) error
	DeleteKtg(c *gin.Context) error
}
