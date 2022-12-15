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

	_, err := task.Exec()
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
	task.Body.Type = "file"
	//task.Proxy = "http://127.0.0.1:8888"
	//task.Body.Data = map[string]string{
	//	"file1": "0cdefae3f68eb4bb5a19181a936fa009",
	//}
	//task.Body.Data.FormData
	task.Expected = append(task.Expected, Expected{
		Path:  "body.files.file1",
		Value: "",
		Vtype: "stringNotEmpty",
	})

	_, err := task.Exec()
	if err != nil {
		log.Panicln(err.Error())
		return
	}
}

func TestSwitch(t *testing.T) {

	t1 := "json"

	log.Println("switch start")
	// form/string(json...)/file/binary
	switch t1 {
	case "form-data":
		log.Println("form-data")
		break

	case "form":
		log.Println("form")
		break

	case "json":
		log.Println("raw")
		break
	case "text", "javascript", "html", "xml":
		log.Println("raw")
		break

	case "binary":
		log.Println("binary")
		break
	default:
		log.Println("未知类型：", t1)
		break
	}
	log.Println("switch end")
}
