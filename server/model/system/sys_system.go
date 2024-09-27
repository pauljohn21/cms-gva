package system

import (
	"github.com/pauljohn21/cms-gva/server/config"
)

// 配置文件结构体
type System struct {
	Config config.Server `json:"config"`
}
