package task

import (
	"log"
	"testing"
)

func TestSingle(t *testing.T) {
	var task Task
	task.method = "GET"
	task.url = "https://api.lolicon.app/setu/v2?uid=123456"
	task.expected = append(task.expected, Expected{
		path:  "body.error",
		value: "",
		vType: "string",
	})
	task.expected = append(task.expected, Expected{
		path:  "body.data",
		value: "0",
		vType: "arrayLength",
	})
	task.expected = append(task.expected, Expected{
		path:  "status",
		value: "200",
		vType: "integer",
	})

	_, err := task.exec()
	if err != nil {
		log.Panicln(err.Error())
		return
	}
}
