package serializer

import (
	"time"

	model "github.com/msojocs/AutoTask/v1/models"
)

// User 用户序列化器
type User struct {
	ID             int64     `json:"id"`
	Email          string    `json:"user_name"`
	Nickname       string    `json:"nickname"`
	Status         int       `json:"status"`
	Avatar         string    `json:"avatar"`
	CreatedAt      time.Time `json:"created_at"`
	PreferredTheme string    `json:"preferred_theme"`
	Anonymous      bool      `json:"anonymous"`
	// Group          group     `json:"group"`
	// Tags           []tag     `json:"tags"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	// tags, _ := model.GetTagsByUID(user.ID)
	return User{
		ID:       user.ID,
		Email:    user.Email,
		Nickname: user.Nick,
		Status:   user.Status,
		Avatar:   user.Avatar,
		// CreatedAt:      user.CreatedAt,
		// PreferredTheme: user.OptionsSerialized.PreferredTheme,
		// Anonymous:      user.IsAnonymous(),
		// Group: group{
		// 	ID:                   user.GroupID,
		// 	Name:                 user.Group.Name,
		// 	AllowShare:           user.Group.ShareEnabled,
		// 	AllowRemoteDownload:  user.Group.OptionsSerialized.Aria2,
		// 	AllowArchiveDownload: user.Group.OptionsSerialized.ArchiveDownload,
		// 	ShareDownload:        user.Group.OptionsSerialized.ShareDownload,
		// 	CompressEnabled:      user.Group.OptionsSerialized.ArchiveTask,
		// 	WebDAVEnabled:        user.Group.WebDAVEnabled,
		// 	SourceBatchSize:      user.Group.OptionsSerialized.SourceBatchSize,
		// },
		// Tags: buildTagRes(tags),
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.User) Response {
	return Response{
		Data: BuildUser(user),
	}
}
