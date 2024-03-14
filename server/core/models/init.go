package models

import (
	"fmt"
	"gcloud/core/define"
	"gcloud/core/internal/config"
	"log"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化数据库
func Init(dataSource string) *gorm.DB {
	// engine, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{})
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/gcloud?charset=utf8mb4&parseTime=True&loc=Local", define.MySQLUser, define.MySQLPassword, define.MySQLAddr)
	fmt.Println(dsn)
	engine, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Gorm New Engine Error:%v", err)
		return nil
	}

	return engine
}

// 初始化redis
func InitRedis(c config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     define.RedisAddr,     // or c.Redis.Addr
		Password: define.RedisPassword, // no password set
		DB:       0,                    // use default DB
	})
}
