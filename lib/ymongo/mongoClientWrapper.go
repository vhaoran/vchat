package ymongo

import (
	"context"
	"errors"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/vhaoran/vchat/common/reflectUtils"
	"github.com/vhaoran/vchat/common/ypage"
)

type (
	MongoClientWrapper struct {
		Base *mongo.Client
	}
)

func (r *MongoClientWrapper) Table(dbName, tbName string) *mongo.Collection {
	client := r.Base
	return client.Database(dbName).Collection(tbName)
}

func (r *MongoClientWrapper) DoInsertOne(dbName, tbName string,
	doc interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	var ctx = context.Background()

	client := r.Base
	//log.Println("cnt ok")

	//var doc = bson.M{"a": 100, "b": 30}
	//d
	//dbName, tbName := "test", "t"
	tb := client.Database(dbName).Collection(tbName)
	return tb.InsertOne(ctx, doc, opts...)
}

func (r *MongoClientWrapper) tran2Slice(a interface{}) ([]interface{}, error) {
	l := make([]interface{}, 0)

	v := reflect.Indirect(reflect.ValueOf(a))
	if !v.IsValid() {
		return nil, errors.New("wrong type,only support array/slice/struct/pointer of struct")
	}

	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		{
			for i := 0; i < v.Len(); i++ {
				l = append(l, v.Index(i).Interface())
			}
			return l, nil
		}
	case reflect.Struct:
		{
			l = append(l, v.Interface())
			return l, nil
		}
	}

	return nil, errors.New("data is empty")
}

func (r *MongoClientWrapper) DoInsertMany(dbName, tbName string,
	doc interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	var ctx = context.Background()
	db := r.Base

	l, err := r.tran2Slice(doc)
	if err != nil {
		return nil, err
	}

	//var doc = bson.M{"a": 100, "b": 30}
	//d
	//dbName, tbName := "test", "t"
	tb := db.Database(dbName).Collection(tbName)
	return tb.InsertMany(ctx, l, opts...)
}

func (r *MongoClientWrapper) DoUpdateOne(dbName, tbName string,
	filter, updateExp bson.D) error {
	//ctx := context.Background()
	ctx := context.TODO()

	db := r.Base
	tb := db.Database(dbName).Collection(tbName)
	//记录不存在时，不新增
	opts := options.FindOneAndUpdate().SetUpsert(false)

	//filter := bson.D{{"a", 1}}
	update := bson.D{{"$set", updateExp}}
	var ret bson.M

	err := tb.FindOneAndUpdate(ctx, filter, update, opts).Decode(&ret)
	if err != nil {
		return err
	}

	return nil
}

func (r *MongoClientWrapper) DoDelMany(dbName, tbName string,
	filter bson.D) (delCount int64, err error) {
	db := r.Base
	tb := db.Database(dbName).Collection(tbName)
	ctx := context.TODO()

	opts := options.Delete().SetCollation(&options.Collation{
		//Locale:    "en_US",
		Locale:    "zh",
		Strength:  1,
		CaseLevel: false,
	})
	ret, err := tb.DeleteMany(ctx, filter, opts)
	if err != nil {
		return
	}
	delCount = ret.DeletedCount
	return
}
func (r *MongoClientWrapper) DoDelOne(dbName, tbName string,
	filter bson.D) (delCount int64, err error) {
	db := r.Base
	tb := db.Database(dbName).Collection(tbName)
	ctx := context.TODO()

	opts := options.Delete().SetCollation(&options.Collation{
		//Locale:    "en_US",
		Locale:    "zh",
		Strength:  1,
		CaseLevel: false,
	})

	ret, err := tb.DeleteOne(ctx, filter, opts)
	if err != nil {
		return
	}
	delCount = ret.DeletedCount
	return
}

func (r *MongoClientWrapper) DoUpdateMany(dbName, tbName string,
	filter, updateExp bson.D) error {
	//ctx := context.Background()
	ctx := context.TODO()

	db := r.Base
	tb := db.Database(dbName).Collection(tbName)
	update := bson.D{{"$set", updateExp}}

	//err := tb.FindOneAndUpdate(ctx, filter, update, opts).Decode(&ret)
	_, err := tb.UpdateMany(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *MongoClientWrapper) DoFindOne(ptr interface{},
	dbName, tbName string,
	filter interface{},
	opts ...*options.FindOneOptions) error {
	var ctx = context.Background()
	db := r.Base
	tb := db.Database(dbName).Collection(tbName)

	ret := tb.FindOne(ctx, filter, opts...)
	if ret.Err() != nil {
		return ret.Err()
	}

	if err := ret.Decode(ptr); err != nil {
		return err
	}
	return nil
}

func (r *MongoClientWrapper) DoFindMany(retSlicePtr interface{},
	dbName, tbName string,
	filter interface{},
	opts ...*options.FindOptions) error {
	if !reflectUtils.IsSlice(retSlicePtr) || !reflectUtils.IsPointer(retSlicePtr) {
		return errors.New("not supported type,muse be pointer of slice,may be you need &slice")
	}

	var ctx = context.Background()
	db := r.Base
	tb := db.Database(dbName).Collection(tbName)

	c, err := tb.Find(ctx, filter, opts...)
	if err != nil {
		return err
	}

	l := reflect.Indirect(reflect.ValueOf(retSlicePtr))
	for c.Next(ctx) {
		if bean, err := reflectUtils.MakeSliceElemPtr(retSlicePtr); err == nil {
			//log.Println("--BBB-beanType---", reflect.TypeOf(bean), "-----------")

			if err = bson.Unmarshal(c.Current, bean); err != nil {
				return err
			}
			//log.Println("--lType----", l.Type(), "-----------")
			//log.Println("---beanType---", reflect.TypeOf(bean), "-----------")
			//log.Println("------", reflect.ValueOf(bean), "-----------")
			l = reflect.Append(l, reflect.ValueOf(bean))
		} else {
			return err
		}
	}

	//rewrite
	v := reflect.ValueOf(retSlicePtr)
	v.Elem().Set(l)
	return nil
}

func (r *MongoClientWrapper) DoCount(db, tb string, filter bson.D,
	opts ...*options.CountOptions) (int64, error) {
	return r.Table(db, tb).CountDocuments(
		context.Background(),
		filter,
		opts...)
}
func (r *MongoClientWrapper) DoCountMap(db,
	tb string, filter bson.M,
	opts ...*options.CountOptions) (int64, error) {
	return r.Table(db, tb).CountDocuments(
		context.Background(),
		filter,
		opts...)
}

func (r *MongoClientWrapper) DoPage(slicePtr interface{},
	db, tb string,
	bean *ypage.PageBean) error {

	bean.ValidateAdjust()

	allCount, err := r.DoCount(db, tb, bean.Where)
	if err != nil {
		return err
	}
	bean.RowsCount = allCount
	bean.PagesCount = bean.GetPagesCount()

	skip := bean.GetSkip()
	filter := bson.D{{}}
	if len(bean.Where) > 0 {
		filter = bean.Where
	}

	limit := int64(bean.RowsPerPage)
	opts := &options.FindOptions{
		Limit: &limit,
		Skip:  &skip,
	}

	if len(bean.Sort) > 0 {
		opts.Sort = bean.Sort
	}

	if err = r.DoFindMany(slicePtr, db, tb, filter, opts); err != nil {
		return err
	}
	//
	bean.Data = slicePtr
	return nil
}
func (r *MongoClientWrapper) DoPageMap(slicePtr interface{},
	db, tb string,
	bean *ypage.PageBeanMap) (*ypage.PageBeanMap, error) {

	bean.ValidateAdjust()

	allCount, err := r.DoCountMap(db, tb, bean.Where)
	if err != nil {
		return nil, err
	}
	bean.RowsCount = allCount
	bean.PagesCount = bean.GetPagesCount()

	skip := bean.GetSkip()
	filter := bson.M{}
	if len(bean.Where) > 0 {
		filter = bean.Where
	}

	limit := int64(bean.RowsPerPage)
	opts := &options.FindOptions{
		Limit: &limit,
		Skip:  &skip,
		Sort:  bean.Sort,
	}

	if len(bean.Sort) > 0 {
		opts.Sort = bean.Sort
	}

	if err = r.DoFindMany(slicePtr, db, tb, filter, opts); err != nil {
		return nil, err
	}
	//
	bean.Data = slicePtr
	return bean, nil
}

// don't use it,it not pass test case
func (r *MongoClientWrapper) DoAggregateOne(ptr interface{},
	dbName, tbName string,
	pip interface{},
	opts ...*options.AggregateOptions) error {
	var ctx = context.Background()
	db := r.Base
	tb := db.Database(dbName).Collection(tbName)
	//pip := bson.D{{"$match",bson.D{{"uid",uid}}},}

	//db.user_addr.aggregate([{$group:{"_id":"$uid",max:{$max:1}}}])
	// db.abc.aggregate([{$group:{"_id":"$ageisok","age":{$max:"$age"}}}])
	c, err := tb.Aggregate(ctx, pip, opts...)

	if err != nil {
		return err
	}
	for c.Next(ctx) {
		if err = bson.Unmarshal(c.Current, ptr); err != nil {
			return err
		}
		break
	}

	return nil
}
