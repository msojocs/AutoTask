package request

import (
	"github.com/gin-gonic/gin"
	"github.com/msojocs/AutoTask/v1/pkg/conf"
	"mime/multipart"
	"os"
)

type FileService struct {
}

func (s FileService) Upload(c *gin.Context, f *multipart.FileHeader) (string, error) {
	userId := c.GetString("user_id")
	path := conf.Conf.Storage.Path + "/" + userId + "/request"
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return "", err
	}
	filePath := path + "/1"
	err = c.SaveUploadedFile(f, filePath)
	if err != nil {
		return "", err
	}
	return filePath, nil
}

func (s FileService) Delete(c *gin.Context, _ string) error {
	userId := c.GetString("user_id")
	path := "/" + userId + "/request/1"
	path = conf.Conf.Storage.Path + path
	if exists, _ := pathExists(path); !exists {
		return nil
	}
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
