package v1

import (
	"github.com/pauljohn21/cms-gva/server/api/v1/example"
	"github.com/pauljohn21/cms-gva/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
