package response

import "github.com/pauljohn21/cms-gva/server/model/example"

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}
