package repository

import (
	"masjiddotexe/models"

	"github.com/gin-gonic/gin"
)

func (repo Repository) GetAllKtg(c *gin.Context) ([]models.Ktg_Inventaris, error) {
	var result []models.Ktg_Inventaris

	data := repo.db.Model(&models.Ktg_Inventaris{}).Preload("Inventaris").Find(&result)

	if data.Error != nil {
		return nil, data.Error
	}

	return result, nil
}

func (repo Repository) CreateKtg(data models.Ktg_Inventaris) error {
	err := repo.db.Create(&data).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo Repository) UpdateKtg(data models.Ktg_Inventaris, param string) error {
	err := repo.db.First(&models.Ktg_Inventaris{}, "id_ktg=?", param).Error

	if err != nil {
		return err
	}

	repo.db.Model(&models.Ktg_Inventaris{}).Where("id_ktg = ?", param).Updates(&data)

	return nil
}

func (repo Repository) DeleteKtg(param string) error {
	err := repo.db.First(&models.Ktg_Inventaris{}, "id_ktg=?", param).Error

	if err != nil {
		return err
	}

	repo.db.Delete(&models.Ktg_Inventaris{}, "id_ktg=?", param)

	return nil
}
