package system

import (
	"fmt"
	"github.com/pauljohn21/cms-gva/server/global"
	"github.com/pauljohn21/cms-gva/server/model/system"
	"testing"
)

func TestUserId(t *testing.T) {
	var user system.SysUser
	err := global.GVA_DB.Where("id = ?", "1").First(&user).Error
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(user.ID)
	}

}
