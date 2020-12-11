package ypg

import (
	"github.com/jinzhu/gorm"

	"github.com/vhaoran/vchat/lib/ylog"
)

func Tx(callback func(*gorm.DB) error) error {
	tx := X.Begin()
	defer func() {
		if err := recover(); err != nil {
			ylog.Error("tx-util.go->", err)
			tx.Rollback()
		}
	}()

	err := callback(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
