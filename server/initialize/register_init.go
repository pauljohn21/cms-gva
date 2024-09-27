package initialize

import (
	_ "github.com/pauljohn21/cms-gva/server/source/example"
	_ "github.com/pauljohn21/cms-gva/server/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
