package ypg

import (
	"errors"

	"github.com/jinzhu/gorm"
)

func DBErr(db *gorm.DB) error {
	if db.Error != nil {
		return db.Error
	}

	if db.RowsAffected <= 0 {
		return errors.New("操作失败，可能是以下原因【未找到符合条件的数据】")
	}
	return nil
}

func DBFindErr(db *gorm.DB) error {
	if db.Error != nil {
		return db.Error
	}
	//if db.RowsAffected <= 0 {
	//	return errors.New("操作失败，可能是以下原因【未找到符合条件的数据】")
	//}
	return nil
}
