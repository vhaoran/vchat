package xorm

import (
	"fmt"
	"log"
	"testing"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "root"
	password = "password"
	dbName   = "test"
)

func Test_x_orm_pg(t *testing.T) {
	user := &UserTbl{
		Id:       1,
		Username: "Windows",
		Sex:      1,
		Info:     "操作系统",
	}

	SessionUserTest(user)
}

func GetCnt() *xorm.Engine {
	url := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	//格式
	cnt, err := xorm.NewEngine("postgres", url)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	cnt.ShowSQL() //菜鸟必备

	err = cnt.Ping()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	fmt.Println("connect postgresql success")
	return cnt
}

//table name 为user_tbl
type UserTbl struct {
	Id       int
	Username string
	Sex      int
	Info     string
}

//查询所有
func selectAll() {
	var user []UserTbl
	engine := GetCnt()
	engine.SQL("select * from user_tbl where ").Find(&user)
	fmt.Println(user)
}

//条件查询
func selectUser(name string) {
	var user []UserTbl
	engine := GetCnt()
	engine.Where("user_tbl.username=?", name).Find(&user)
	fmt.Println(user)
}

//可以用Get查询单个元素
func selectOne(id int) {
	var user UserTbl
	engine := GetCnt()
	engine.Id(id).Get(&user)
	//engine.Alias("u").Where("u.id=?",id).Get(&user)
	fmt.Println(user)
}

//添加
func InsertUser(user *UserTbl) bool {
	engine := GetCnt()
	rows, err := engine.Insert(user)
	if err != nil {
		log.Println(err)
		return false
	}
	if rows == 0 {
		return false
	}
	return true
}

//删除(根据名称删除)
func DeleteUser(name string) bool {
	user := UserTbl{
		Username: name,
	}
	engine := GetCnt()
	rows, err := engine.Delete(&user)
	if err != nil {
		log.Println(err)
		return false
	}
	if rows == 0 {
		return false
	}
	return true
}

//利用sql删除
func DeleteUserBySQL(name string) bool {
	engine := GetCnt()
	result, err := engine.Exec("delete from user_tbl where username=?", name)
	if err != nil {
		log.Println(err)
		return false
	}
	rows, err := result.RowsAffected()
	if err == nil && rows > 0 {
		return true
	}
	return false
}

//更新
func UpdateUser(user *UserTbl) bool {
	engine := GetCnt()
	//Update(bean interface{}, condiBeans ...interface{}) bean是需要更新的bean,condiBeans是条件
	rows, err := engine.Update(user, UserTbl{Id: user.Id})
	if err != nil {
		log.Println(err)
		return false
	}
	if rows > 0 {
		return true
	}
	return false
}

//利用session进行增删改
//用session的好处就是可以事务处理
func SessionUserTest(user *UserTbl) {
	engine := GetCnt()
	if er1 := engine.Sync(new(UserTbl)); er1 != nil {
		fmt.Println(er1)
		return
	}

	session := engine.NewSession()
	session.Begin()
	_, err := session.Insert(user)
	if err != nil {
		session.Rollback()
		log.Fatal(err)
	}

	user.Username = "windows"
	_, err = session.Update(user, UserTbl{Id: user.Id})
	if err != nil {
		session.Rollback()
		log.Fatal(err)
	}

	_, err = session.Delete(user)
	if err != nil {
		session.Rollback()
		log.Fatal(err)
	}

	err = session.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
