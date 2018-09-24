package people

import (
	"github.com/kataras/iris"
	"tourism_erp/handler/people"
	"tourism_erp/util"
)

func GetRouting (group iris.Party) {
	group.Get("/people", util.ApiHandlerWrap(people.List))
}
