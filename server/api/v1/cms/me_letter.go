package cms

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pauljohn21/cms-gva/server/global"
	"github.com/pauljohn21/cms-gva/server/model/cms"
	cmsReq "github.com/pauljohn21/cms-gva/server/model/cms/request"
	"github.com/pauljohn21/cms-gva/server/model/common/response"
	"github.com/pauljohn21/cms-gva/server/service"
	"github.com/pauljohn21/cms-gva/server/service/esgin"
	"github.com/pauljohn21/cms-gva/server/utils"
	"go.uber.org/zap"
)

type MeLetterApi struct{}

var meLetterService = service.ServiceGroupApp.CmsServiceGroup.MeLetterService

// CreateMeLetter 创建我的保函
// @Tags MeLetter
// @Summary 创建我的保函
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cms.MeLetter true "创建我的保函"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /meLetter/createMeLetter [post]
func (meLetterApi *MeLetterApi) CreateMeLetter(c *gin.Context) {
	var meLetter cms.MeLetter
	err := c.ShouldBindJSON(&meLetter)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println(meLetter)

	meLetter.CreatedBy = utils.GetUserID(c)
	page, err := esgin.CrateTemplate(&meLetter)
	fmt.Printf("签章页%d", page)

	if err := meLetterService.CreateMeLetter(&meLetter); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMeLetter 删除我的保函
// @Tags MeLetter
// @Summary 删除我的保函
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cms.MeLetter true "删除我的保函"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /meLetter/deleteMeLetter [delete]
func (meLetterApi *MeLetterApi) DeleteMeLetter(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := meLetterService.DeleteMeLetter(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMeLetterByIds 批量删除我的保函
// @Tags MeLetter
// @Summary 批量删除我的保函
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /meLetter/deleteMeLetterByIds [delete]
func (meLetterApi *MeLetterApi) DeleteMeLetterByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := meLetterService.DeleteMeLetterByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMeLetter 更新我的保函
// @Tags MeLetter
// @Summary 更新我的保函
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cms.MeLetter true "更新我的保函"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /meLetter/updateMeLetter [put]
func (meLetterApi *MeLetterApi) UpdateMeLetter(c *gin.Context) {
	var meLetter cms.MeLetter
	err := c.ShouldBindJSON(&meLetter)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	meLetter.UpdatedBy = utils.GetUserID(c)

	if err := meLetterService.UpdateMeLetter(meLetter); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMeLetter 用id查询我的保函
// @Tags MeLetter
// @Summary 用id查询我的保函
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cms.MeLetter true "用id查询我的保函"
// @Success 200 {object} response.Response{data=object{remeLetter=cms.MeLetter},msg=string} "查询成功"
// @Router /meLetter/findMeLetter [get]
func (meLetterApi *MeLetterApi) FindMeLetter(c *gin.Context) {
	ID := c.Query("ID")
	if remeLetter, err := meLetterService.GetMeLetter(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(remeLetter, c)
	}
}

// GetMeLetterList 分页获取我的保函列表
// @Tags MeLetter
// @Summary 分页获取我的保函列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cmsReq.MeLetterSearch true "分页获取我的保函列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /meLetter/getMeLetterList [get]
func (meLetterApi *MeLetterApi) GetMeLetterList(c *gin.Context) {
	var pageInfo cmsReq.MeLetterSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := meLetterService.GetMeLetterInfoList(pageInfo); err != nil {
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

// GetMeLetterPublic 不需要鉴权的我的保函接口
// @Tags MeLetter
// @Summary 不需要鉴权的我的保函接口
// @accept application/json
// @Produce application/json
// @Param data query cmsReq.MeLetterSearch true "分页获取我的保函列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /meLetter/getMeLetterPublic [get]
func (meLetterApi *MeLetterApi) GetMeLetterPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的我的保函接口信息",
	}, "获取成功", c)
}
