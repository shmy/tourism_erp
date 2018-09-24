package model

/// 人员表
type People struct {
	*Model
	UserName string `json:"username" gorm:"column:username"`
	Password string `json:"-" gorm:"column:password"`
	RealName  string   `json:"realname" gorm:"column:realname"`
	Avatar string `json:"avatar" gorm:"column:avatar"`
	Phone string `json:"phone" gorm:"column:phone"`
	SocketId string `json:"socket_id" gorm:"column:socket_id"`
	Locked bool `json:"locked" gorm:"column:locked"`
	LockWhy string `json:"lock_why" gorm:"column:lock_why"`
}

func (People) TableName() string {
	return "people"
}