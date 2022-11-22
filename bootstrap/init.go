package bootstrap

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/msojocs/AutoTask/v1/db"
	"github.com/msojocs/AutoTask/v1/pkg/conf"
	"gopkg.in/yaml.v3"
)

func Init() {
	log.Println("Bootstrap init start")

	// 处理配置文件路径
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	path = path[:index]

	log.Println("基础路径：", path)
	buf, err := os.ReadFile(filepath.Join(path, "./config.yml"))
	buf, err = os.ReadFile("./config.yml")

	if err != nil {
		log.Fatalln("failed to read config file")
		return
	}
	err = yaml.Unmarshal(buf, &conf.Conf)
	if err != nil {
		log.Fatalln("failed to parse config file")
		return
	}
	db.Init()

	log.Println("Bootstrap init end")
}
