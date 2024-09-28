package cms

import (
	"github.com/pauljohn21/cms-gva/server/global"
    "github.com/pauljohn21/cms-gva/server/model/cms"
    cmsReq "github.com/pauljohn21/cms-gva/server/model/cms/request"
    "github.com/pauljohn21/cms-gva/server/model/common/response"
    "github.com/pauljohn21/cms-gva/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/pauljohn21/cms-gva/server/utils"
)

type ApplicantApi struct {}

var applicantService = service.ServiceGroupApp.CmsServiceGroup.ApplicantService


// CreateApplicant 创建申请人
// @Tags Applicant
// @Summary 创建申请人
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cms.Applicant true "创建申请人"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /applicant/createApplicant [post]
func (applicantApi *ApplicantApi) CreateApplicant(c *gin.Context) {
	var applicant cms.Applicant
	err := c.ShouldBindJSON(&applicant)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    applicant.CreatedBy = utils.GetUserID(c)

	if err := applicantService.CreateApplicant(&applicant); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteApplicant 删除申请人
// @Tags Applicant
// @Summary 删除申请人
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cms.Applicant true "删除申请人"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /applicant/deleteApplicant [delete]
func (applicantApi *ApplicantApi) DeleteApplicant(c *gin.Context) {
	ID := c.Query("ID")
    	userID := utils.GetUserID(c)
	if err := applicantService.DeleteApplicant(ID,userID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteApplicantByIds 批量删除申请人
// @Tags Applicant
// @Summary 批量删除申请人
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /applicant/deleteApplicantByIds [delete]
func (applicantApi *ApplicantApi) DeleteApplicantByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	if err := applicantService.DeleteApplicantByIds(IDs,userID); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateApplicant 更新申请人
// @Tags Applicant
// @Summary 更新申请人
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cms.Applicant true "更新申请人"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /applicant/updateApplicant [put]
func (applicantApi *ApplicantApi) UpdateApplicant(c *gin.Context) {
	var applicant cms.Applicant
	err := c.ShouldBindJSON(&applicant)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    applicant.UpdatedBy = utils.GetUserID(c)

	if err := applicantService.UpdateApplicant(applicant); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindApplicant 用id查询申请人
// @Tags Applicant
// @Summary 用id查询申请人
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cms.Applicant true "用id查询申请人"
// @Success 200 {object} response.Response{data=object{reapplicant=cms.Applicant},msg=string} "查询成功"
// @Router /applicant/findApplicant [get]
func (applicantApi *ApplicantApi) FindApplicant(c *gin.Context) {
	ID := c.Query("ID")
	if reapplicant, err := applicantService.GetApplicant(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(reapplicant, c)
	}
}

// GetApplicantList 分页获取申请人列表
// @Tags Applicant
// @Summary 分页获取申请人列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cmsReq.ApplicantSearch true "分页获取申请人列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /applicant/getApplicantList [get]
func (applicantApi *ApplicantApi) GetApplicantList(c *gin.Context) {
	var pageInfo cmsReq.ApplicantSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := applicantService.GetApplicantInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}

// GetApplicantPublic 不需要鉴权的申请人接口
// @Tags Applicant
// @Summary 不需要鉴权的申请人接口
// @accept application/json
// @Produce application/json
// @Param data query cmsReq.ApplicantSearch true "分页获取申请人列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /applicant/getApplicantPublic [get]
func (applicantApi *ApplicantApi) GetApplicantPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的申请人接口信息",
    }, "获取成功", c)
}
