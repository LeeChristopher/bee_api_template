package packages

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"time"
)

var (
	Db      *gorm.DB
	Redis   *redis.Client
	Limiter *redis.Client
)

/*
 * 获取Mysql连接
 */
func GetDbClient() (err error) {
	dbConStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		beego.AppConfig.String("dbUsername"), beego.AppConfig.String("dbPassword"),
		beego.AppConfig.String("dbHost"), beego.AppConfig.String("dbDatabase"))
	db, err := gorm.Open("mysql", dbConStr)
	if err != nil {
		return err
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(1000)
	db.DB().SetConnMaxLifetime(time.Minute)
	Db = db
	return
}

/*
 * 获取Redis连接
 */
func GetRedisClient() (err error) {
	dbNumber, err := beego.AppConfig.Int("redisDatabase")
	if err != nil {
		return err
	}

	redisOption := &redis.Options{
		Addr: beego.AppConfig.String("redisHost"),
		DB:   dbNumber,
	}
	if len(beego.AppConfig.String("redisUsername")) > 0 {
		redisOption.Password = beego.AppConfig.String("redisPassword")
	}
	redisClient := redis.NewClient(redisOption)
	_, err = redisClient.Ping().Result()
	if err != nil {
		return err
	}
	Redis = redisClient

	return
}
