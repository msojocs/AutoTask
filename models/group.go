package model

type Group struct {
	ID   int64
	Name string
}

func (Group) TableName() string {
	return "at_groups"
}
