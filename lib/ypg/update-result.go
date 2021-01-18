package ypg

import (
	"errors"

	"github.com/jinzhu/gorm"
)

func ResultOfOrm(db *gorm.DB) error {
	i := db.RowsAffected
	err := db.Error
	if i > 0 {
		return nil
	}
	if err != nil {
		return tranErr(err)
	}
	return errors.New("没有找到需要处理的内容")
}
