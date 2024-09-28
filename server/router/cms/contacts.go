package cms

import (
	"github.com/pauljohn21/cms-gva/server/api/v1"
	"github.com/pauljohn21/cms-gva/server/middleware"
	"github.com/gin-gonic/gin"
)

type ContactsRouter struct {}

// InitContactsRouter 初始化 联系人 路由信息
func (s *ContactsRouter) InitContactsRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	contactsRouter := Router.Group("contacts").Use(middleware.OperationRecord())
	contactsRouterWithoutRecord := Router.Group("contacts")
	contactsRouterWithoutAuth := PublicRouter.Group("contacts")

	var contactsApi = v1.ApiGroupApp.CmsApiGroup.ContactsApi
	{
		contactsRouter.POST("createContacts", contactsApi.CreateContacts)   // 新建联系人
		contactsRouter.DELETE("deleteContacts", contactsApi.DeleteContacts) // 删除联系人
		contactsRouter.DELETE("deleteContactsByIds", contactsApi.DeleteContactsByIds) // 批量删除联系人
		contactsRouter.PUT("updateContacts", contactsApi.UpdateContacts)    // 更新联系人
	}
	{
		contactsRouterWithoutRecord.GET("findContacts", contactsApi.FindContacts)        // 根据ID获取联系人
		contactsRouterWithoutRecord.GET("getContactsList", contactsApi.GetContactsList)  // 获取联系人列表
	}
	{
	    contactsRouterWithoutAuth.GET("getContactsPublic", contactsApi.GetContactsPublic)  // 获取联系人列表
	}
}
