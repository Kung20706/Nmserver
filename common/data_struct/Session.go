package datastruct

import "time"

// SessionData session的資料
type SessionData struct {
	UserID    int64
	LoginTime time.Time
	Auth      string
	IsNew     bool
	Session   string
}
