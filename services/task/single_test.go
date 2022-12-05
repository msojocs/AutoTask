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
		Path:  "body.error",
		Value: "",
		Vtype: "stringEqual",
	})
	task.Expected = append(task.Expected, Expected{
		Path:  "body.data",
		Value: "0",
		Vtype: "arrayLength",
	})
	task.Expected = append(task.Expected, Expected{
		Path:  "status",
		Value: "200",
		Vtype: "integerEqual",
	})

	_, err := task.exec()
	if err != nil {
		log.Panicln(err.Error())
		return
	}
}

// 测试文件上传
func TestUploadFile(t *testing.T) {
	var task Task
	task.Method = "POST"
	task.Url = "https://httpbin.org/post"
	task.Body.t = "file"
	//task.Proxy = "http://127.0.0.1:8888"
	task.Body.data = map[string]string{
		"file1": "0cdefae3f68eb4bb5a19181a936fa009",
	}
	task.Expected = append(task.Expected, Expected{
		Path:  "body.files.file1",
		Value: "",
		Vtype: "stringNotEmpty",
	})

	_, err := task.exec()
	if err != nil {
		log.Panicln(err.Error())
		return
	}
}
