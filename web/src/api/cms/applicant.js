import service from '@/utils/request'

// @Tags Applicant
// @Summary 创建申请人
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Applicant true "创建申请人"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /applicant/createApplicant [post]
export const createApplicant = (data) => {
  return service({
    url: '/applicant/createApplicant',
    method: 'post',
    data
  })
}

// @Tags Applicant
// @Summary 删除申请人
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Applicant true "删除申请人"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /applicant/deleteApplicant [delete]
export const deleteApplicant = (params) => {
  return service({
    url: '/applicant/deleteApplicant',
    method: 'delete',
    params
  })
}

// @Tags Applicant
// @Summary 批量删除申请人
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除申请人"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /applicant/deleteApplicant [delete]
export const deleteApplicantByIds = (params) => {
  return service({
    url: '/applicant/deleteApplicantByIds',
    method: 'delete',
    params
  })
}

// @Tags Applicant
// @Summary 更新申请人
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Applicant true "更新申请人"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /applicant/updateApplicant [put]
export const updateApplicant = (data) => {
  return service({
    url: '/applicant/updateApplicant',
    method: 'put',
    data
  })
}

// @Tags Applicant
// @Summary 用id查询申请人
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Applicant true "用id查询申请人"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /applicant/findApplicant [get]
export const findApplicant = (params) => {
  return service({
    url: '/applicant/findApplicant',
    method: 'get',
    params
  })
}

// @Tags Applicant
// @Summary 分页获取申请人列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取申请人列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /applicant/getApplicantList [get]
export const getApplicantList = (params) => {
  return service({
    url: '/applicant/getApplicantList',
    method: 'get',
    params
  })
}
