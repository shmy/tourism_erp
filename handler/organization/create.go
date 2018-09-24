package organization

import (
	"tourism_erp/config"
	"tourism_erp/util"
)

func Create(c util.ApiContext) error {
	// 重新从数据库加载规则
	config.GetEnforcer().LoadPolicy()
	return c.Success("post")
}
