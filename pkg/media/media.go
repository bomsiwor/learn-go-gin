package media

import (
	"golang-bootcamp-1/pkg/response"
	"mime/multipart"
)

type IMedia interface {
	Upload(src *multipart.FileHeader, location string) (*string, *response.ErrorResp)
	UpdateFile(identifier string) *response.ErrorResp
	DeleteFile(identifier string) *response.ErrorResp
}
