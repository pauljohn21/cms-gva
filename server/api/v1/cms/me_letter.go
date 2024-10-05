package cms

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pauljohn21/cms-gva/server/global"
	"github.com/pauljohn21/cms-gva/server/model/cms"
	cmsReq "github.com/pauljohn21/cms-gva/server/model/cms/request"
	"github.com/pauljohn21/cms-gva/server/model/common/response"
	Model "github.com/pauljohn21/cms-gva/server/model/esgin"
	"github.com/pauljohn21/cms-gva/server/service"
	"github.com/pauljohn21/cms-gva/server/service/esgin"
	"github.com/pauljohn21/cms-gva/server/utils"
	"github.com/pauljohn21/cms-gva/server/utils/upload"
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
	// 获取当前工作目录
	currentDir, _ := os.Getwd()

	dataPath := filepath.Join(currentDir, "resource/doc/demo.docx")
	fmt.Println(dataPath)

	var meLetter cms.MeLetter
	err := c.ShouldBindJSON(&meLetter)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println(meLetter)

	meLetter.CreatedBy = utils.GetUserID(c)
	pids, err := esgin.CrateTemplate(&meLetter)
	if err != nil {
		global.GVA_LOG.Error("创建模板失败", zap.Error(err))
		response.FailWithMessage("创建模板失败", c)
		return
	}
	contentMd5, size := utils.CountFileMd5(dataPath)
	fileUploadUrlInfo := Model.FileUploadUrlInfo{
		ContentMd5:   contentMd5,
		ContentType:  "application/octet-stream",
		ConvertToPDF: true,
		FileName:     fmt.Sprintf("%s.docx", pids),
		FileSize:     size,
	}
	initResult, err := esgin.GetFileUploadUrl(fileUploadUrlInfo)
	if err != nil {
		global.GVA_LOG.Error("文件md5上传失败", zap.Error(err))

		response.FailWithMessage(err.Error(), c)
		return
	}
	global.GVA_LOG.Info("文件上传id", zap.Any("文件上传id", initResult.Data.FileId))
	global.GVA_LOG.Info("文件上传URL", zap.Any("文件上传URL", initResult.Data.FileUploadUrl))

	rest, err := utils.UpLoadFile(initResult.Data.FileUploadUrl, dataPath, contentMd5, "application/octet-stream")
	if err != nil {
		global.GVA_LOG.Error("文件上传失败", zap.Error(err))
		response.FailWithMessage("文件上传失败", c)
	}
	if !rest {
		global.GVA_LOG.Error("文件上传失败")
		response.FailWithMessage("文件上传失败", c)

	}
	var PageResult int32

	meLetter.FileID = initResult.Data.FileId
	for {
		pageNumber, err := esgin.GetFileUploadUrlInfo(initResult.Data.FileId)
		if err != nil {
			global.GVA_LOG.Error("文件上传失败", zap.Error(err))
			response.FailWithMessage("文件上传失败", c)
		}
		if pageNumber.Data.FileStatus == 5 {
			PageResult = pageNumber.Data.FileTotalPageCount
			break
		}

		time.Sleep(1 * time.Second)
	}

	pages := strconv.FormatInt(int64(PageResult-4), 10)
	global.GVA_LOG.Info("签署文件页数", zap.Any("签署文件页数:", pages))
	time.Sleep(5 * time.Second)

	// 创建签署流程-start

	flow := Model.SignFlowModel{
		Docs: []Model.Docs{
			{
				FileId:   initResult.Data.FileId,
				FileName: "测试合同.pdf",
			},
		},
		SignFlowConfig: Model.SignFlowConfig{
			SignFlowTitle: "测试合同",
			AutoFinish:    true,
			NotifyUrl:     "",
		},
		Signers: []Model.Signers{
			{
				SignerType: 1,

				SignFields: []Model.SignFields{
					{
						FileId: initResult.Data.FileId,
						NormalSignFieldConfig: Model.NormalSignFieldConfig{
							FreeMode:       false,
							AutoSign:       true,
							AssignedSealId: "1aec676a-74e6-4319-9467-c86e09c5521a",
							SignFieldStyle: "1",
							SignFieldPosition: Model.SignFieldPosition{
								AcrossPageMode: "0",
								PositionPage:   pages,
								PositionX:      480,
								PositionY:      110,
							},
						},
						SignDateConfig: Model.SignDateConfig{
							ShowSignDate: 1,
							DateFormat:   "yyyy年MM月dd日",
							// SignDatePositionX: 360,
							// SignDatePositionY: 80,
						},
					},
				},
			},
		},
	}
	flowresult := esgin.SignFlowCreateByFile(flow)
	fmt.Println("创建签署流程：--------------")
	fmt.Println(flowresult.Data.SignFlowId)
	var flowUrl string
	for {
		FlowId, err := esgin.SignFlowFileDownloadUrl(flowresult.Data.SignFlowId)
		if err != nil {
			global.GVA_LOG.Error("文件上传失败", zap.Error(err))
			response.FailWithMessage("文件上传失败", c)
		}
		if FlowId.Code == 0 {
			flowUrl = FlowId.Data.Files[0].DownloadUrl
			break
		}

		time.Sleep(1 * time.Second)
	}
	fmt.Println(flowUrl)

	filename := upload.DownloadFile(flowUrl, fmt.Sprintf("%s.pdf", pids))

	meLetter.TemplateFileUrl = filename
	// 创建签署流程-end

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
