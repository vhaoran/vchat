package ymongo

import (
	"context"
	"log"

	"github.com/vhaoran/vchat/lib/yconfig"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	X *MongoClientWrapper
)

func InitMongo(cfg yconfig.MongoConfig) error {
	cnt, err := NewMongoClientWrapper(cfg)
	if err != nil {
		return err
	}
	X = cnt
	return nil
}

func NewMongoClientWrapper(cfg yconfig.MongoConfig) (*MongoClientWrapper, error) {
	bean, err := newMongoClient(cfg)
	if err != nil {
		return nil, err
	}
	return &MongoClientWrapper{
		Base: bean,
	}, nil
}

func newMongoClient(cfg yconfig.MongoConfig) (*mongo.Client, error) {
	var err error
	var client *mongo.Client
	// for cluster
	//uri := "mongodb://localhost/argos?replicaSet=replset&authSource=admin"
	// for single host
	//mongodb://mongodb0.example.com:27017/admin

	//uri := "mongodb://root:password@192.168.0.99:27017/test?&authSource=admin"
	uri := cfg.URL
	//uri := "mongodb://root:password@192.168.0.99:27017/test"

	opts := options.Client()
	opts.ApplyURI(uri)

	//set other properties
	//todo whr
	opts.SetMaxPoolSize(cfg.Options.MaxPoolSize)
	opts.SetMinPoolSize(cfg.Options.MinPoolSize)
	//opts.SetMaxConnIdleTime(cfg.Options.MaxConnIdleTime * time.Second)

	if client, err = mongo.Connect(context.Background(), opts); err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return client, nil
}
