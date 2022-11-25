package router

import (
	"github.com/msojocs/AutoTask/v1/middleware"
	"net/http"
	"strings"
	"time"

	controller "github.com/msojocs/AutoTask/v1/routers/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func retHelloGinAndMethod(ctx *gin.Context) {
	ctx.String(http.StatusOK, "hello gin "+strings.ToLower(ctx.Request.Method)+" method")
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// 跨域配置
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	api := router.Group("/api")
	api.GET("/", retHelloGinAndMethod)
	api.GET("/test", middleware.Auth(), func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"test": 123,
		})
	})
	api.POST("/", retHelloGinAndMethod)
	api.PATCH("/", retHelloGinAndMethod)
	user := api.Group("/user")
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
