package model

type MergeOrganizationDepartmentPeople struct {
	*Model
	OrganizationId uint `json:"organization_id" gorm:"column:organization_id"`
	DepartmentId uint `json:"department_id" gorm:"column:department_id"`
	PeopleId uint `json:"people_id" gorm:"column:people_id"`
}
/// 人员组织部门表
func (MergeOrganizationDepartmentPeople) TableName() string {
	return "merge_organization_department_people"
}