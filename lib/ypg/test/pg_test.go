package test

import (
	"fmt"
	"log"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/vhaoran/vchat/common/ypage"
	"github.com/vhaoran/vchat/common/ytime"
	"github.com/vhaoran/vchat/lib"
	"github.com/vhaoran/vchat/lib/ypg"
)

type GoodA struct {
	ID   int64
	Name string
	T    ytime.Date
	TM   ytime.Date
	B    ytime.DateM
}

func (GoodA) TableName() string {
	return "good_a"
}

func (r *GoodA) AfterSave(scope *gorm.Scope) (err error) {
	//bean := scope.Value.(*GoodA)
	log.Println("....after save...")
	spew.Dump(r)

	return nil
}

func init() {
	_, err := lib.InitModulesOfAll()
	if err != nil {
		panic(err.Error())
	}
	ypg.X.AutoMigrate(new(GoodA))
}

func Test_trigger(t *testing.T) {
	//
	bean := &GoodA{
		Name: "abc",
	}
	err := ypg.X.Save(bean).Error
	if err != nil {
		log.Println(err)
	}
}

func Test_trigger_2(t *testing.T) {
	err := ypg.X.Model(new(GoodA)).Where("id=?", 1).Update("name", "abcde").Error
	if err != nil {
		log.Println(err)
	}
}

func Test_insert(t *testing.T) {
	ypg.X.AutoMigrate(new(GoodA))
	ypg.X.LogMode(true)
	ytime.SetTimeZone()

	for i := 3; i < 100; i++ {
		//c := ytime.OfNowM()
		bean := &GoodA{
			Name: fmt.Sprint(i, " name_"),
			T:    ytime.OfNow(),
			TM:   ytime.OfNow(),
			//B:    ytime.OfNowM(),
		}
		err := ypg.X.Save(bean).Error
		if err != nil {
			fmt.Println("pg_test.go->", err)
			return
		}
		fmt.Println(i, " ok")
	}
}

func Test_page(t *testing.T) {
	ypg.X.AutoMigrate(new(GoodA))
	ypg.X.LogMode(true)
	ytime.SetTimeZone()

	where := bson.D{{
		"id", bson.M{"$gte": 1},
	}}
	sort := bson.D{
		{"id", 1},
		{"name", 1}}

	exp, p := new(ypage.SqlWhere).GetWhere(where)

	//
	l := make([]*GoodA, 0)
	err := ypg.X.Order(ypage.GetSort(sort)).Limit(3).Offset(5).Where(exp, p...).Find(&l).Error
	fmt.Println(err)
	log.Println("----------", "ok", "------------")
	spew.Dump(l)

	log.Println("----------", "aaa", "------------")
	//
	count := 0
	err = ypg.X.Model(new(GoodA)).Where(exp, p...).Count(&count).Error
	fmt.Println("count:", count)

}

func Test_page1(t *testing.T) {
	ypg.X.AutoMigrate(new(GoodA))
	ypg.X.LogMode(true)
	ytime.SetTimeZone()

	where := bson.D{{
		"id", bson.M{"$gte": 1},
	}}
	sort := bson.D{
		{"id", 1},
		{"name", 1}}
	bean := &ypage.PageBean{
		PageNo:      1,
		RowsPerPage: 2,
		Where:       where,
		Sort:        sort,
	}

	l := make([]*GoodA, 0)
	ret, err := ypg.Page(ypg.X, &l, bean, "name")
	log.Println("----------", err, "------------")
	spew.Dump(ret)
}

func Test_page_map_1(t *testing.T) {
	ypg.X.AutoMigrate(new(GoodA))
	ypg.X.LogMode(true)
	ytime.SetTimeZone()

	where := bson.M{
		"id": bson.M{"$gte": 1},
	}
	sort := bson.M{
		"id":   1,
		"name": 1}
	bean := &ypage.PageBeanMap{
		PageNo:      1,
		RowsPerPage: 2,
		Where:       where,
		Sort:        sort,
	}

	l := make([]*GoodA, 0)
	ret, err := ypg.PageMap(ypg.X, &l, bean)
	log.Println("----------", err, "------------")
	spew.Dump(ret)
}
