package ypg

import (
	"github.com/vhaoran/vchat/lib/ylog"
	"log"

	gorm "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/vhaoran/vchat/lib/yconfig"
)

var (
	X *gorm.DB
)

/*--auth: whr  date:2019/12/511:44--------------------------
 ####请勿擅改此功能代码####
 用途：连接到数据库功能
 --->：yconfig
--------------------------------------- */

func InitPG(cfg yconfig.PGConfig, debug ...bool) (err error) {
	if cnt, err := NewPGCnt(&cfg, debug...); err != nil {
		return err
	} else {
		cnt.Callback().Create().Replace("gorm:update_time_stamp", createCallback)
		cnt.Callback().Update().Replace("gorm:update_time_stamp", updateCallback)
		//cnt.Callback().Delete().Replace("gorm:delete", deleteCallback)
		X = cnt
		return nil
	}
}

func NewPGCnt(cfg *yconfig.PGConfig, debug ...bool) (*gorm.DB, error) {
	connStr := cfg.URL

	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	db.DB().SetMaxOpenConns(cfg.PoolMax)
	db.DB().SetMaxIdleConns(cfg.PoolMax)

	if len(debug) > 0 {
		db.LogMode(true)
		db.SetLogger(ylog.GetLogger())
	}

	return db, nil
}
