package handler

import (
	interfaces "masjiddotexe/masjid"

	"net/http"

	"github.com/gin-gonic/gin"
)

func InvRouters(handler interfaces.InvUsecase, g *gin.RouterGroup) {
	uc := InvHandler{
		handler,
	}

	l := g.Group("inventaris")
	l.GET("", uc.GetAllInv)
	l.POST("post", uc.CreateInv)
	l.PUT(":id", uc.UpdateInv)
	l.DELETE(":id", uc.DeleteInv)
}

type InvHandler struct {
	InventarisUc interfaces.InvUsecase
}

func (inv InvHandler) GetAllInv(c *gin.Context) {
	result, err := inv.InventarisUc.GetAllInv(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, result)
}

func (inv InvHandler) CreateInv(c *gin.Context) {
	err := inv.InventarisUc.CreateInv(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": err})
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Post Data"})
}

func (inv InvHandler) UpdateInv(c *gin.Context) {
	err := inv.InventarisUc.UpdateInv(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Update Data"})
}

func (inv InvHandler) DeleteInv(c *gin.Context) {
	err := inv.InventarisUc.DeleteInv(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Delete Data"})
}
