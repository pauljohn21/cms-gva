package cms

import (
	"github.com/pauljohn21/cms-gva/server/api/v1"
	"github.com/pauljohn21/cms-gva/server/middleware"
	"github.com/gin-gonic/gin"
)

type MeLetterRouter struct {}

// InitMeLetterRouter 初始化 我的保函 路由信息
func (s *MeLetterRouter) InitMeLetterRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	meLetterRouter := Router.Group("meLetter").Use(middleware.OperationRecord())
	meLetterRouterWithoutRecord := Router.Group("meLetter")
	meLetterRouterWithoutAuth := PublicRouter.Group("meLetter")

	var meLetterApi = v1.ApiGroupApp.CmsApiGroup.MeLetterApi
	{
		meLetterRouter.POST("createMeLetter", meLetterApi.CreateMeLetter)   // 新建我的保函
		meLetterRouter.DELETE("deleteMeLetter", meLetterApi.DeleteMeLetter) // 删除我的保函
		meLetterRouter.DELETE("deleteMeLetterByIds", meLetterApi.DeleteMeLetterByIds) // 批量删除我的保函
		meLetterRouter.PUT("updateMeLetter", meLetterApi.UpdateMeLetter)    // 更新我的保函
	}
	{
		meLetterRouterWithoutRecord.GET("findMeLetter", meLetterApi.FindMeLetter)        // 根据ID获取我的保函
		meLetterRouterWithoutRecord.GET("getMeLetterList", meLetterApi.GetMeLetterList)  // 获取我的保函列表
	}
	{
	    meLetterRouterWithoutAuth.GET("getMeLetterPublic", meLetterApi.GetMeLetterPublic)  // 获取我的保函列表
	}
}
