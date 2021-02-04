package ypage

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"testing"
)

func Test_sql_where_1(t *testing.T) {
	obj := new(SqlWhere)
	m := bson.D{
		{"id", 1},
		{"user_name", "whr"},
		{"$in", []int{1}},
		{"$nin", []int{1, 2, 3}},
	}

	s, p := obj.GetWhere(m)
	log.Println("----------", s, "------------")
	spew.Dump("p:", p)
}

func Test_sql_where_2(t *testing.T) {
	obj := new(SqlWhere)
	m := bson.D{
		{"id", bson.E{"$gte", 3}},
		{"user_name", "whr"},
	}

	s, p := obj.GetWhere(m)
	log.Println("----------", s, "------------")
	spew.Dump("p:", p)
}

func Test_sql_where_21(t *testing.T) {
	obj := new(SqlWhere)
	m := bson.D{
		{"id", bson.M{"$gte": 3}},
		{"id", bson.M{"$in": []int{1, 2, 3, 4}}},
		{"user_name", "whr"},
	}

	s, p := obj.GetWhere(m)
	log.Println("----------", s, "------------")
	spew.Dump("p:", p)
}

func TestSqlWhere_IsMap(t *testing.T) {
	obj := new(SqlWhere)
	v := bson.M{
		"a": "b",
	}
	log.Println("-----bson.M-----", obj.IsMap(v), "------------")
	//
	v1 := []bson.E{bson.E{
		Key: "a", Value: "b",
	},
	}

	log.Println("-----bson.M-----", obj.IsMap(v1), "------------")
	//

}

func Test_map2_e(t *testing.T) {
	obj := new(SqlWhere)
	v := bson.M{
		"a": "a_value",
		"b": "b_value",
		"c": "c_value",
	}
	bean := obj.Map2E(v)
	log.Println("-----bson.M-----", "------------")
	spew.Dump("bean:", bean)
	//
}

func Test_slice_d(t *testing.T) {
	obj := new(SqlWhere)
	v := []bson.M{
		bson.M{
			"a": "a_value",
			"b": "b_value",
			"c": "c_value",
		},
		bson.M{
			"a": "a_value)1",
			"b": "b_value)1",
			"c": "c_value)1",
		},
	}

	bean := obj.Slice2D(v)
	log.Println("-----bson.M-----", "------------")
	spew.Dump("bean:", bean)
	//
}

func Test_GetWhere_map(t *testing.T) {
	//
	sql, p := new(SqlWhereMap).GetWhere(bson.M{
		"id":   1,
		" 1 = 1 \";and \"code": "23",
		"$in":  []int{},
		"x":    bson.M{"$nin": []int{1, 2, 3}},
	})

	spew.Dump(p)
	fmt.Println("--------------")
	spew.Dump(sql)

}
