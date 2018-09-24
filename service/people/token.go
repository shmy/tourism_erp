package people

import (
	"github.com/jinzhu/gorm"
	"tourism_erp/model"
)

func UpsertToken (id uint, token string, _type int ) error {
	t := model.Token{}
	err := model.GetDB().
		Table("token").
		Where("`people_id`=? AND `type`=?", id, _type).
		First(&t).Error
	// 存在则更新
	if err == nil {
		t.Token = token
		err = model.GetDB().Save(&t).Error
	}
	// 不存在则新增
	if err == gorm.ErrRecordNotFound {
		t.Token = token
		t.PeopleId = id
		t.Type = _type
		err = model.GetDB().Create(&t).Error
	}
	if err != nil {
		return err
	}
	return nil
}