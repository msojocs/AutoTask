package task

import (
	"encoding/json"
	"github.com/msojocs/AutoTask/v1/bootstrap"
	"github.com/msojocs/AutoTask/v1/db"
	model "github.com/msojocs/AutoTask/v1/models"
	"log"
	"os"
	"testing"
)

func init() {

	wd, err := os.Getwd()
	if err != nil {
		log.Panicln("获取工作目录失败！")
	}
	bootstrap.Init(wd + "/../../")
}

func TestRequest(t *testing.T) {
	var req []model.Request
	_db := db.DB.Find(&req)
	if _db.Error != nil {
		log.Panicln(_db.Error.Error())
	}
	log.Println("数量：", len(req))

	var task Task
	err := json.Unmarshal([]byte(req[0].Main), &task)
	if err != nil {
		log.Panicln(err.Error())
		return
	}
	result, err := task.Exec()
	if err != nil {
		log.Panicln(err.Error())
		return
	}
	log.Println(result)
}

func TestType(t *testing.T) {
	var a interface{} = "123"
	switch a.(type) {
	case int64:
		log.Println("int64")
		break
	default:
		log.Println("unknown")
		break

	}
}
