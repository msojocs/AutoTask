package controllers

import (
	"github.com/msojocs/AutoTask/v1/services/user"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册
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

// UserMenus 获取用户的菜单
func UserMenus(c *gin.Context) {
	c.JSON(http.StatusOK, 123)
}
