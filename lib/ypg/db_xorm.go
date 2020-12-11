package ypg

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"github.com/vhaoran/vchat/lib/yconfig"
	"log"
	_ "xorm.io/core"

	_ "github.com/go-xorm/xorm"
)

var (
	xdb *xorm.Engine
)

func InitPGOfXOrm_NoUsed(cfg yconfig.PGConfig) (err error) {
	if cnt, err := NewPGCntOfXOrm(cfg); err != nil {
		return err
	} else {
		//cnt.Callback().Delete().Replace("gorm:delete", deleteCallback)
		xdb = cnt
		return nil
	}
}

func NewPGCntOfXOrm(cfg yconfig.PGConfig) (*xorm.Engine, error) {
	//url := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	url := cfg.URL
	cnt, err := xorm.NewEngine("postgres", url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	cnt.SetMaxOpenConns(cfg.PoolMax)
	cnt.SetMaxIdleConns(cfg.PoolMin)
	cnt.ShowSQL() //菜鸟必备

	err = cnt.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	log.Println("connect postgres success")
	return cnt, nil
}
