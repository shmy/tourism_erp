package model

/// 部门表
type Department struct {
	*Model
	Name string `json:"name" gorm:"column:name"`
	Describe string `json:"describe" gorm:"column:describe"`
	Pid  uint   `json:"pid" gorm:"column:pid"`
	OrganizationId uint `json:"organization_id" gorm:"column:organization_id"`
	Organization Organization `json:"organization" gorm:"foreignKey:OrganizationId"`
}

func (Department) TableName() string {
	return "department"
}