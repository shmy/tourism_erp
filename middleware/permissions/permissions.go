package permissions

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/errors"
	"tourism_erp/model"
	"tourism_erp/util"
)
var key = "authUser"

// 检测权限
func Serve (_type int) iris.Handler {
	return func(c iris.Context) {
		cc := util.ApiContext{ c }
		// 获取用户信息
		authUser := cc.Values().Get(key).(*jwt.Token)
		u := authUser.Claims.(jwt.MapClaims)
		t := model.Token{}
		err := model.GetDB().Where("`people_id`=? AND `type`=?", u["id"], _type).First(&t).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				err = errors.New("非法登录")
			}
		} else {
			if authUser.Raw != t.Token {
				err = errors.New("你的账号已在别处登录")
			}
		}
		if err != nil {
			cc.Fail(err)
			cc.StopExecution()
			return
		}
		// TODO 获取权限 并设置
		cc.Values().Set(key, u)
		cc.Next()
	}
}