package auth

import (
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"
	"time"
	datastruct "BearApp/common/data_struct"
	"BearApp/common/helper"
	constant "BearApp/constant"
	"BearApp/model"

	"github.com/go-redis/redis"
)

// DecryptSession session解密
func DecryptSession(session string) (authData datastruct.SessionData, err error) {
	var data []byte
	data, err = constant.DecryptSession(session)
	if err != nil {
		log.Println("common.DecryptSession: 呼叫global.DecryptSession解密, 發生錯誤, ", session)
		return
	}

	// 開始解json

	err = helper.FromJSON(data, &authData)
	if err != nil {
		log.Println("common.DecryptSession: 解JSON失敗, ", err.Error())
		return
	}
	return
}

// GetUserBySession 用session取會員資料
func GetUserBySession(session string) (
	userID int64,
	loginTime time.Time,
	err error,
) {
	// 如果沒傳入連線，自動建立redis連線
	conn, apiErr := model.NewRedis(true)
	if apiErr != nil {
		log.Println("redis_conn_err")
	}

	// 從 redis 取 session
	key := getSessionRedisKey(session)
	r := conn.Get(key)
	err = r.Err()
	if err != nil && err != redis.Nil {
		log.Println("redis_conn_err")
		return
	}

	if err == redis.Nil {
		return
	}

	// 將字串轉成User資料
	userID, loginTime = getUserFromString(r.Val())

	return
}

// getSessionRedisKey 取session的redis中的key值
func getSessionRedisKey(session string) string {
	return fmt.Sprintf("Wonder:Auth:Session:%s", session)
}

// getUserFromString 從字串分解出使用者資料
func getUserFromString(userStr string) (userID int64, loginTime time.Time) {
	query, err := url.ParseQuery(userStr)
	if err != nil {
		return
	}

	val, err := strconv.ParseInt(strings.TrimSpace(query.Get("user_id")), 10, 64)
	if err == nil {
		userID = val
	}

	loginTime, _ = time.Parse(time.RFC3339Nano, query.Get("login_time"))
	return
}

