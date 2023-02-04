package repository

import (
	"masjiddotexe/models"

	"github.com/gin-gonic/gin"
)

func (repo Repository) GetAllKM(c *gin.Context) ([]models.InfaqMasuk, error) {
	var result []models.InfaqMasuk
	data := repo.db.Find(&result)
	if data.Error != nil {
		return nil, data.Error
	}

	return result, nil
}

func (repo Repository) CreateKM(data models.InfaqMasuk) error {
	err := repo.db.Create(&data).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo Repository) UpdateKM(data models.InfaqMasuk, param string) error {
	err := repo.db.First(&models.InfaqMasuk{}, "id_barang=?", param).Error

	if err != nil {
		return err
	}

	repo.db.Model(&models.InfaqMasuk{}).Where("id_barang = ?", param).Updates(&data)

	return nil
}

func (repo Repository) DeleteKM(param string) error {
	err := repo.db.First(&models.InfaqMasuk{}, "id_barang=?", param).Error

	if err != nil {
		return err
	}

	repo.db.Delete(&models.InfaqMasuk{}, "id_barang=?", param)

	return nil
}
