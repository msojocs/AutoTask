package router

import (
	"net/http"
	"strings"

	controller "github.com/msojocs/AutoTask/v1/routers/controllers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func retHelloGinAndMethod(ctx *gin.Context) {
	ctx.String(http.StatusOK, "hello gin "+strings.ToLower(ctx.Request.Method)+" method")
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", retHelloGinAndMethod)
	router.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"test": 123,
		})
	})
	router.POST("/", retHelloGinAndMethod)
	router.PATCH("/", retHelloGinAndMethod)
	user := router.Group("/user")
	{

		user.GET("/:name", controller.UserSave)
		user.GET("", controller.UserSaveByQuery)
		user.POST("/register", controller.UserRegister)
		user.POST("/login", controller.UserLogin)
	}
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
