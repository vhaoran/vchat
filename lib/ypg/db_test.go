package ypg

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	gorm "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/magiconair/properties/assert"
	"github.com/vhaoran/vchat/common/ytime"
	"log"
	"testing"

	"github.com/vhaoran/vchat/lib/yconfig"
)

type (
	personX struct {
		BaseModel
		ID       uint32 `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
		Name     string `gorm:"size:255;not null;unique" json:"Name"`
		Nickname string `gorm:"size:255" json:"nickname"`
		Email    string `gorm:"size:255" json:"email"`
		Password string `gorm:"size:255" json:"password"`
	}
)

func (personX) TableName() string {
	return "person"
}

func newDBCnt() *gorm.DB {
	connStr := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		"127.0.0.1", //viper.GetString("DB_HOST"),
		"root",      //viper.GetString("DB_USER"),
		"test",      //viper.GetString("DB_NAME"),
		"password",  ///viper.GetString("DB_PASS"),
	)

	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		log.Println(err)
	}
	db.DB().SetMaxOpenConns(500)
	db.LogMode(true)
	return db
}

func Test_cnt(t *testing.T) {
	bean := &personX{}

	db := newDBCnt()

	if db.HasTable(bean) {
		if err := db.DropTable(bean).Error; err != nil {
			log.Println("err:", err)
			return
		}
	}
	log.Println("after cnt ")

	db.CreateTable(bean)
	db.AutoMigrate(bean)

	log.Println("create table okt ")
}

func Test_cnt_insert(t *testing.T) {
	db := newDBCnt()

	ytime.SetTimeZone()

	h := 100
	for i := 0; i < h; i++ {
		func(k int) {
			bean := &personX{
				BaseModel: BaseModel{
					CreatedAt: ytime.OfNow(),
					UpdatedAt: ytime.OfNow(),
				},
				Name:     fmt.Sprint("aaaaa", k, "bbb", k),
				Nickname: "",
				Email:    "",
				Password: "",
			}

			if err := db.Save(bean).Error; err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("append ok")
			}
		}(i)
	}


}

var db *gorm.DB

func prepareCnt() error {
	cfg := &yconfig.PGConfig{
		URL:     "host=127.0.0.1 user=root dbname=test sslmode=disable password=password",
		PoolMax: 100,
		PoolMin: 19,
	}
	cnt, err := NewPGCnt(cfg)
	if err != nil {
		log.Println(err)
		return err
	}
	db = cnt
	return nil
}

func Test_db_cnt_config(t *testing.T) {
	err := prepareCnt()
	if err != nil {
		log.Println("cnt fail:", err)
	}

	//
	l := make([]*personX, 0)
	if err = db.Find(&l).Error; err != nil {
		log.Println(err)
	}
	spew.Dump(l)

	//condiation find
	bean := &personX{ID: 2}
	l = make([]*personX, 0)
	db.Find(&l, )
	assert.Equal(t, bean.ID, 2, "found!")
	spew.Dump(bean)

	//raw find
	l = make([]*personX, 0)
	if err = db.Raw("select * from person where id > ? limit 2", 50).Find(&l).Error; err != nil {
		log.Println("raw sql err:", err)
	}
	//
	assert.Equal(t, len(l) > 0, true, "raw sql passed")
	fmt.Println("-----------------")
	spew.Dump(l)

	//-------- -----------------------------
	// db.Model(&user).Updates(User{Name: "hello", Age: 18})
}

func Test_db_time(t *testing.T) {
	err := prepareCnt()
	if err != nil {
		log.Println(err)
		return
	}

	//db.Delete(personX{})

	for i := 0; i < 10; i++ {
		bean := &personX{
			Name: fmt.Sprint(i, "_name"),
		}

		if err = db.Save(bean).Error; err != nil {
			log.Println(err)
			return
		}
		log.Println("ok->", i)
	}
}

func Test_db_time_marshal(t *testing.T) {
	err := prepareCnt()
	if err != nil {
		log.Println(err)
		return
	}

	l := make([]personX, 0)
	if err = db.Find(&l).Error; err != nil {
		log.Println(err)
		return
	}

	//-------------------------------------
	buffer, err := json.Marshal(l)
	if err != nil {
		log.Println(err)
		return
	}
	spew.Dump(string(buffer))

	fmt.Println("------unmarshal-----------")
	l = make([]personX, 0)
	if err = json.Unmarshal(buffer, &l); err != nil {
		log.Println(err)
		return
	}
	spew.Dump(l)
}
