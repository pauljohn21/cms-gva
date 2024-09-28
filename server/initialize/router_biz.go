package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/pauljohn21/cms-gva/server/router"
)

func holder(routers ...*gin.RouterGroup) {
	_ = routers
	_ = router.RouterGroupApp
}

func initBizRouter(routers ...*gin.RouterGroup) {
	privateGroup := routers[0]
	publicGroup := routers[1]
	{
		cmsRouter := router.RouterGroupApp.Cms
		cmsRouter.InitApplicantRouter(privateGroup, publicGroup)

		cmsRouter.InitCourtRouter(privateGroup, publicGroup)
		cmsRouter.InitMeLetterRouter(privateGroup, publicGroup)
		cmsRouter.InitContactsRouter(privateGroup, publicGroup)
	}

	holder(publicGroup, privateGroup)
}
