package esgin

import (
	"encoding/json"
	"log"

	"github.com/pauljohn21/cms-gva/server/model/esgin"

	Tools "github.com/pauljohn21/cms-gva/server/utils"
)

func GetFileUploadUrl(data esgin.FileUploadUrlInfo) (Tools.Res[esgin.FileTemplatereq], error) {
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
	return initResult, err
}

func GetFileUploadUrlInfo(fileid string) (Tools.Res[esgin.GetFileUploadStatusRes], error) {
	apiUrl := "/v3/files/" + fileid
	log.Println("获取文件上传地址：--------------")

	initResult, err := Tools.SendCommHttp[esgin.GetFileUploadStatusRes](apiUrl, "", "GET")
	log.Println("返回参数：------------------")
	log.Println(initResult)
	log.Println("错误信息：-----------------------")
	log.Println(err)
	return initResult, err
}
