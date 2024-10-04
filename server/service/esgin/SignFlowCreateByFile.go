package esgin

import (
	"encoding/json"
	"log"

	"github.com/pauljohn21/cms-gva/server/model/esgin"
	Tools "github.com/pauljohn21/cms-gva/server/utils"
)

func SignFlowCreateByFile(flow esgin.SignFlowModel) Tools.Res[esgin.SignFlowReq] {
	apiUrl := "/v3/sign-flow/create-by-file"
	log.Println("创建签署流程：--------------")
	var dataJsonStr string
	if data, err := json.Marshal(flow); err == nil {
		dataJsonStr = string(data)
	}
	initResult, err := Tools.SendCommHttp[esgin.SignFlowReq](apiUrl, dataJsonStr, "POST")
	log.Println("返回参数：------------------")
	log.Println(initResult)
	log.Println("错误信息：-----------------------")
	log.Println(err)
	return initResult
}
