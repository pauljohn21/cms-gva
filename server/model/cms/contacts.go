// 自动生成模板Contacts
package cms

import (
	"github.com/pauljohn21/cms-gva/server/global"
)

// 联系人 结构体  Contacts
type Contacts struct {
    global.GVA_MODEL
    Name  string `json:"name" form:"name" gorm:"column:name;comment:;size:191;" binding:"required"`  //姓名 
    Phone  string `json:"phone" form:"phone" gorm:"column:phone;comment:电话;size:19;" binding:"required"`  //电话 
    Addres  string `json:"addres" form:"addres" gorm:"column:addres;comment:地址;size:191;"`  //地址 
    Email  string `json:"email" form:"email" gorm:"column:email;comment:邮箱;size:191;"`  //邮箱 
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 联系人 Contacts自定义表名 contacts
func (Contacts) TableName() string {
    return "contacts"
}

