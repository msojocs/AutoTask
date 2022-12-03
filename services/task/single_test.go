package task

import (
	"log"
	"testing"
)

func TestSingle(t *testing.T) {
	var task Task
	task.method = "get"
	task.url = "https://api.lolicon.app/setu/v2?uid=123456"
	//task.proxy = "http://127.0.0.1:8888"
	task.expected = append(task.expected, Expected{
		path:  "body.error",
		value: "",
		vType: "stringEqual",
	})
	task.expected = append(task.expected, Expected{
		path:  "body.data",
		value: "0",
		vType: "arrayLength",
	})
	task.expected = append(task.expected, Expected{
		path:  "status",
		value: "200",
		vType: "integerEqual",
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
	task.method = "POST"
	task.url = "https://httpbin.org/post"
	task.body.t = "file"
	//task.proxy = "http://127.0.0.1:8888"
	task.body.data = map[string]string{
		"file1": "D:/tmp/0cdefae3f68eb4bb5a19181a936fa009",
	}
	task.expected = append(task.expected, Expected{
		path:  "body.files.file1",
		value: "",
		vType: "stringNotEmpty",
	})

	_, err := task.exec()
	if err != nil {
		log.Panicln(err.Error())
		return
	}
}
