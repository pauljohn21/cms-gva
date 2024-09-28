package request

import (
	"github.com/pauljohn21/cms-gva/server/model/common/request"
	"time"
)

type ApplicantSearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
