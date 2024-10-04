package esgin

type (
	SignFlowModel struct {
		Docs           []Docs         `json:"docs"`
		SignFlowConfig SignFlowConfig `json:"signFlowConfig"`
		Signers        []Signers      `json:"signers"`
	}
	Docs struct {
		FileId   string `json:"fileId"`
		FileName string `json:"fileName"`
	}
	SignFlowConfig struct {
		SignFlowTitle string `json:"signFlowTitle"`
		// SignFlowExpireTime int64  `json:"signFlowExpireTime"`
		AutoFinish bool   `json:"autoFinish"`
		NotifyUrl  string `json:"notifyUrl"`
	}
	Signers struct {
		SignerType int32        `json:"signerType"`
		SignFields []SignFields `json:"signFields"`
	}
	SignFields struct {
		FileId                string                `json:"fileId"`
		NormalSignFieldConfig NormalSignFieldConfig `json:"normalSignFieldConfig"`
		SignDateConfig        SignDateConfig        `json:"signDateConfig"`
		// SignFieldType         int32                 `json:"signFieldType"`
	}
	NormalSignFieldConfig struct {
		FreeMode          bool              `json:"freeMode"`
		AutoSign          bool              `json:"autoSign"`
		AssignedSealId    string            `json:"assignedSealId"`
		SignFieldSize     string            `json:"signFieldSize"`
		SignFieldStyle    string            `json:"signFieldStyle"`
		SignFieldPosition SignFieldPosition `json:"signFieldPosition"`
	}
	SignFieldPosition struct {
		AcrossPageMode string  `json:"acrossPageMode"`
		PositionPage   string  `json:"positionPage"`
		PositionX      float32 `json:"positionX"`
		PositionY      float32 `json:"positionY"`
	}
	OrgSignerInfo struct {
		OrgName string `json:"orgName"`
	}
	SignDateConfig struct {
		ShowSignDate int32  `json:"showSignDate"`
		DateFormat   string `json:"dateFormat"`
		// SignDatePositionX float32 `json:"signDatePositionX"`
		// SignDatePositionY float32 `json:"signDatePositionY"`
	}

	SignFlowReq struct {
		SignFlowId string `json:"signFlowId"`
	}
	SignFlowFileDownloadUrl struct {
		Files []SignFlowFileDownloadUrlFile `json:"files"`
	}
	SignFlowFileDownloadUrlFile struct {
		FileId      string `json:"fileId"`
		FileName    string `json:"fileName"`
		DownloadUrl string `json:"downloadUrl"`
	}
)
