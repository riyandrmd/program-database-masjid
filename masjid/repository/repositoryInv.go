package repository

import (
	"masjiddotexe/models"

	"github.com/gin-gonic/gin"
)

func (repo Repository) GetAllInv(c *gin.Context) ([]models.Inventaris, error) {
	var result []models.Inventaris
	data := repo.db.Find(&result)
	if data.Error != nil {
		return nil, data.Error
	}

	return result, nil
}

func (repo Repository) CreateInv(data models.Inventaris) error {
	err := repo.db.Create(&data).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo Repository) UpdateInv(data models.Inventaris, param string) error {
	err := repo.db.First(&models.Inventaris{}, "id_barang=?", param).Error

	if err != nil {
		return err
	}

	repo.db.Model(&models.Inventaris{}).Where("id_barang = ?", param).Updates(&data)

	return nil
}

func (repo Repository) DeleteInv(param string) error {
	err := repo.db.First(&models.Inventaris{}, "id_barang=?", param).Error

	if err != nil {
		return err
	}

	repo.db.Delete(&models.Inventaris{}, "id_barang=?", param)

	return nil
}
