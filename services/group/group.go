package group

import (
	"github.com/gin-gonic/gin"
	"github.com/msojocs/AutoTask/v1/db"
	model "github.com/msojocs/AutoTask/v1/models"
	"github.com/msojocs/AutoTask/v1/pkg/serializer"
	"net/http"
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

func (s DeleteGroupService) delete(c *gin.Context) {
	groupId := c.Param("group_id")
	ret := db.DB.Table("at_groups").Delete("id = ?", groupId)
	if ret.Error != nil {

	}
	c.JSON(http.StatusNoContent, nil)
}
