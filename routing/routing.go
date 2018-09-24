package routing

import (
	"github.com/kataras/iris"
	"tourism_erp/routing/admin"
	"tourism_erp/routing/svcd"
)

func RouteMapping(app *iris.Application) {
	// 静态服务
	app.StaticWeb("/static", "./public")
	// 系统监控
	svcdGroup := app.Party("/svcd")
	svcd.GetRouting(svcdGroup)

	// api admin v1
	apiV1Party := app.Party("/api/v1/admin")
	{
		admin.GetRouting(apiV1Party)
	}
}
