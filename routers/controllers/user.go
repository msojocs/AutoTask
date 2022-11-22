package controllers

import (
	"log"
	"net/http"

	"github.com/msojocs/AutoTask/v1/services/user"

	"github.com/gin-gonic/gin"
)

func UserSave(ctx *gin.Context) {
	username := ctx.Param("name")
	ctx.String(http.StatusOK, "用户"+username+"已保存")
}

func UserSaveByQuery(ctx *gin.Context) {
	username := ctx.Query("name")
	age := ctx.DefaultQuery("age", "20")
	ctx.String(http.StatusOK, "用户："+username+",年龄："+age+"已保存")
}

func UserRegister(ctx *gin.Context) {

	var service user.UserRegisterService
	if err := ctx.ShouldBindJSON(&service); err == nil {
		res := service.Register(ctx)
		ctx.JSON(200, res)
	} else {
		ctx.JSON(200, ErrorResponse(err))
	}

}

// UserLogin 用户登录
func UserLogin(c *gin.Context) {
	log.Println("Controllers: UserLogin")
	var service user.UserLoginService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
