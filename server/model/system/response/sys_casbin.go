package response

import (
	"github.com/pauljohn21/cms-gva/server/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
