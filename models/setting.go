package model

import (
	"strconv"

	"github.com/msojocs/AutoTask/v1/db"
)

type Setting struct {
	ID    int64
	Name  string
	Value string
}

func (Setting) TableName() string {
	return "at_settings"
}

// GetSettingByName 用 Name 获取设置值
func GetSettingByName(name string) string {
	var setting Setting
	result := db.DB.Where("name = ?", name).Select("value").First(&setting)
	if result == nil || result.Error == nil {
		return setting.Value
	}
	return ""
}

// GetIntSetting 获取整形设置值，如果转换失败则返回默认值defaultVal
func GetIntSetting(key string, defaultVal int) int {
	res, err := strconv.Atoi(GetSettingByName(key))
	if err != nil {
		return defaultVal
	}
	return res
}
