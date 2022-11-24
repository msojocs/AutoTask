package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	model "github.com/msojocs/AutoTask/v1/models"
	"github.com/msojocs/AutoTask/v1/pkg/serializer"
	"net/http"
	"strings"
)

func parseToken(token string) (*model.MyCustomClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &model.MyCustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(""), nil
	})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*model.MyCustomClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, err
}

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		result := serializer.Response{
			Code: http.StatusUnauthorized,
			Msg:  "无法认证，重新登录",
			Data: nil,
		}
		auth := context.Request.Header.Get("Authorization")
		if len(auth) == 0 {
			context.Abort()
			context.JSON(http.StatusUnauthorized, gin.H{
				"result": result,
			})
		}
		auth = strings.Fields(auth)[1]
		// 校验token
		_, err := parseToken(auth)
		if err != nil {
			context.Abort()
			result.Msg = "token 过期" + err.Error()
			context.JSON(http.StatusUnauthorized, gin.H{
				"result": result,
			})
		} else {
			println("token 正确")
		}
		context.Next()
	}
}
