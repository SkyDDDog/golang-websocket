package main

import (
	"demo04/config"
	"demo04/internal/cache"
	"demo04/internal/mongo"
	"demo04/internal/repository"
	"demo04/internal/service"
	"demo04/routes"
	"fmt"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

func main() {

	// 初始化配置
	config.InitConfig()
	// 初始化redis
	cache.InitCache()
	// 初始化数据库
	repository.InitDB()
	// 初始化MongoDB
	mongo.InitMongoDB()
	// 监听ws信道
	go service.Manager.Start()
	//
	ginRouter := routes.NewRouter()

	server := &http.Server{
		Addr:           fmt.Sprintf(":%s", viper.GetString("server.port")),
		Handler:        ginRouter,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := server.ListenAndServe()
	//err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
