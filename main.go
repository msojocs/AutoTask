package main

import (
	"github.com/msojocs/AutoTask/v1/bootstrap"
	_ "github.com/msojocs/AutoTask/v1/docs"
	router "github.com/msojocs/AutoTask/v1/routers"
	"log"
	"os"
)

// @title Gin swagger
// @version 1.0
// @description Gin swagger AutoTask

// @contact.name msojocs
// @contact.url https://jysafe.cn
// @contact.email msojocs@g mail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
func main() {

	wd, err := os.Getwd()
	if err != nil {
		log.Panicln("获取工作目录失败！")
	}

	bootstrap.Init(wd)
	route := router.SetupRouter()
	err = route.Run()
	if err != nil {
		return
	}
}
