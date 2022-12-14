package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/msojocs/AutoTask/v1/services/group"
	"net/http"
)

func GetAllGroups(ctx *gin.Context) {
	var service group.AllGroupService
	res := service.Get()
	ctx.JSON(200, res)
}

func DeleteGroups(ctx *gin.Context) {
	var service group.DeleteGroupService
	service.Delete(ctx)

	ctx.JSON(http.StatusNoContent, nil)
}

func UpdateGroups(ctx *gin.Context) {
	var service group.UpdateGroupService
	if err := ctx.ShouldBindJSON(&service); err == nil {
		err = service.Update(ctx)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
		} else {
			ctx.JSON(http.StatusCreated, nil)
		}
	} else {
		ctx.JSON(http.StatusBadRequest, err)
	}
}
