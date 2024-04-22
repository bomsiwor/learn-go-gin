package media

import (
	"golang-bootcamp-1/pkg/response"
	"mime/multipart"
)

type IMedia interface {
	Upload(src *multipart.FileHeader, location string) (*RemoteResponse, *response.ErrorResp)
	UpdateFile(identifier string) *response.ErrorResp
	DeleteFile(identifier string) *response.ErrorResp
}

type RemoteResponse struct {
	Url        *string
	Identifier *string
}
