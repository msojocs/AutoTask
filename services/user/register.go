package user

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/msojocs/AutoTask/v1/db"
	model "github.com/msojocs/AutoTask/v1/models"
	"github.com/msojocs/AutoTask/v1/pkg/serializer"
)

// RegisterService 管理用户注册的服务
type RegisterService struct {
	//TODO 细致调整验证规则
	Email    string `form:"userName" json:"userName" binding:"required,email"`
	Nick     string `form:"nick" json:"nick" binding:"required"`
	Password string `form:"Password" json:"Password" binding:"required,min=4,max=64"`
}

// Register 用户注册函数
func (service *RegisterService) Register(c *gin.Context) serializer.Response {
	log.Println("UserRegisterService: register start")
	user := model.User{
		Email:  service.Email,
		Status: model.NotActivated,
	}
	// 密码加密
	err := user.SetPassword(service.Password)
	if err != nil {
		return serializer.Response{}
	}

	// 设置默认用户组
	log.Println("UserRegisterService: get default group id")
	defaultGroup := model.GetIntSetting("default_group", -1)
	if defaultGroup == -1 {
		return serializer.Err(serializer.CodeGroupNotFound, "failed to get default group", nil)
	}
	user.GroupId = uint(defaultGroup)

	// 添加用户
	if err := db.DB.Create(&user).Error; err != nil {
		log.Println("创建用户失败！", err.Error())

		if driverErr, ok := err.(*mysql.MySQLError); ok { // 现在可以直接访问错误编号
			if driverErr.Number != 1062 {
				// 处理未知的错误
				return serializer.Err(serializer.CodeNotSet, "unexpected error", err)
			}
		}

		exists, err := model.IsEmailExists(service.Email)
		if exists {
			return serializer.Err(serializer.CodeEmailExisted, "user email already in use", err)
		}

		return serializer.Err(serializer.CodeNotSet, "unexpected error", err)
	} else {
		return serializer.Response{
			Msg: "success",
		}
	}

}
