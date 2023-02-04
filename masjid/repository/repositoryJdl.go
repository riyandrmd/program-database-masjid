package repository

import (
	"masjiddotexe/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (repo Repository) GetAll(c *gin.Context) ([]models.JadwalSholat, error) {
	var result []models.JadwalSholat
	data := repo.db.Find(&result)
	if data.Error != nil {
		return nil, data.Error
	}

	return result, nil
}

func (repo Repository) Create(data models.JadwalSholat) error {
	err := repo.db.Create(&data).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo Repository) Update(data models.JadwalSholat, param string) error {
	err := repo.db.First(&models.JadwalSholat{}, "id_jadwal=?", param).Error

	if err != nil {
		return err
	}

	repo.db.Model(&models.JadwalSholat{}).Where("id_jadwal = ?", param).Updates(&data)

	return nil
}

func (repo Repository) Delete(param string) error {
	err := repo.db.First(&models.JadwalSholat{}, "id_jadwal=?", param).Error

	if err != nil {
		return err
	}

	repo.db.Delete(&models.JadwalSholat{}, "id_jadwal=?", param)

	return nil
}
