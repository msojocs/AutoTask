package group

import (
	"github.com/gin-gonic/gin"
	"github.com/msojocs/AutoTask/v1/db"
	model "github.com/msojocs/AutoTask/v1/models"
	"github.com/msojocs/AutoTask/v1/pkg/serializer"
)

type AllGroupService struct {
}

func (s AllGroupService) Get() serializer.Response {
	var groupList []model.Group
	db.DB.Find(&groupList)

	return serializer.Response{
		Code: 0,
		Data: map[string]interface{}{
			"group": groupList,
		},
		Msg: "success",
	}

}

type DeleteGroupService struct {
}

func (s DeleteGroupService) Delete(c *gin.Context) bool {
	groupId := c.Param("group_id")
	ret := db.DB.Model(model.Group{}).Table("at_groups").Delete("id = ?", groupId)
	if ret.Error != nil {
		return false
	}
	return true
}

type UpdateGroupService struct {
	Name string `json:"name"`
}

func (s UpdateGroupService) Update(c *gin.Context) error {
	groupId := c.Param("group_id")
	ret := db.DB.Table("at_groups").Where("id = ?", groupId).Update("name", s.Name)
	if ret.Error != nil {
		return ret.Error
	}
	return nil
}
