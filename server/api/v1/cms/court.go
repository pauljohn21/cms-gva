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

type CourtApi struct {}

var courtService = service.ServiceGroupApp.CmsServiceGroup.CourtService


// CreateCourt 创建法院
// @Tags Court
// @Summary 创建法院
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cms.Court true "创建法院"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /court/createCourt [post]
func (courtApi *CourtApi) CreateCourt(c *gin.Context) {
	var court cms.Court
	err := c.ShouldBindJSON(&court)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    court.CreatedBy = utils.GetUserID(c)

	if err := courtService.CreateCourt(&court); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCourt 删除法院
// @Tags Court
// @Summary 删除法院
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cms.Court true "删除法院"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /court/deleteCourt [delete]
func (courtApi *CourtApi) DeleteCourt(c *gin.Context) {
	ID := c.Query("ID")
    	userID := utils.GetUserID(c)
	if err := courtService.DeleteCourt(ID,userID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCourtByIds 批量删除法院
// @Tags Court
// @Summary 批量删除法院
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /court/deleteCourtByIds [delete]
func (courtApi *CourtApi) DeleteCourtByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	if err := courtService.DeleteCourtByIds(IDs,userID); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCourt 更新法院
// @Tags Court
// @Summary 更新法院
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cms.Court true "更新法院"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /court/updateCourt [put]
func (courtApi *CourtApi) UpdateCourt(c *gin.Context) {
	var court cms.Court
	err := c.ShouldBindJSON(&court)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    court.UpdatedBy = utils.GetUserID(c)

	if err := courtService.UpdateCourt(court); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCourt 用id查询法院
// @Tags Court
// @Summary 用id查询法院
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cms.Court true "用id查询法院"
// @Success 200 {object} response.Response{data=object{recourt=cms.Court},msg=string} "查询成功"
// @Router /court/findCourt [get]
func (courtApi *CourtApi) FindCourt(c *gin.Context) {
	ID := c.Query("ID")
	if recourt, err := courtService.GetCourt(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(recourt, c)
	}
}

// GetCourtList 分页获取法院列表
// @Tags Court
// @Summary 分页获取法院列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cmsReq.CourtSearch true "分页获取法院列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /court/getCourtList [get]
func (courtApi *CourtApi) GetCourtList(c *gin.Context) {
	var pageInfo cmsReq.CourtSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := courtService.GetCourtInfoList(pageInfo); err != nil {
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

// GetCourtPublic 不需要鉴权的法院接口
// @Tags Court
// @Summary 不需要鉴权的法院接口
// @accept application/json
// @Produce application/json
// @Param data query cmsReq.CourtSearch true "分页获取法院列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /court/getCourtPublic [get]
func (courtApi *CourtApi) GetCourtPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的法院接口信息",
    }, "获取成功", c)
}
