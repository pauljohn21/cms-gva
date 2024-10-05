package esgin

// 获取文件上传地址json信息配置
type FileUploadUrlInfo struct {
	ContentMd5   string `json:"contentMd5,omitempty"`
	ContentType  string `json:"contentType,omitempty"`
	ConvertToPDF bool   `json:"convertToPDF,omitempty"`
	FileName     string `json:"fileName,omitempty"`
	FileSize     int64  `json:"fileSize,omitempty"`
}
type FileTemplatereq struct {
	FileId        string `json:"fileId"`
	FileUploadUrl string `json:"fileUploadUrl"`
}
type GetFileUploadStatusRes struct {
	FileId             string  `json:"fileId"`
	FileName           string  `json:"fileName"`
	FileSize           int32   `json:"fileSize"`
	FileStatus         int32   `json:"fileStatus"`
	FileDownloadUrl    string  `json:"fileDownloadUrl"`
	FileTotalPageCount int32   `json:"fileTotalPageCount"`
	PageWidth          float32 `json:"pageWidth"`
	PageHeight         float32 `json:"pageHeight"`
}
type GetFileUploadStatus struct {
	FileId   string `json:"fileId"`
	PageSize bool   `json:"pageSize"`
}
