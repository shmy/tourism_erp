package organization

import (
	"github.com/kataras/iris"
	"tourism_erp/handler/organization"
	"tourism_erp/util"
)

func GetRouting(group iris.Party) {
	group.Get("/organization", util.ApiHandlerWrap(organization.List))
	group.Post("/organization", util.ApiHandlerWrap(organization.Create))
	group.Put("/organization", util.ApiHandlerWrap(organization.Update))
}
