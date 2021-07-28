package repository

import (
	"crypto/sha256"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"
	"time"
	datastruct "BearApp/common/data_struct"
	errorcode "BearApp/common/error_code"
	"BearApp/model"

	"github.com/go-redis/redis"
)

// GetSession 取Session
func GetSession(userID int64) (*datastruct.SessionData, errorcode.Error) {

	// 如果沒傳入連線，自動建立redis連線
	conn, err := model.NewRedis(true)
	log.Println("redis conn to --->:", conn)
	if err != nil {
		return nil, errorcode.CheckRedisConnError("get_redis_conn", err)
	}
	key := getUserRedisKey(userID)
	r := conn.Get(key)
	err = r.Err()
	if err != nil && err != redis.Nil {
		return nil, errorcode.CheckRedisConnError("get_user_session_failed", err)
	}

	session := r.Val()

R:
	// 如果session為空字串，重新產生
	if session == "" {
		sData, apiErr := ResetSession(userID)
		if apiErr != nil {
			return nil, apiErr
		}
		sData.IsNew = true
		return sData, nil
	}

	userID, loginTime, apiErr := GetUserBySession(session)
	if apiErr != nil {
		if apiErr.ErrorCode() == "114030104" { // Session過期
			session = ""
			goto R
		}
		return nil, apiErr
	}

	sData := &datastruct.SessionData{
		UserID:    userID,
		Session:   session,
		LoginTime: loginTime,
	}
	return sData, nil
}

// CleanSession 清除session
func CleanSession(userID int64) (apiErr errorcode.Error) {

	// 如果沒傳入連線，自動建立redis連線
	conn, err := model.NewRedis(true)
	log.Println("connform--->:", conn)
	if err != nil {
		return nil
	}

	key := getUserRedisKey(userID)
	r := conn.Get(key)

	err = r.Err()
	if err != nil && err != redis.Nil {
		apiErr = errorcode.CheckRedisConnError("get_user_session_failed", err)
		return
	}

	removeKey := []string{key}
	session := r.Val()
	if err == nil {
		sessionKey := getSessionRedisKey(session)
		removeKey = append(removeKey, sessionKey)
	}

	err = conn.Del(removeKey...).Err()
	if err != nil {
		apiErr = errorcode.CheckRedisConnError("delete_user_session_failed", err)
		return
	}
	log.Println("delete connform--->:", conn, session, userID, key)
	return
}

// GetUserBySession 用session取會員資料
func GetUserBySession(session string) (
	userID int64,
	loginTime time.Time,
	apiErr errorcode.Error,
) {
	// 如果沒傳入連線，自動建立redis連線
	conn, err := model.NewRedis(true)
	if err != nil {
		apiErr = errorcode.CheckRedisConnError("get_redis_conn", err)
		return
	}
	// 從 redis 取 session
	key := getSessionRedisKey(session)
	r := conn.Get(key)
	err = r.Err()
	if err != nil && err != redis.Nil {
		apiErr = errorcode.CheckRedisConnError("get_user_session_failed", err)
		return
	}

	if err == redis.Nil {
		apiErr = errorcode.GetAPIError("session_is_expired", nil)
		return
	}

	// 將字串轉成User資料
	userID, loginTime = getUserFromString(r.Val())

	return
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

// ResetSession 重設Session
func ResetSession(userID int64) (*datastruct.SessionData, errorcode.Error) {
	// 如果沒傳入連線，自動建立redis連線
	conn, err := model.NewRedis(true)
	if err != nil {
		return nil, errorcode.CheckRedisConnError("get_redis_conn", err)
	}

	// 取使用者的 Redis Key
	key := getUserRedisKey(userID)

	now := time.Now()
	// 產生session
	session := encryptSession(userID, now.UnixNano())

	// 先取舊的session，用來清除
	getResult := conn.Get(key)
	err = getResult.Err()
	if err != nil && err != redis.Nil {
		return nil, errorcode.CheckRedisConnError("get_user_session_failed", err)
	}
	oldSession := getResult.Val()

	// 存入新session
	setResult := conn.Set(key, session, 0)
	err = setResult.Err()
	if err != nil {
		return nil, errorcode.CheckRedisConnError("set_user_session_failed", err)
	}

	// 清除舊的session
	if oldSession != "" {
		delResult := conn.Del(getSessionRedisKey(oldSession))
		if err := delResult.Err(); err != nil {
			return nil, errorcode.CheckRedisConnError("delete_old_session_failed", err)
		}
	}

	// 存入新的user資料
	userStr := setUserString(userID, now)
	sessionKey := getSessionRedisKey(session)
	err = conn.Set(sessionKey, userStr, 0).Err()
	if err != nil {
		return nil, errorcode.CheckRedisConnError("set_user_session_failed", err)
	}

	sData := &datastruct.SessionData{
		UserID:    userID,
		Session:   session,
		LoginTime: now,
	}
	return sData, nil
}

// getUserRedisKey 取使用者的redis中的key值
func getUserRedisKey(id int64) string {
	return fmt.Sprintf("Auth:User:ID_%d", id)
}

// getSessionRedisKey 取session的redis中的key值
func getSessionRedisKey(session string) string {
	return fmt.Sprintf("Auth:Session:%s", session)
}

// encryptSession 加密session
func encryptSession(userID, ts int64) (session string) {
	// 原資料
	data := fmt.Sprintf("user_id=%d&ts=%d", userID, ts)

	//加密字符串
	salt := "wEgyaPhGhbRfscwPWjpMHqpeHLHD7cK9"

	// 進行加密
	hashCode := sha256.Sum256([]byte(salt + data))

	// 轉成字串
	session = fmt.Sprintf("%x", hashCode)
	return
}

// setUserString 將使用者資料轉成字串
func setUserString(id int64, now time.Time) string {
	query := url.Values{}
	query.Set("user_id", strconv.FormatInt(id, 10))
	query.Set("login_time", now.Format(time.RFC3339Nano))
	return query.Encode()
}
