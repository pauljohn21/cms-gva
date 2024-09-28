package cms

import (
	"github.com/pauljohn21/cms-gva/server/api/v1"
	"github.com/pauljohn21/cms-gva/server/middleware"
	"github.com/gin-gonic/gin"
)

type CourtRouter struct {}

// InitCourtRouter 初始化 法院 路由信息
func (s *CourtRouter) InitCourtRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	courtRouter := Router.Group("court").Use(middleware.OperationRecord())
	courtRouterWithoutRecord := Router.Group("court")
	courtRouterWithoutAuth := PublicRouter.Group("court")

	var courtApi = v1.ApiGroupApp.CmsApiGroup.CourtApi
	{
		courtRouter.POST("createCourt", courtApi.CreateCourt)   // 新建法院
		courtRouter.DELETE("deleteCourt", courtApi.DeleteCourt) // 删除法院
		courtRouter.DELETE("deleteCourtByIds", courtApi.DeleteCourtByIds) // 批量删除法院
		courtRouter.PUT("updateCourt", courtApi.UpdateCourt)    // 更新法院
	}
	{
		courtRouterWithoutRecord.GET("findCourt", courtApi.FindCourt)        // 根据ID获取法院
		courtRouterWithoutRecord.GET("getCourtList", courtApi.GetCourtList)  // 获取法院列表
	}
	{
	    courtRouterWithoutAuth.GET("getCourtPublic", courtApi.GetCourtPublic)  // 获取法院列表
	}
}
