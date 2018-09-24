package department

import (
	"github.com/kataras/iris"
	"tourism_erp/handler/department"
	"tourism_erp/util"
)

func GetRouting (group iris.Party) {
	group.Get("/department", util.ApiHandlerWrap(department.List))
}