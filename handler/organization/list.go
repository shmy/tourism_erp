package organization

import (
	"tourism_erp/model"
	"tourism_erp/util"
)

func List (c util.ApiContext) error {
	results := make([]model.Organization, 0)
	err := model.GetDB().Preload("Department").Find(&results).Error
	//err := model.GetDB().Select("`id`, `name`, `describe`, `pid`, `created_at`").Find(&results).Error
	if err != nil {
		return c.Fail(err)
	}
	return c.Success(results)
}