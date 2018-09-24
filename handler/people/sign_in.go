package people

import (
	"errors"
	"github.com/jinzhu/gorm"
	"tourism_erp/model"
	"tourism_erp/service/people"
	"tourism_erp/util"
)

type SignInBody struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func SignIn(c util.ApiContext) error {
	body := SignInBody{}
	err := c.ReadJSON(&body)
	if err != nil {
		return c.Fail(err)
	}
	if body.UserName == "" {
		return c.Fail(errors.New("用户名不能为空"))
	}
	if body.Password == "" {
		return c.Fail(errors.New("密码不能为空"))
	}
	p := model.People{}
	// 查询数据库
	if err = model.GetDB().
		Table("people").
		Where("`username`=? AND `password`=?", body.UserName, body.Password).
		Select("id, locked, lock_why").
		First(&p).Error; err == nil {
		if p.Locked {
			return c.Fail(errors.New("你的账号已被锁定, 锁定原因：" + p.LockWhy))
		}
		// 生成jwt
		tokenStr, err := util.GenerateTheToken(p.ID, "admin")
		if err != nil {
			return c.Fail(err)
		}
		// 写入数据库
		if err = people.UpsertToken(p.ID, tokenStr, 1); err == nil {
			return c.Success(tokenStr)
		}
	}
	if err == gorm.ErrRecordNotFound {
		err = errors.New("用户名或密码错误")
	}
	return c.Fail(err)
}
