// 自动生成模板MeLetter
package cms

import (
	"github.com/pauljohn21/cms-gva/server/global"
	"time"
)

// 我的保函 结构体  MeLetter
type MeLetter struct {
    global.GVA_MODEL
    Applicant  string `json:"applicant" form:"applicant" gorm:"column:applicant;comment:申请人;size:191;" binding:"required"`  //申请人 
    Code  string `json:"code" form:"code" gorm:"column:code;comment:保单号;size:191;"`  //保单号 
    Court  string `json:"court" form:"court" gorm:"column:court;comment:法院;size:191;" binding:"required"`  //法院 
    Coverage  string `json:"coverage" form:"coverage" gorm:"column:coverage;comment:保费;size:191;" binding:"required"`  //保费 
    CoverageAll  string `json:"coverageAll" form:"coverageAll" gorm:"column:coverageAll;comment:保额;size:191;" binding:"required"`  //保额 
    EndCreatedAt  *time.Time `json:"endCreatedAt" form:"endCreatedAt" gorm:"column:end_created_at;comment:结束时间;"`  //结束时间 
    FileID  string `json:"fileID" form:"fileID" gorm:"column:fileID;comment:文件id;size:191;"`  //文件id 
    Info  string `json:"info" form:"info" gorm:"column:info;comment:保险信息;size:191;"`  //保险信息 
    Policyholder  string `json:"policyholder" form:"policyholder" gorm:"column:policyholder;comment:投保人;size:191;"`  //投保人 
    Respondent  string `json:"respondent" form:"respondent" gorm:"column:respondent;comment:被申请人;size:191;"`  //被申请人 
    SignStatus  string `json:"signStatus" form:"signStatus" gorm:"column:sign_status;comment:;size:191;"`  //签署状态 
    StartCreatedAt  *time.Time `json:"startCreatedAt" form:"startCreatedAt" gorm:"column:start_created_at;comment:开始时间;"`  //开始时间 
    TemplateFileUrl  string `json:"templateFileUrl" form:"templateFileUrl" gorm:"column:template_file_url;comment:文件下载;size:191;"`  //文件下载 
    Type  string `json:"type" form:"type" gorm:"column:type;comment:出单方式;size:191;"`  //出单方式 
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 我的保函 MeLetter自定义表名 me_letter
func (MeLetter) TableName() string {
    return "me_letter"
}

