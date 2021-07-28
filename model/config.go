package model

import (
	"BearApp/internal/bootstrap"
	"BearApp/internal/database"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

// 資料表名稱
const (
	/*
	 * ===============
	 *       DB 定義出tablename
	 * ===============
	 */

	TableGameName    = "gamerecord"
	TableUserClothes = "user_clothes"
	TableAccount     = "account"
	TableOperation   = "operation"
)

// 資料庫常數
const (
	DB = database.Type("Bearlab") // 一般資料庫
)

func dbConf(t database.Type, master bool) *bootstrap.DBConf {
	switch t {
	case DB:
		if master {
			return bootstrap.GetAppConf().DBMaster
		}
		return bootstrap.GetAppConf().DBSlave
	default:
		panic("DB型態錯誤")
	}
}

// NewModelDB 建立新的Model的DB連線
func NewModelDB(m IModel, master bool) (*gorm.DB, error) {
	dbType := m.Database()
	return database.GetPoolDB(dbType, dbConf(dbType, master))
}

// IModel Model有的func
type IModel interface {
	TableName() string
	Database() database.Type
}

// NewRedis 建立新的Redis連線
func NewRedis(master bool) (*redis.Client, error) {
	if master {
		return database.GetPoolRedis(database.DefaultCacheType, bootstrap.GetAppConf().CacheMaster)
	}
	return database.GetPoolRedis(database.DefaultCacheType, bootstrap.GetAppConf().CacheSlave)
}
