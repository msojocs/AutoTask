package task

import (
	"encoding/json"
	"github.com/msojocs/AutoTask/v1/db"
	model "github.com/msojocs/AutoTask/v1/models"
	"log"
)

// Run 执行所有任务
func Run() {
	var requests []model.Request
	db.DB.Find(&requests)
	log.Println("request count: ", len(requests))
	for _, req := range requests {
		var task Task
		err := json.Unmarshal([]byte(req.Main), &task)
		if err != nil {
			return
		}
		result, err := task.Exec()
		if err != nil {
			return
		}
		log.Println("body: ", result.Body)
	}
}
