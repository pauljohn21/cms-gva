package esgin

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"gorm.io/gorm"

	"github.com/pauljohn21/cms-gva/server/global"
	cms "github.com/pauljohn21/cms-gva/server/model/esgin"
)

func TestCrateTemplate(t *testing.T) {
	data := cms.MeLetter{
		Policyholder:    "丁健",
		Applicant:       "黑龙江省信士正融资担保有限公司",
		Respondent:      "2024-09-27-02.xlsx",
		Info:            "",
		Type:            "",
		SignStatus:      "",
		TemplateFileUrl: "",
		FileID:          "",
		Court:           "哈尔滨市道里区人民法院",
		StartCreatedAt:  &time.Time{},
		EndCreatedAt:    &time.Time{},
		Coverage:        "1500",
		CoverageAll:     "150000000",
	}
	CrateTemplate(data)
	t.Log("test")
}

func TestName(t *testing.T) {
	// 定义变量
	var applicant *cms.Applicant
	var comId string

	// 查询数据
	err := global.GVA_DB.Where("id = ?", "1").First(&applicant).Error

	// 错误处理
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("未找到公司记录")
		} else {
			fmt.Println("查询错误:", err)
		}
	} else {
		comId = applicant.Code
		fmt.Println(comId)
	}
}
