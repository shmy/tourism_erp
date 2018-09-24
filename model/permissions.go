package model

/// 权限表
type Permissions struct {
	*Model
	Name string `json:"name" gorm:"column:name"`
	Describe string `json:"describe" gorm:"column:describe"`
	Code string `json:"code" gorm:"column:code"`
}

func (Permissions) TableName() string {
	return "permissions"
}