package department

import (
	"tourism_erp/model"
	"tourism_erp/util"
)

func List(c util.ApiContext) error {
	results := make([]model.Department, 0)
	err := model.GetDB().Preload("Organization").Find(&results).Error
	//err := model.GetDB().Preload("Organization").Select("`id`, `name`, `describe`, `pid`, `created_at`").Find(&results).Error
	if err != nil {
		return c.Fail(err)
	}
	return c.Success(results)
}
