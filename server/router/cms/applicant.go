package cms

import (
	"github.com/pauljohn21/cms-gva/server/api/v1"
	"github.com/pauljohn21/cms-gva/server/middleware"
	"github.com/gin-gonic/gin"
)

type ApplicantRouter struct {}

// InitApplicantRouter 初始化 申请人 路由信息
func (s *ApplicantRouter) InitApplicantRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	applicantRouter := Router.Group("applicant").Use(middleware.OperationRecord())
	applicantRouterWithoutRecord := Router.Group("applicant")
	applicantRouterWithoutAuth := PublicRouter.Group("applicant")

	var applicantApi = v1.ApiGroupApp.CmsApiGroup.ApplicantApi
	{
		applicantRouter.POST("createApplicant", applicantApi.CreateApplicant)   // 新建申请人
		applicantRouter.DELETE("deleteApplicant", applicantApi.DeleteApplicant) // 删除申请人
		applicantRouter.DELETE("deleteApplicantByIds", applicantApi.DeleteApplicantByIds) // 批量删除申请人
		applicantRouter.PUT("updateApplicant", applicantApi.UpdateApplicant)    // 更新申请人
	}
	{
		applicantRouterWithoutRecord.GET("findApplicant", applicantApi.FindApplicant)        // 根据ID获取申请人
		applicantRouterWithoutRecord.GET("getApplicantList", applicantApi.GetApplicantList)  // 获取申请人列表
	}
	{
	    applicantRouterWithoutAuth.GET("getApplicantPublic", applicantApi.GetApplicantPublic)  // 获取申请人列表
	}
}
