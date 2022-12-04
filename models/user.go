package model

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"strings"
	"time"

	"github.com/msojocs/AutoTask/v1/db"
	"github.com/msojocs/AutoTask/v1/pkg/util"
	"gorm.io/gorm"
)

const (
	// NotActivated 未激活
	NotActivated = iota
	// Active 账户正常状态
	Active
	// Baned 被封禁
	Baned
)

type User struct {
	ID        int64
	Email     string `form:"email" json:"email"`
	Nick      string
	Password  string `form:"password" json:"password"`
	Status    int
	Avatar    string
	GroupId   uint
	CreatedAt time.Time
}

func (*User) TableName() string {
	return "at_users"
}

func (user *User) Save() int64 {
	db.DB.Create(user)
	return user.ID
}

func (user *User) CheckPassword(pass string) (bool, error) {

	// 原生密码
	if user.Password == pass {
		err := user.SetPassword(pass)
		if err != nil {
			return false, err
		}

		return true, nil
	}

	// 根据存储密码拆分为 Salt 和 Digest
	passwordStore := strings.Split(user.Password, ":")
	if len(passwordStore) != 2 && len(passwordStore) != 3 {
		return false, errors.New("unknown password type")
	}

	// 兼容V2密码，升级后存储格式为: md5:$HASH:$SALT
	if len(passwordStore) == 3 {
		if passwordStore[0] != "md5" {
			return false, errors.New("unknown password type")
		}
		hash := md5.New()
		_, err := hash.Write([]byte(passwordStore[2] + pass))
		bs := hex.EncodeToString(hash.Sum(nil))
		if err != nil {
			return false, err
		}
		return bs == passwordStore[1], nil
	}

	//计算 Salt 和密码组合的SHA1摘要
	hash := sha1.New()
	_, err := hash.Write([]byte(pass + passwordStore[0]))
	bs := hex.EncodeToString(hash.Sum(nil))
	if err != nil {
		return false, err
	}

	return bs == passwordStore[1], nil
}

// SetPassword 根据给定明文设定 User 的 Password 字段
func (user *User) SetPassword(password string) error {
	//生成16位 Salt
	salt := util.RandStringRunes(16)

	//计算 Salt 和密码组合的SHA1摘要
	hash := sha1.New()
	_, err := hash.Write([]byte(password + salt))
	bs := hex.EncodeToString(hash.Sum(nil))

	if err != nil {
		return err
	}

	//存储 Salt 值和摘要， ":"分割
	user.Password = salt + ":" + string(bs)
	return nil
}

// GetUserByEmail 用Email获取用户
func GetUserByEmail(email string) (User, error) {
	var user User
	result := db.DB.Set("gorm:auto_preload", true).Where("email = ?", email).First(&user)
	return user, result.Error
}

func IsEmailExists(email string) (bool, error) {
	var user User
	err := db.DB.Set("gorm:auto_preload", true).Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, err
	}
	return true, err
}
