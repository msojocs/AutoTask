package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/msojocs/AutoTask/v1/pkg/serializer"
	"github.com/msojocs/AutoTask/v1/services/request"
	"github.com/msojocs/AutoTask/v1/services/task"
	"net/http"
)

func Test(c *gin.Context) {
	var t task.Task
	if err := c.ShouldBindJSON(&t); err == nil {
		resp, err2 := t.Exec()
		if err2 != nil {
			return
		}
		c.JSON(http.StatusOK, serializer.Response{
			Code: 0,
			Data: resp,
		})
	} else {
		c.JSON(http.StatusBadRequest, serializer.Err(1, "数据解析失败", err))
	}
}

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err == nil {
		var s request.FileService
		path, err := s.Upload(c, file)
		if err != nil {
			c.JSON(http.StatusOK, serializer.Err(1, "文件存储失败", err))
			return
		}
		ret := make(map[string]string)
		ret["path"] = path
		c.JSON(http.StatusOK, serializer.Response{
			Code: 0,
			Data: ret,
		})
	} else {
		c.JSON(http.StatusBadRequest, serializer.Err(2, "未上传指定文件", err))
	}
}

func DeleteFile(c *gin.Context) {
	file := c.Query("file")
	if len(file) > 0 {
		var s request.FileService
		err := s.Delete(c, file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, serializer.Err(1, "文件删除失败", err))
			return
		}
		c.JSON(http.StatusNoContent, nil)
	} else {
		c.JSON(http.StatusBadRequest, serializer.Err(2, "参数错误", nil))
	}
}
