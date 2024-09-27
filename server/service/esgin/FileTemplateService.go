package esgin

import (
	"encoding/json"
	"github.com/pauljohn21/cms-gva/server/model/esgin"
	"log"

	Tools "github.com/pauljohn21/cms-gva/server/utils"
)

func GetFileUploadUrl(data esgin.FileUploadUrlInfo) Tools.Res[esgin.FileTemplatereq] {
	apiUrl := "/v3/files/file-upload-url"
	log.Println("获取文件上传地址：--------------")
	var dataJsonStr string
	if data, err := json.Marshal(data); err == nil {
		dataJsonStr = string(data)
	}
	initResult, err := Tools.SendCommHttp[esgin.FileTemplatereq](apiUrl, dataJsonStr, "POST")
	log.Println("返回参数：------------------")
	log.Println(initResult)
	log.Println("错误信息：-----------------------")
	log.Println(err)
	return initResult
}
