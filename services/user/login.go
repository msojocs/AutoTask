package user

import (
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
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
	if expectedUser.Status == model.Baned {
		return serializer.Err(serializer.CodeUserBaned, "This account has been blocked", nil)
	}
	if expectedUser.Status == model.NotActivicated {
		return serializer.Err(serializer.CodeUserNotActivated, "This account is not activated", nil)
	}

	// if expectedUser.TwoFactor != "" {
	// 	// 需要二步验证
	// 	util.SetSession(c, map[string]interface{}{
	// 		"2fa_user_id": expectedUser.ID,
	// 	})
	// 	return serializer.Response{Code: 203}
	// }

	// 省略代码
	expiresTime := time.Now().Unix() + int64(24*60*60)
	claims := jwt.RegisteredClaims{
		Audience:  []string{expectedUser.Nick},                   // 受众
		ExpiresAt: jwt.NewNumericDate(time.Unix(expiresTime, 0)), // 失效时间
		IssuedAt:  jwt.NewNumericDate(time.Now()),                // 签发时间
		Issuer:    "gin hello",                                   // 签发人
		NotBefore: jwt.NewNumericDate(time.Now()),                // 生效时间
		Subject:   "login",                                       // 主题
	}
	finalClaims := model.MyCustomClaims{
		Id:               strconv.FormatInt(expectedUser.ID, 10), // 编号
		RegisteredClaims: claims,
	}
	var jwtSecret = []byte("")
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, finalClaims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return serializer.Response{
		Code: 0,
		Data: map[string]string{
			"token": token,
		},
		Msg: "success",
	}

}
