package model

/// 角色表
type Role struct {
	*Model
	Name string `json:"name" gorm:"column:name"`
	Describe string `json:"describe" gorm:"column:describe"`
}

func (Role) TableName() string {
	return "role"
}