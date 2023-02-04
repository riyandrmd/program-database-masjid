package usecase

import (
	"log"
	interfaces "masjiddotexe/masjid"
	"masjiddotexe/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InvUsecase struct {
	jRepo interfaces.InvRepository
}

func NewUsecaseInv(inv interfaces.InvRepository) *InvUsecase {
	return &InvUsecase{inv}
}

func (Jw *InvUsecase) GetAllInv(c *gin.Context) ([]models.Inventaris, error) {
	result, err := Jw.jRepo.GetAllInv(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return nil, err
	}

	return result, nil
}

func (Jw *InvUsecase) CreateInv(c *gin.Context) error {
	var data models.Inventaris

	err := c.ShouldBindJSON(&data)
	if err != nil {
		return err
	}

	err = Jw.jRepo.CreateInv(data)
	if err != nil {
		return err
	}

	return nil
}

func (Jw *InvUsecase) UpdateInv(c *gin.Context) error {
	var result models.Inventaris

	ID := c.Param("id")
	if ID == "" {
		log.Fatal("ID NOT FOUND")
	}

	err := c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	err = Jw.jRepo.UpdateInv(result, ID)
	if err != nil {
		return err
	}

	return nil
}

func (Jw *InvUsecase) DeleteInv(c *gin.Context) error {
	ID := c.Param("id")
	if ID == "" {
		log.Fatal("ID NOT FOUND")
	}

	err := Jw.jRepo.DeleteInv(ID)
	if err != nil {
		return err
	}

	return nil
}
