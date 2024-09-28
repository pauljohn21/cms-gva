import service from '@/utils/request'

// @Tags Contacts
// @Summary 创建联系人
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Contacts true "创建联系人"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /contacts/createContacts [post]
export const createContacts = (data) => {
  return service({
    url: '/contacts/createContacts',
    method: 'post',
    data
  })
}

// @Tags Contacts
// @Summary 删除联系人
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Contacts true "删除联系人"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /contacts/deleteContacts [delete]
export const deleteContacts = (params) => {
  return service({
    url: '/contacts/deleteContacts',
    method: 'delete',
    params
  })
}

// @Tags Contacts
// @Summary 批量删除联系人
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除联系人"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /contacts/deleteContacts [delete]
export const deleteContactsByIds = (params) => {
  return service({
    url: '/contacts/deleteContactsByIds',
    method: 'delete',
    params
  })
}

// @Tags Contacts
// @Summary 更新联系人
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Contacts true "更新联系人"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /contacts/updateContacts [put]
export const updateContacts = (data) => {
  return service({
    url: '/contacts/updateContacts',
    method: 'put',
    data
  })
}

// @Tags Contacts
// @Summary 用id查询联系人
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Contacts true "用id查询联系人"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /contacts/findContacts [get]
export const findContacts = (params) => {
  return service({
    url: '/contacts/findContacts',
    method: 'get',
    params
  })
}

// @Tags Contacts
// @Summary 分页获取联系人列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取联系人列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /contacts/getContactsList [get]
export const getContactsList = (params) => {
  return service({
    url: '/contacts/getContactsList',
    method: 'get',
    params
  })
}
