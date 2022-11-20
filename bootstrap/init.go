package bootstrap

import (
	"os"

	"github.com/msojocs/AutoTask/v1/db"
	"github.com/msojocs/AutoTask/v1/pkg/conf"
	"gopkg.in/yaml.v3"
)

func Init() {
	buf, err := os.ReadFile("../config.yml")
	if err != nil {
		return
	}
	err = yaml.Unmarshal(buf, &conf.Conf)
	if err != nil {
		return
	}
	db.Init()

}
