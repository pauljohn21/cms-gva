package request

import (
	"time"

	"github.com/pauljohn21/cms-gva/server/model/common/request"
	"github.com/pauljohn21/cms-gva/server/model/system"
)

type SysExportTemplateSearch struct {
	system.SysExportTemplate
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
