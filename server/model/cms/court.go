// 自动生成模板Court
package cms

import (
	"github.com/pauljohn21/cms-gva/server/global"
)

// 法院 结构体  Court
type Court struct {
    global.GVA_MODEL
    Addr  string `json:"addr" form:"addr" gorm:"column:addr;comment:地址;size:191;" binding:"required"`  //地址 
    Name  string `json:"name" form:"name" gorm:"column:name;comment:法院名;size:191;" binding:"required"`  //法院名 
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 法院 Court自定义表名 court
func (Court) TableName() string {
    return "court"
}

