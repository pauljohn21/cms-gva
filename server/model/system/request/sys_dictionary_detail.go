package request

import (
	"github.com/pauljohn21/cms-gva/server/model/common/request"
	"github.com/pauljohn21/cms-gva/server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
