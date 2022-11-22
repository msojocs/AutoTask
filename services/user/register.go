package user

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/msojocs/AutoTask/v1/db"
	model "github.com/msojocs/AutoTask/v1/models"
	"github.com/msojocs/AutoTask/v1/pkg/serializer"
)

// UserRegisterService 管理用户注册的服务
type UserRegisterService struct {
	//TODO 细致调整验证规则
	Email    string `form:"userName" json:"userName" binding:"required,email"`
	Nick     string `form:"nick" json:"nick" binding:"required"`
	Password string `form:"Password" json:"Password" binding:"required,min=4,max=64"`
}

// Register 用户注册函数
func (service *UserRegisterService) Register(c *gin.Context) serializer.Response {
	log.Println("UserRegisterService: ")
	user := model.User{
		Email: service.Email,
	}
	// 密码加密
	user.SetPassword(service.Password)

	// 设置默认用户组
	defaultGroup := model.GetIntSetting("default_group", -1)
	if defaultGroup == -1 {
		return serializer.Err(serializer.CodeGroupNotFound, "failed to get default group", nil)
	}
	user.GroupId = uint(defaultGroup)

	// 添加用户
	if err := db.DB.Create(&user).Error; err != nil {
		expectedUser, err := model.IsEmailExists(service.Email)
		if expectedUser {
			return serializer.Err(serializer.CodeEmailExisted, "user email already in use", err)
		}

	}

	return serializer.Response{}

}
