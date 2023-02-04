package usecase

import (
	"log"
	interfaces "masjiddotexe/masjid"
	"masjiddotexe/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type KmUsecase struct {
	jRepo interfaces.InfaqMasukRepo
}

func NewUsecaseKm(KasMasuk interfaces.InfaqMasukRepo) *KmUsecase {
	return &KmUsecase{KasMasuk}
}

func (Jw *KmUsecase) GetAllKM(c *gin.Context) ([]models.InfaqMasuk, error) {
	result, err := Jw.jRepo.GetAllKM(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return nil, err
	}

	return result, nil
}

func (Jw *KmUsecase) CreateKM(c *gin.Context) error {
	var data models.InfaqMasuk

	err := c.ShouldBindJSON(&data)
	if err != nil {
		return err
	}

	err = Jw.jRepo.CreateKM(data)
	if err != nil {
		return err
	}

	return nil
}

func (Jw *KmUsecase) UpdateKM(c *gin.Context) error {
	var result models.InfaqMasuk

	ID := c.Param("id")
	if ID == "" {
		log.Fatal("ID NOT FOUND")
	}

	err := c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	err = Jw.jRepo.UpdateKM(result, ID)
	if err != nil {
		return err
	}

	return nil
}

func (Jw *KmUsecase) DeleteKM(c *gin.Context) error {
	ID := c.Param("id")
	if ID == "" {
		log.Fatal("ID NOT FOUND")
	}

	err := Jw.jRepo.DeleteKM(ID)
	if err != nil {
		return err
	}

	return nil
}
