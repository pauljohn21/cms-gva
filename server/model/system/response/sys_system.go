package response

import "github.com/pauljohn21/cms-gva/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
