package controllers

import (
	"log"
	"net/http"
	"strconv"

	model "github.com/msojocs/AutoTask/v1/models"
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
	var user model.User
	err := ctx.ShouldBind(&user)
	if err != nil {
		log.Fatalln("err->", err.Error())
		return
	}
	// 通过
	log.Println("success:", user.Login, " - ", user.Password)
	id := user.Save()
	ctx.String(http.StatusOK, "ok"+strconv.FormatInt(id, 10))
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
