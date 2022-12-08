package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/msojocs/AutoTask/v1/services/group"
)

func GetAllGroups(ctx *gin.Context) {
	var service group.AllGroupService
	res := service.Get()
	ctx.JSON(200, res)
}
