package model

/// 人员表
type Token struct {
	*Model
	Token  string   `json:"token" gorm:"column:token"`
	PeopleId uint `json:"people_id" gorm:"column:people_id"`
	Type int `json:"type" gorm:"column:type"`
}

func (Token) TableName() string {
	return "token"
}