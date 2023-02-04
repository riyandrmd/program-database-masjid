package handler

import (
	interfaces "masjiddotexe/masjid"

	"net/http"

	"github.com/gin-gonic/gin"
)

func KmRouters(handler interfaces.InfaqMasukUsecase, g *gin.RouterGroup) {
	uc := KmHandler{
		handler,
	}

	l := g.Group("infaqmasuk")
	l.GET("", uc.GetAllKM)
	l.POST("post", uc.CreateKM)
	l.PUT(":id", uc.UpdateKM)
	l.DELETE(":id", uc.DeleteKM)
}

type KmHandler struct {
	KmUc interfaces.InfaqMasukUsecase
}

func (km KmHandler) GetAllKM(c *gin.Context) {
	result, err := km.KmUc.GetAllKM(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, result)
}

func (km KmHandler) CreateKM(c *gin.Context) {
	err := km.KmUc.CreateKM(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": err})
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Post Data"})
}

func (km KmHandler) UpdateKM(c *gin.Context) {
	err := km.KmUc.UpdateKM(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Update Data"})
}

func (km KmHandler) DeleteKM(c *gin.Context) {
	err := km.KmUc.DeleteKM(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Delete Data"})
}
