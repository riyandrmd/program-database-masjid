package routers

import (
	"masjiddotexe/masjid/handler"
	"masjiddotexe/masjid/repository"
	"masjiddotexe/masjid/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HandlerStr struct {
	DB *gorm.DB
	R  *gin.Engine
}

func (H HandlerStr) Routers() {
	Repository := repository.NewRepository(H.DB)
	JdlUC := usecase.NewUsecase(Repository)
	InvUC := usecase.NewUsecaseInv(Repository)
	KtgUC := usecase.NewUsecaseKtg(Repository)

	ep := H.R.Group("masjid")
	handler.JwRouters(JdlUC, ep)
	handler.InvRouters(InvUC, ep)
	handler.KtgRouters(KtgUC, ep)
}
