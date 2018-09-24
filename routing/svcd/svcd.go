package svcd

import (
	"github.com/kataras/iris"
	"tourism_erp/handler/svcd"
)

func GetRouting(group iris.Party) {
	group.Get("/health", svcd.HealthCheck)
	group.Get("/disk", svcd.DiskCheck)
	group.Get("/cpu", svcd.CPUCheck)
	group.Get("/ram", svcd.RAMCheck)
}
