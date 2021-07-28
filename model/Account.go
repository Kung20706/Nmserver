package model

import (
	"BearApp/internal/database"
	"time"
)

// Account 帳號資料表
type Account struct {
	ID               int64      `gorm:"column:id;primary_key;type:int(20);NOT NULL;DEFAULT:0"` // gorm 格式ID
	Auth             string     `gorm:"column:auth;type:varchar(30);NOT NULL;"`                // 權限 1為管理者 0為會員
	Open             int        `gorm:"column:open;type:int(1);NOT NULL;DEFAULT:0"`            // 是否已開通
	ResetPassword    int        `gorm:"column:reset_password;type:int(1);NOT NULL;DEFAULT:0"`  // 是否同意修改密碼
	UserID           string     `gorm:"column:user_id;type:varchar(50);"`
	FID              string     `gorm:"column:fid;type:varchar(50);"`
	GID              string     `gorm:"column:gid;type:varchar(50);"`
	FacebookID       string     `gorm:"column:facebook_id;type:varchar(50);"`
	DeviceID         string     `gorm:"column:device_id;type:varchar(100);"`
	Username         string     `gorm:"column:username;type:varchar(50);NOT NULL;"`              // 帳號
	Password         string     `gorm:"column:password;type:varchar(255);NOT NULL;"`             // 密碼
	Alias            string     `gorm:"column:alias;type:varchar(50);NOT NULL;"`                 // 暱稱
	Status           int64      `gorm:"column:status;type:int(2);NOT NULL;DEFAULT:0"`            // 狀態
	LoginFailedCount int        `gorm:"column:login_fail_count;type:int(11);NOT NULL;DEFAULT:0"` // 登入帳號錯誤次數
	IsFreeze         bool       `gorm:"column:is_freeze;type:tinyint(1);NOT NULL;DEFAULT:0"`     // 凍結帳號
	FirstName        string     `gorm:"column:firstname;type:varchar(30);NOT NULL;"`
	LastName         string     `gorm:"column:lastname;type:varchar(30);NOT NULL;"`
	Birthday         string     `gorm:"column:birthday;type:varchar(30);NOT NULL;"`
	Gender           string     `gorm:"column:gender;type:varchar(50);NOT NULL;"`
	AreaCode         string     `gorm:"column:areacode;type:varchar(50);NOT NULL;"`
	Email            string     `gorm:"column:email;type:varchar(50);NOT NULL;"`
	Phone            string     `gorm:"column:phone;type:varchar(50);NOT NULL;"`
	Money            int        `gorm:"column:money;type:int(50);NOT NULL;"`
	UnLock1          bool       `gorm:"column:unlock1;type:bool(1);NOT NULL;"`
	UnLock2          bool       `gorm:"column:unlock2;type:bool(1);NOT NULL;"`
	CreatedAt        time.Time  // gorm 格式
	UpdatedAt        time.Time  // gorm 格式
	DeletedAt        *time.Time `sql:"index"` // gorm 格式
}

// TableName 資料表
func (m Account) TableName() string {
	return TableAccount
}

// Database 資料庫
func (m Account) Database() database.Type {
	return DB
}
