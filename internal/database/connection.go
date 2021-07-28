package database

import (
	"log"
	"time"
	"BearApp/internal/bootstrap"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

// NewOrmConnectionWithConf 用Conf建立新DB連線
func NewOrmConnectionWithConf(conf *bootstrap.DBConf) (db *gorm.DB, err error) {
	connectName := getConnectName(
		"mysql",
		conf.Host,
		conf.Port,
		conf.DB,
		conf.Username,
		conf.Password, // 本地環境  ＥＣ還徑
	)
	log.Println("connectName:", connectName)
	db, err = gorm.Open("mysql", connectName)
	if bootstrap.GetAppConf().App.Env == "local" {
		db.LogMode(true)
	}
	return
}

// NewRedisConnWithConf 用設定建立Redis連線
func NewRedisConnWithConf(conf *bootstrap.CacheConf) *redis.Client {
	redisConn := redis.NewClient(&redis.Options{
		Addr:        conf.Host + conf.Port,
		Password:    conf.Password,
		DB:          conf.DB, // use default DB
		PoolSize:    conf.MaxConn,
		IdleTimeout: time.Minute,
	})
	return redisConn
}
