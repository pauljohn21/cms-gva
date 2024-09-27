package initialize

import (
	"github.com/songzhibin97/gkit/cache/local_cache"

	"github.com/pauljohn21/cms-gva/server/global"
	"github.com/pauljohn21/cms-gva/server/utils"
)

func OtherInit() {
	dr, err := utils.ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
	if err != nil {
		panic(err)
	}
	_, err = utils.ParseDuration(global.GVA_CONFIG.JWT.BufferTime)
	if err != nil {
		panic(err)
	}

	global.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(dr),
	)
}
