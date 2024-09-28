// 自动生成模板Applicant
package cms

import (
	"github.com/pauljohn21/cms-gva/server/global"
)

// 申请人 结构体  Applicant
type Applicant struct {
    global.GVA_MODEL
    Code  string `json:"code" form:"code" gorm:"column:code;comment:统一社会信用代码;size:191;" binding:"required"`  //统一社会信用代码 
    Company  string `json:"company" form:"company" gorm:"column:company;comment:公司名;size:191;" binding:"required"`  //公司名 
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 申请人 Applicant自定义表名 applicant
func (Applicant) TableName() string {
    return "applicant"
}

