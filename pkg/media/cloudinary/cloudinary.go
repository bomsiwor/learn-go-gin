package media

import (
	"context"
	globalMedia "golang-bootcamp-1/pkg/media"
	"golang-bootcamp-1/pkg/response"
	"mime/multipart"
	"os"
	"strings"

	cldPkg "github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

type mediaUsecase struct {
}

func startConn() (*cldPkg.Cloudinary, context.Context) {
	cld, _ := cldPkg.NewFromParams(
		os.Getenv("CLOUDINARY_CLOUD_NAME"),
		os.Getenv("CLOUDINARY_API_KEY"),
		os.Getenv("CLOUDINARY_API_SECRET"),
	)
	cld.Config.URL.Secure = true

	ctx := context.Background()

	return cld, ctx
}

// DeleteFile implements media.IMedia.
func (*mediaUsecase) DeleteFile(identifier string) *response.ErrorResp {
	panic("unimplemented")
}

// UpdateFile implements media.IMedia.
func (*mediaUsecase) UpdateFile(identifier string) *response.ErrorResp {
	panic("unimplemented")
}

// Upload implements media.IMedia.
func (*mediaUsecase) Upload(src *multipart.FileHeader, location string) (string, *response.ErrorResp) {
	// Init connection
	cld, ctx := startConn()

	// Adapt from gin saveFileTo function
	// Open file first

	file, err := src.Open()
	if err != nil {
		return "", &response.ErrorResp{
			Code:    500,
			Message: "Error proccessing file",
			Err:     err,
		}
	}

	defer file.Close()

	// Uplaod file using uploader
	publicId := strings.Join([]string{location, uuid.NewString()}, "/")
	if file != nil {
		resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
			PublicID:       publicId,
			UniqueFilename: api.Bool(false),
			Overwrite:      api.Bool(true),
		})
		if err != nil {
			return "", &response.ErrorResp{
				Code:    500,
				Message: "Error proccessing file",
				Err:     err,
			}
		}

		return resp.SecureURL, nil
	}

	return "", &response.ErrorResp{
		Code:    500,
		Message: "Error proccessing file",
		Err:     err,
	}
}

func NewMediaUsecase() globalMedia.IMedia {
	return &mediaUsecase{}
}
