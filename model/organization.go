package model

/// 组织机构表
type Organization struct {
	*Model
	Name string `json:"name" gorm:"column:name"`
	Describe string `json:"describe" gorm:"column:describe"`
	Pid  uint   `json:"pid" gorm:"column:pid"`
	Department []Department `json:"department" gorm:"foreignKey:OrganizationId"`
}

func (Organization) TableName() string {
	return "organization"
}
