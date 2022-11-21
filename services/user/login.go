package user

import (
	"log"

	"github.com/gin-gonic/gin"
	model "github.com/msojocs/AutoTask/v1/models"
	"github.com/msojocs/AutoTask/v1/pkg/serializer"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	//TODO 细致调整验证规则
	UserName string `form:"userName" json:"userName" binding:"required,email"`
	Password string `form:"Password" json:"Password" binding:"required,min=4,max=64"`
}

// Login 用户登录函数
func (service *UserLoginService) Login(c *gin.Context) serializer.Response {
	log.Println("UserLoginService: Login")
	expectedUser, err := model.GetUserByEmail(service.UserName)
	// 一系列校验
	if err != nil {
		return serializer.Err(serializer.CodeCredentialInvalid, "Wrong password or email address", err)
	}
	if authOK, _ := expectedUser.CheckPassword(service.Password); !authOK {
		return serializer.Err(serializer.CodeCredentialInvalid, "Wrong password or email address", nil)
	}
	// if expectedUser.Status == model.Baned || expectedUser.Status == model.OveruseBaned {
	// 	return serializer.Err(serializer.CodeUserBaned, "This account has been blocked", nil)
	// }
	// if expectedUser.Status == model.NotActivicated {
	// 	return serializer.Err(serializer.CodeUserNotActivated, "This account is not activated", nil)
	// }

	// if expectedUser.TwoFactor != "" {
	// 	// 需要二步验证
	// 	util.SetSession(c, map[string]interface{}{
	// 		"2fa_user_id": expectedUser.ID,
	// 	})
	// 	return serializer.Response{Code: 203}
	// }

	// //登陆成功，清空并设置session
	// util.SetSession(c, map[string]interface{}{
	// 	"user_id": expectedUser.ID,
	// })

	return serializer.BuildUserResponse(expectedUser)

}
