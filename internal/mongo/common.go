package mongo

import (
	"context"
	"demo04/config"
	"demo04/pkg/util/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDBClient *mongo.Client

func InitMongoDB() {
	clientOptions := options.Client().ApplyURI("mongodb://" + config.MongoAddress + ":" + config.MongoPort)
	var err error
	MongoDBClient, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Infof("mongoDb连接失败, {}", err)
		panic(err)
	}
	log.Infof("MongoDb Connect Successfully")
}
