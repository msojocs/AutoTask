package conf

type database struct {
	Host     string `yaml:"host"`
	Port     int64  `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
type jwt struct {
	Secret string `yaml:"secret"`
}
type conf struct {
	Db  database `yaml:"db"`
	Jwt jwt      `yaml:"jwt"`
}

// Conf 初始化
var Conf conf
var DbConf *database = &Conf.Db
