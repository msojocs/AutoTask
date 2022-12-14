package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/msojocs/AutoTask/v1/pkg/serializer"
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
