package system

import (
	"testing"

	"github.com/pauljohn21/cms-gva/server/service"
)

func TestSys(t *testing.T) {
	userService := service.ServiceGroupApp.SystemServiceGroup
	userService.FindUserById(1)
}
