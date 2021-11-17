package main

import (
	"fmt"
	"gin_test/app/log"
	"gin_test/app/manager"
	"gin_test/app/order"
	"gin_test/conf"
	"gin_test/middleware"
	"gin_test/routers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"os"
)

func main() {

	// InitConfig viper初始化调用
	InitConfig()
	// 数据库连接初始化
	db := conf.InitDB()
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)
	// 分包陆游
	routers.Include(log.Routers, manager.Routers, order.Routers)
	// 初始化路由
	r := routers.Init()
	r.Use(middleware.LoggerToFile())
	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}
}

// InitConfig viper初始化
func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
