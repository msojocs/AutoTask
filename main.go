package main

import (
	"github.com/msojocs/AutoTask/v1/bootstrap"
	"github.com/msojocs/AutoTask/v1/cron"
	_ "github.com/msojocs/AutoTask/v1/docs"
	router "github.com/msojocs/AutoTask/v1/routers"
	"log"
	"os"
	"path/filepath"
)

// @title Gin swagger
// @version 1.0
// @description Gin swagger AutoTask

// @contact.name msojocs
// @contact.url https://jysafe.cn
// @contact.email msojocs@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
func main() {

	ed, err := os.Executable()
	wd := filepath.Dir(ed)
	if err != nil {
		log.Panicln("获取程序所在目录失败！")
	}

	bootstrap.Init(wd)
	route := router.SetupRouter()

	//	计划任务
	log.Println("计划任务...")
	cron.Init()

	err = route.Run()
	if err != nil {
		log.Panicln(err.Error())
		return
	}
}
