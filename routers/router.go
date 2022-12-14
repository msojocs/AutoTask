package router

import (
	"github.com/msojocs/AutoTask/v1/middleware"
	"time"

	controller "github.com/msojocs/AutoTask/v1/routers/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// 跨域配置
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	api := router.Group("/api/v1")
	// 用户相关操作
	user := api.Group("/user")
	{
		// 用户注册
		user.POST("/register", middleware.Auth(), controller.UserRegister)
		// 用户登录
		user.POST("/login", controller.UserLogin)
		// 获取用户菜单
		user.GET("/menus", middleware.Auth(), controller.UserMenus)
	}
	// 用户组相关
	authGroup := api.Group("/auth/group") //, middleware.Auth()
	{
		// 获取所有用户组
		authGroup.GET("", controller.GetAllGroups)
		// 创建用户组
		authGroup.POST("/create", nil)
		// 修改用户组信息
		authGroup.PUT("/:group_id", controller.UpdateGroups)
		// 删除用户组
		authGroup.DELETE("/:group_id", controller.DeleteGroups)
		// 修改用户组下的菜单
		authGroup.PUT("/menus", nil)
	}
	// 菜单相关
	authMenu := api.Group("/auth/menu", middleware.Auth())
	{
		// 获取所有菜单
		authMenu.GET("", nil)
		// 创建菜单
		authMenu.POST("/create", nil)
		// 修改菜单
		authMenu.PUT("/:menu_id", nil)
		// 删除菜单
		authMenu.DELETE("/:menu_id", nil)
	}
	// 任务相关
	job := api.Group("/job", middleware.Auth())
	{
		// 获取任务信息
		job.GET("/:job_id", nil)
		// 创建任务
		job.POST("/create", nil)
		// 修改任务信息
		job.PUT("/:job_id", nil)
		// 删除任务
		job.DELETE("/:job_id", nil)
	}
	// 请求相关
	req := api.Group("/request")
	{
		// 获取请求信息
		req.GET("/:request_id", middleware.Auth(), nil)
		// 创建请求
		req.POST("/create", middleware.Auth(), nil)
		// 修改请求
		req.PUT("/:request_id", middleware.Auth(), nil)
		//	删除请求
		req.DELETE("/:request_id", middleware.Auth(), nil)
		// 请求测试
		req.POST("/test", controller.Test)
	}

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
