package service

import (
	"github.com/pauljohn21/cms-gva/server/service/example"
	"github.com/pauljohn21/cms-gva/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
