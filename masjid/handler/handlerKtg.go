package handler

import (
	interfaces "masjiddotexe/masjid"

	"net/http"

	"github.com/gin-gonic/gin"
)

func KtgRouters(handler interfaces.KtgUsecase, g *gin.RouterGroup) {
	uc := KtgHandler{
		handler,
	}

	l := g.Group("kategori")
	l.GET("", uc.GetAllKtg)
	l.POST("post", uc.CreateKtg)
	l.PUT(":id", uc.UpdateKtg)
	l.DELETE(":id", uc.DeleteKtg)
}

type KtgHandler struct {
	KtgUc interfaces.KtgUsecase
}

func (inv KtgHandler) GetAllKtg(c *gin.Context) {
	result, err := inv.KtgUc.GetAllKtg(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, result)
}

func (inv KtgHandler) CreateKtg(c *gin.Context) {
	err := inv.KtgUc.CreateKtg(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": err})
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Post Data"})
}

func (inv KtgHandler) UpdateKtg(c *gin.Context) {
	err := inv.KtgUc.UpdateKtg(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Update Data"})
}

func (inv KtgHandler) DeleteKtg(c *gin.Context) {
	err := inv.KtgUc.DeleteKtg(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Delete Data"})
}
