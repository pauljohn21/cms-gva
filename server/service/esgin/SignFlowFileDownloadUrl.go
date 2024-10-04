package esgin

import (
	"fmt"
	"log"

	"github.com/pauljohn21/cms-gva/server/model/esgin"
	Tools "github.com/pauljohn21/cms-gva/server/utils"
)

func SignFlowFileDownloadUrl(signFlowId string) Tools.Res[esgin.SignFlowFileDownloadUrl] {
	apiUrl := fmt.Sprintf("/v3/sign-flow/%s/file-download-url", signFlowId)
	log.Println("创建签署流程：--------------")
	// var dataJsonStr string
	// if data, err := json.Marshal(signFlowId); err == nil {
	// 	dataJsonStr = string(data)
	// }
	initResult, err := Tools.SendCommHttp[esgin.SignFlowFileDownloadUrl](apiUrl, "", "GET")
	log.Println("返回参数：------------------")
	log.Println(initResult)
	log.Println("错误信息：-----------------------")
	log.Println(err)
	return initResult
}
