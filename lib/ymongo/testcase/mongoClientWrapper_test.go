package testcase

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/vhaoran/vchat/common/ytime"

	"github.com/davecgh/go-spew/spew"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/vhaoran/vchat/common/ypage"
	"github.com/vhaoran/vchat/lib"
	"github.com/vhaoran/vchat/lib/ylog"
	"github.com/vhaoran/vchat/lib/ymongo"
)

type ABC struct {
	ID int `json:"id,omitempty"`
	//
	Name      string      `json:"name,omitempty"`
	Age       int         `json:"age,omitempty"`
	AgeIsOk   int         `json:"test_b,omitempty"`
	CreatedAt ytime.Date  `json:"created_at,omitempty"   bson:"created_at,omitempty"`
	T         time.Time   `json:"t,omitempty"   bson:"t,omitempty"`
	M         ytime.DateM `json:"m,omitempty"   bson:"m,omitempty"`
	Salary    int         `json:"salary,omitempty"   bson:"salary,omitempty"`
}

var (
	db *ymongo.MongoClientWrapper
)

func init() {
	err := lib.InitModules(
		false,
		false,
		false,
		false,
		true,
	)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	//set time zone
	ytime.SetTimeZone()

	db = ymongo.X
	if db == nil {
		log.Println("----------", "init", "------------")
		panic("db not prepare")
	} else {
		log.Println("----------", "db cnt passed", "------------")
	}
}

func Test_wrapper_insert_one(t *testing.T) {

	bean := &ABC{
		ID:        5,
		Name:      "whr",
		Age:       3,
		CreatedAt: ytime.OfNow(),
		T:         time.Now(),
		M:         ytime.OfTimeM(time.Now().AddDate(0, 0, -5)),
		Salary:    50,
	}

	ret, err := db.DoInsertOne("test", "abc", bean)
	log.Println("--------", "", "--------------")
	log.Println(ret)
	log.Println("----------------------")
	log.Println(err)
}

func Test_wrapper_insert_many(t *testing.T) {
	l := make([]*ABC, 0)
	for i := 0; i < 1234; i++ {
		bean := &ABC{
			ID:   i,
			Name: fmt.Sprint("wrh_", i),
			Age:  i,
		}
		l = append(l, bean)
	}

	ret, err := db.DoInsertMany("test", "t", l)
	log.Println("--------", "", "--------------")
	log.Println(ret)
	log.Println("----------------------")
	log.Println(err)
	//

}

func Test_mongo_client_wrapper_find(t *testing.T) {
	client := db.Base
	ctx := context.Background()

	dbName, tbName := "test", "abc"
	tb := client.Database(dbName).Collection(tbName)
	c, err := tb.Find(ctx, bson.M{"id": 1})
	log.Println("----------", err, "------------")
	for c.Next(ctx) {
		log.Println(c.Current)
		bean := &ABC{}
		if er1 := bson.Unmarshal(c.Current, &bean); er1 != nil {
			log.Println(err)
		} else {
			log.Println("----------", "unmarshal ok", "------------")
			spew.Dump(bean)
		}
	}

}

func Test_test_wrapper_find(t *testing.T) {
	l := make([]*ABC, 0)
	err := db.DoFindMany(&l, "test",
		"abc",
		bson.M{"_id": 1, "b": 10})
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("------", "", "-----------")
	spew.Dump(l)
}

func Test_mongo_update_one(t *testing.T) {
	//var id primitive.ObjectID

	opts := options.FindOneAndUpdate().SetUpsert(true)
	filter := bson.D{bson.E{Key: "a", Value: 1}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "b", Value: 333}}}}
	var updatedDocument bson.M

	client := db.Base
	tb := client.Database("test").Collection("t")
	err := tb.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updatedDocument)

	if err != nil {
		log.Println("########## err:", err)
		return
	}

	fmt.Printf("updated document %v", updatedDocument)
}

func Test_mongo_update_many(t *testing.T) {
	//var id primitive.ObjectID

	//opts := options.FindOneAndUpdate().SetUpsert(true)
	filter := bson.D{{"a", 1}}
	update := bson.D{{"$set", bson.D{{"b", 100}}}}
	var updatedDocument bson.M

	client := db.Base
	tb := client.Database("test").Collection("abc")
	//err := tb.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updatedDocument)
	_, err := tb.UpdateMany(context.TODO(), filter, update)

	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			log.Println("########## err:", err)
			return
		}
		log.Fatal(err)
	}

	fmt.Printf("updated document %v", updatedDocument)
}

func Test_wrap_update_one(t *testing.T) {
	err := db.DoUpdateOne("test", "t",
		bson.D{{"a", 1}},
		bson.D{{"b", 32}},
	)
	if err != nil {
		ylog.Error("mongoClientWrapper_test.go->", err)
		return
	}
	fmt.Println("------", "ok", "-----------")
}

func Test_wrap_update_many(t *testing.T) {
	err := db.DoUpdateMany("test", "t",
		bson.D{{"a", 1}},
		bson.D{{"b", 42}})

	if err != nil {
		ylog.Error("mongoClientWrapper_test.go->", err)
		return
	}
	fmt.Println("------", "ok", "-----------")
}

func Test_del_many(t *testing.T) {
	i, err := db.DoDelMany("test", "t",
		bson.D{{"a", 1}})
	log.Println("----------", "aaa", "------------")
	ylog.Debug("i", i, "  err", err)
}

func Test_del_one(t *testing.T) {
	i, err := db.DoDelOne("test", "t",
		bson.D{{"a", ""}})
	log.Println("----------", "aaa", "------------")
	ylog.Debug("i", i, "  err", err)
}

type T struct {
	Key   string
	Value string
}

func Test_a_test(t *testing.T) {
	c := []T{{"a", "b"}, {"c", "d"}}
	for _, v := range c {
		fmt.Println(v)
	}
}

func Test_a_insert_One(t *testing.T) {
	bean := &ABC{
		ID:        1,
		Name:      "abc",
		Age:       0,
		AgeIsOk:   0,
		CreatedAt: ytime.OfNow(),
		Salary:    50,
	}
	_, err := db.DoInsertOne("test", "abc", bean)
	if err != nil {
		ylog.Error("mongoClientWrapper_test.go->", err)
		return
	}
}

func Test_a_b_findOne(t *testing.T) {
	bean := &ABC{

	}

	err := db.DoFindOne(bean, "test", "abc",
		bson.M{"name": "abc"})
	if err != nil {
		ylog.Error("### mongoClientWrapper_test.go->", err)
		return
	}
	//
	log.Println("----------", "aaa", "------------")
	ylog.DebugDump(bean)
}

func Test_count(t *testing.T) {
	var opts *options.CountOptions

	i, err := db.Table("test", "abc").CountDocuments(
		context.Background(),
		bson.D{{"name", "abc"}}, opts)
	if err != nil {
		ylog.Error("####mongoClientWrapper_test.go->", err)
		return
	}
	log.Println("---------count: ", i, "------------")
}

func Test_count_1(t *testing.T) {
	i, err := db.DoCount("test", "abc",
		bson.D{{"name", "abc"}})
	if err != nil {
		ylog.Error("####mongoClientWrapper_test.go->", err)
		return
	}
	log.Println("---------count: ", i, "------------")
}

func Test_a_b_find_many(t *testing.T) {
	l := make([]*ABC, 0)

	i := int64(5)
	opts := &options.FindOptions{
		Limit: &i,
		Sort: bson.D{{
			"Name",
			1,
		}},
	}

	err := db.DoFindMany(&l, "test", "abc",
		bson.M{"name": bson.M{"$ne": "abc"}}, opts)
	if err != nil {
		ylog.Error("### mongoClientWrapper_test.go->", err)
		return
	}

	//
	log.Println("----------", "aaa", "------------")
	ylog.DebugDump(l)
}

func Test_page_bean(t *testing.T) {
	bean := &ypage.PageBean{
		PageNo:      2,
		RowsPerPage: 10,
		PagesCount:  0,
		RowsCount:   0,
		Where:       bson.D{{}},
		Sort:        nil, //bson.D{{"name", 1}},
		Data:        nil,
	}

	l := make([]*ABC, 0)

	err := db.DoPage(&l, "test", "abc", bean)

	if err != nil {
		ylog.Error("mongoClientWrapper_test.go->", err)
		return
	}
	//
	fmt.Println("------", "", "-----------")
	ylog.DebugDump("", bean)
}

func Test_find_many_1(t *testing.T) {
	l := make([]*ABC, 0)
	//err := db.DoFindMany(&l, "test",
	//	"abc",
	//	bson.D{{"_id", 1}, {"b", 10}})
	//err := db.DoFindMany(&l, "test",
	//	"abc",
	//	bson.D{{"name", "whr"},
	//		{"age", bson.M{"$gte": 3}}})

	//err := db.DoFindMany(&l, "test",
	//	"abc",
	//	bson.M{"name": "whr",
	//		"created_at.time": bson.M{"$lt": ytime.OfNow())}})

	err := db.DoFindMany(&l, "test",
		"abc",
		bson.M{"name": "whr",
			"created_at.time":
			bson.M{"$lt": ytime.OfNow().TimeShanghai()}})

	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("------", "", "-----------")
	spew.Dump(l)
	log.Println("----------", "json", "------------")
	s, _ := json.Marshal(l)
	log.Println("----------", "s", string(s), "------------")

}

func Test_find_many_of_map(t *testing.T) {
	l := make([]*ABC, 0)

	limit := int64(2)
	skip := int64(1)
	opts := &options.FindOptions{
		Limit: &limit,
		Skip:  &skip,
		Sort: bson.M{
			"age": 1,
		},
	}

	err := db.DoFindMany(&l, "test",
		"abc",
		bson.M{"name": bson.M{
			"$ne": "whr",
		},
		}, opts)

	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("------", "", "-----------")
	spew.Dump(l)
	log.Println("----------", "json", "------------")
	s, _ := json.Marshal(l)
	log.Println("----------", "s", string(s), "------------")

}

func Test_PageBean_map(t *testing.T) {
	pb := &ypage.PageBeanMap{
		PageNo:      1,
		RowsPerPage: 2,
		Where: bson.M{
			"name": bson.M{
				"$ne": "whr",
			},
		},
		Sort: bson.M{
			"name": 1,
			"age":  1,
		},
	}

	l := make([]*ABC, 0)
	bean, err := ymongo.X.DoPageMap(&l, "test", "abc", pb)
	if err != nil {
		ylog.Error("mongoClientWrapper_test.go->", err)
		return
	}
	log.Println("-----aaa-----", "ok", "------------")
	spew.Dump(bean)
}

func Test_PageBean_map_2(t *testing.T) {
	pb := &ypage.PageBeanMap{
		PageNo:      1,
		RowsPerPage: 5,
		Where: bson.M{
			"name": bson.M{
				"$ne": "whr",
			},
			"$or": []bson.M{
				{
					"age": 70,
				},
				{
					"age": 80,
				},
			},
		},
		Sort: bson.M{
			"name": 1,
			"age":  1,
		},
	}

	l := make([]*ABC, 0)
	bean, err := ymongo.X.DoPageMap(&l, "test", "abc", pb)
	if err != nil {
		ylog.Error("mongoClientWrapper_test.go->", err)
		return
	}
	log.Println("-----aaa-----", "ok", "------------")
	spew.Dump(bean)
}
