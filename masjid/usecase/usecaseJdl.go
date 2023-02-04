package usecase

import (
	"log"
	interfaces "masjiddotexe/masjid"
	"masjiddotexe/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JadwalUsecase struct {
	jRepo interfaces.JadwalRepository
}

func NewUsecase(jadwal interfaces.JadwalRepository) *JadwalUsecase {
	return &JadwalUsecase{jadwal}
}

func (Jw *JadwalUsecase) GetAll(c *gin.Context) ([]models.JadwalSholat, error) {
	result, err := Jw.jRepo.GetAll(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return nil, err
	}

	return result, nil
}

func (Jw *JadwalUsecase) Create(c *gin.Context) error {
	var data models.JadwalSholat

	err := c.ShouldBindJSON(&data)
	if err != nil {
		return err
	}

	err = Jw.jRepo.Create(data)
	if err != nil {
		return err
	}

	return nil
}

func (Jw *JadwalUsecase) Update(c *gin.Context) error {
	var result models.JadwalSholat

	ID := c.Param("id")
	if ID == "" {
		log.Fatal("ID NOT FOUND")
	}

	err := c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	err = Jw.jRepo.Update(result, ID)
	if err != nil {
		return err
	}

	return nil
}

func (Jw *JadwalUsecase) Delete(c *gin.Context) error {
	ID := c.Param("id")
	if ID == "" {
		log.Fatal("ID NOT FOUND")
	}

	err := Jw.jRepo.Delete(ID)
	if err != nil {
		return err
	}

	return nil
}
