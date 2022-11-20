package db

import (
	"log"
	"strconv"

	"github.com/msojocs/AutoTask/v1/pkg/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

// 通过bootstrap初始化，使用自带init由于执行顺序问题会造成初始化异常
func Init() {
	var err error
	connectionStr := conf.DbConf.Username + ":" + conf.DbConf.Password + "@tcp(" + conf.DbConf.Host + ":" + strconv.FormatInt(conf.DbConf.Port, 10) + ")/auto_task"
	DB, err = gorm.Open(mysql.Open(connectionStr), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "at_",
			SingularTable: true,
		},
	})
	if err != nil {
		log.Panicln("err:", err.Error())
	}
}
