package conf

type database struct {
	Host     string `yaml:"host"`
	Port     int64  `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
type conf struct {
	Db   database `yaml:"db"`
	Test int      `yaml:"test"`
}

// 初始化
var Conf conf
var DbConf *database = &Conf.Db
