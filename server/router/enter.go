package router

import (
	"github.com/pauljohn21/cms-gva/server/router/cms"
	"github.com/pauljohn21/cms-gva/server/router/example"
	"github.com/pauljohn21/cms-gva/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Cms     cms.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
