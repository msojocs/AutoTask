package task

import (
	"log"
	"testing"
)

//func init() {
//
//	wd, err := os.Getwd()
//	if err != nil {
//		log.Panicln("获取工作目录失败！")
//	}
//	bootstrap.Init(wd + "/../../")
//}

func TestSingle(t *testing.T) {
	var task Task
	task.Method = "GET"
	task.Url = "https://api.lolicon.app/setu/v2?uid=123456"
	//task.Proxy = "http://127.0.0.1:8888"
	task.Expected = append(task.Expected, Expected{
		Enable: true,
		Path:   "jsonBody.error",
		Value:  "",
		Exp:    "stringEqual",
	})
	task.Expected = append(task.Expected, Expected{
		Path:  "body.data",
		Value: "0",
		Exp:   "arrayLength",
	})
	task.Expected = append(task.Expected, Expected{
		Enable: true,
		Path:   "status",
		Value:  "200",
		Exp:    "integerEqual",
	})

	_, err := task.Exec()
	if err != nil {
		log.Panicln(err.Error())
		return
	}
}
