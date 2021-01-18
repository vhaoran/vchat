package ypg

import (
	"errors"

	"github.com/jinzhu/gorm"
)

func DBErr(db *gorm.DB) error {
	if db.Error != nil {
		return tranErr(db.Error)
	}

	if db.RowsAffected <= 0 {
		return errors.New("操作失败，可能是以下原因【未找到符合条件的数据】")
	}
	return nil
}

func tranErr(err error) error {
	if err == gorm.ErrRecordNotFound {
		return errors.New("记录没有发现")
	}
	if err == gorm.ErrInvalidSQL {
		return errors.New("无效的sql(ErrInvalidSQL)")
	}
	if err == gorm.ErrInvalidTransaction {
		return errors.New("无效的事务处理(ErrInvalidTransaction)")
	}
	if err == gorm.ErrCantStartTransaction {
		return errors.New("无法开始事务(ErrCantStartTransaction)")
	}
	if err == gorm.ErrUnaddressable {
		return errors.New("非法指针(ErrUnaddressable)")
	}
	return err
}

func DBFindErr(db *gorm.DB) error {
	if db.Error != nil {
		return tranErr(db.Error)
	}
	//if db.RowsAffected <= 0 {
	//	return errors.New("操作失败，可能是以下原因【未找到符合条件的数据】")
	//}
	return nil
}
