package handler

import (
	interfaces "masjiddotexe/masjid"

	"net/http"

	"github.com/gin-gonic/gin"
)

func JwRouters(handler interfaces.JadwalUsecase, g *gin.RouterGroup) {
	uc := JdHandler{
		handler,
	}

	l := g.Group("jadwal")
	l.GET("", uc.GetAllJadwal)
	l.POST("post", uc.CreateJadwal)
	l.PUT(":id", uc.UpdateJadwal)
	l.DELETE(":id", uc.DeleteJadwal)
}

type JdHandler struct {
	JadwalUc interfaces.JadwalUsecase
}

func (jh JdHandler) GetAllJadwal(c *gin.Context) {
	result, err := jh.JadwalUc.GetAll(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, result)
}

func (jh JdHandler) CreateJadwal(c *gin.Context) {
	err := jh.JadwalUc.Create(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": err})
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Post Data"})
}

func (jh JdHandler) UpdateJadwal(c *gin.Context) {
	err := jh.JadwalUc.Update(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Update Data"})
}

func (jh JdHandler) DeleteJadwal(c *gin.Context) {
	err := jh.JadwalUc.Delete(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Delete Data"})
}
