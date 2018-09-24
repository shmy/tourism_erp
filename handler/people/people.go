package people


import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"tourism_erp/model"
	"tourism_erp/util"
)

func List (c util.ApiContext) error {
	// 获取用户信息
	authUser := c.Values().Get("authUser").(jwt.MapClaims)
	fmt.Println(authUser["id"])
	results := make([]model.People, 0)
	err := model.GetDB().Find(&results).Error
	if err != nil {
		return c.Fail(err)
	}
	return c.Success(results)
}