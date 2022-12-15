package request

import (
	"github.com/gin-gonic/gin"
	"github.com/msojocs/AutoTask/v1/pkg/conf"
	"mime/multipart"
)

type FileUploadService struct {
}

func (s FileUploadService) Upload(c *gin.Context, f *multipart.FileHeader) (string, error) {
	path := "/request/1"
	err := c.SaveUploadedFile(f, conf.Conf.Storage.Path+path)
	if err != nil {
		return "", err
	}
	return path, nil
}
