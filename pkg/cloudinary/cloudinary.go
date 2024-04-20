package cloudinary

import (
	"os"

	cldPkg "github.com/cloudinary/cloudinary-go/v2"
)

func Init() *cldPkg.Cloudinary {
	cld, _ := cldPkg.NewFromParams(
		os.Getenv("CLOUDINARY_CLOUD_NAME"),
		os.Getenv("CLOUDINARY_API_KEY"),
		os.Getenv("CLOUDINARY_API_SECRET"),
	)

	return cld
}
