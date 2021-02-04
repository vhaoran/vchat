package ymongo

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/mitchellh/mapstructure"

	"github.com/vhaoran/vchat/lib/yconfig"

	"go.mongodb.org/mongo-driver/bson"
)

var cfg = yconfig.MongoConfig{
	//URL: "mongodb://root:password@172.2.2.2:27017/test?&authSource=admin",
	URL: "mongodb://192.168.0.99:27017/?&slaveOk=true",
	//URL: "mongodb://root:password@192.168.0.99:27017/test?&authSource=admin",

	Options: &yconfig.MongoOptions{
		MaxPoolSize: 100,
		MinPoolSize: 10,
	},
}

type (
	Good struct {
		ID     int    `json:"id,omitempty"`
		Name   string `json:"name,omitempty"`
		Age    int    `json:"age,omitempty"   bson:"age,omitempty"`
		Salary int    `json:"salary,omitempty"   bson:"salary,omitempty"`
	}
)

func Test_mongo_insert_one(t *testing.T) {
	var ctx = context.Background()
	var doc = bson.M{"a": 100, "b": 30}

	client, err := newMongoClient(cfg)
	defer client.Disconnect(ctx)
	if err != nil {
		fmt.Println("------aa-----------")
		log.Println(err)
		return
	}
	log.Println("cnt ok")

	//d
	dbName, tbName := "test", "t"
	tb := client.Database(dbName).Collection(tbName)
	result, err := tb.InsertOne(ctx, doc)
	if err != nil {
		log.Println("insert One err:", err)
	}
	fmt.Println("-----------------")
	log.Println(result)

	//
	c, err := tb.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println("find err:", err)
		return
	}

	//--------result -----------------------------
	for c.Next(ctx) {
		log.Println(c.Current.String())
	}

}
func Test_mongo_insert2(t *testing.T) {
	var ctx = context.Background()
	var doc = bson.M{"a": 100, "b": 30}

	client, err := newMongoClient(cfg)
	defer client.Disconnect(ctx)
	if err != nil {
		fmt.Println("------aa-----------")
		log.Println(err)
		return
	}
	log.Println("cnt ok")

	//d
	dbName, tbName := "test", "t"
	tb := client.Database(dbName).Collection(tbName)

	h := 10000
	t0 := time.Now()
	for i := 0; i < h; i++ {
		doc = bson.M{"a": i, "b": i * 10}
		_, err := tb.InsertOne(ctx, doc)
		if err != nil {
			log.Println("insert One err:", err)
		}
	}
	fmt.Println("-------------time----", time.Since(t0))

	//
	c, err := tb.CountDocuments(ctx, bson.M{})
	if err != nil {
		fmt.Println("count err:", err)
		return
	}
	fmt.Println("------count:----", c)
	//--------result -----------------------------
}

func Test_mongo_insert_batch(t *testing.T) {
	var ctx = context.Background()
	var docs []interface{}

	client, err := newMongoClient(cfg)
	defer client.Disconnect(ctx)
	if err != nil {
		fmt.Println("------aa-----------")
		log.Println(err)
		return
	}
	log.Println("cnt ok")

	//d
	dbName, tbName := "test", "t"
	tb := client.Database(dbName).Collection(tbName)

	h := 1000
	t0 := time.Now()
	for i := 0; i < h; i++ {
		docs = append(docs, bson.M{"a": i, "b": i * 10})

	}

	if _, err := tb.InsertMany(ctx, docs); err != nil {
		log.Println("insert err:", err)
	}
	fmt.Println("-------------time----", time.Since(t0))

	fmt.Println("------count-----------")
	c, err := tb.CountDocuments(ctx, bson.M{})
	if err != nil {
		fmt.Println("count err:", err)
		return
	}
	fmt.Println("------count:----", c)
	//--------result -----------------------------

}

func Test_mongo_struct(t *testing.T) {
	var ctx = context.Background()
	//var doc = bson.M{"a": 100, "b": 30, "c": bson.M{"AA": 1}}
	doc := &Good{
		ID:   1,
		Name: "aaaaa",
	}

	client, err := newMongoClient(cfg)
	defer client.Disconnect(ctx)
	if err != nil {
		fmt.Println("------aa-----------")
		log.Println(err)
		return
	}
	log.Println("cnt ok")

	//d
	dbName, tbName := "test", "t"
	tb := client.Database(dbName).Collection(tbName)
	result, err := tb.InsertOne(ctx, doc)
	if err != nil {
		log.Println("insert One err:", err)
	}
	fmt.Println("-----------------")
	log.Println(result)

	//
	c, err := tb.Find(ctx, Good{})
	if err != nil {
		fmt.Println("find err:", err)
		return
	}

	//--------result -----------------------------
	for c.Next(ctx) {
		log.Println(c.Current.Elements())
	}
}

func Test_m2s(t *testing.T) {
	bean := &Good{
		ID:   1,
		Name: "whr",
	}
	m := make(map[string]interface{})

	err := mapstructure.Decode(bean, &m)
	if err != nil {
		log.Println(err)
		return
	}

	//
	fmt.Println("------aa-----------")
	spew.Dump(m)
}

func Test_test_aggregate(t *testing.T) {

}
