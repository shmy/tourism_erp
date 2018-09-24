package admin

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	"github.com/spf13/viper"
	peopleHandler "tourism_erp/handler/people"
	jwtMiddleware "tourism_erp/middleware/jwt"
	"tourism_erp/middleware/permissions"
	"tourism_erp/routing/admin/department"
	"tourism_erp/routing/admin/organization"
	"tourism_erp/routing/admin/people"
	"tourism_erp/util"
)

func GetRouting(group iris.Party) {
	var secret = viper.GetString("jsonwebtoken.admin.secret")
	jwtHandler := jwtMiddleware.New(jwtMiddleware.Config{
		ContextKey: "authUser",
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
		ErrorHandler: func(c iris.Context, s string) {
			// jwt验证失败调用
			c.JSON(map[string]interface{}{
				"code": 1,
				"msg":  s,
				"data": nil,
			})
		},
	})
	group.Post("/sign_in", util.ApiHandlerWrap(peopleHandler.SignIn))
	// Token 有效验证
	group.Use(jwtHandler.Serve)
	// Token 持久化和权限验证
	group.Use(permissions.Serve(1))
	organization.GetRouting(group)
	department.GetRouting(group)
	people.GetRouting(group)
}
