package config

import (
	"github.com/spf13/viper"
	"os"
)

var (
	RedisDb       string
	RedisAddress  string
	RedisPassword string
	RedisDbNumber int
)

var (
	MongoDatabase string
	MongoAddress  string
	MongoPassword string
	MongoPort     string
)

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	initRedisConfig()
	initMongoConfig()
}

func initRedisConfig() {
	RedisDb = viper.GetString("redis.db")
	RedisAddress = viper.GetString("redis.addr")
	RedisPassword = viper.GetString("redis.pwd")
	RedisDbNumber = viper.GetInt("redis.number")
}

func initMongoConfig() {
	MongoDatabase = viper.GetString("mongo.database")
	MongoAddress = viper.GetString("mongo.address")
	MongoPassword = viper.GetString("mongo.password")
	MongoPort = viper.GetString("mongo.port")
}
