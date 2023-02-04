package usecase

import (
	"log"
	interfaces "masjiddotexe/masjid"
	"masjiddotexe/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type KtgUsecase struct {
	jRepo interfaces.KtgRepository
}

func NewUsecaseKtg(inv interfaces.KtgRepository) *KtgUsecase {
	return &KtgUsecase{inv}
}

func (Ktg *KtgUsecase) GetAllKtg(c *gin.Context) ([]models.Ktg_Inventaris, error) {
	result, err := Ktg.jRepo.GetAllKtg(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return nil, err
	}

	return result, nil
}

func (Ktg *KtgUsecase) CreateKtg(c *gin.Context) error {
	var data models.Ktg_Inventaris

	err := c.ShouldBindJSON(&data)
	if err != nil {
		return err
	}

	err = Ktg.jRepo.CreateKtg(data)
	if err != nil {
		return err
	}

	return nil
}

func (Ktg *KtgUsecase) UpdateKtg(c *gin.Context) error {
	var result models.Ktg_Inventaris

	ID := c.Param("id")
	if ID == "" {
		log.Fatal("ID NOT FOUND")
	}

	err := c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	err = Ktg.jRepo.UpdateKtg(result, ID)
	if err != nil {
		return err
	}

	return nil
}

func (Ktg *KtgUsecase) DeleteKtg(c *gin.Context) error {
	ID := c.Param("id")
	if ID == "" {
		log.Fatal("ID NOT FOUND")
	}

	err := Ktg.jRepo.DeleteKtg(ID)
	if err != nil {
		return err
	}

	return nil
}
