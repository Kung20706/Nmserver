package model

import (
	"BearApp/internal/database"
	"time"
)

// Operation 操作記錄資料表
type Operation struct {
	ID     int64  `gorm:"column:id;primary_key;type:int(10);NOT NULL;DEFAULT:0"` // gorm 格式ID
	UserID string `gorm:"column:user_id;type:varchar(50);"`
	// 開始遊戲到完成本關累計的花費時間
	T  float64 `gorm:"column:t" example:"100"`
	T0 float64 `gorm:"column:t0;type:decimal(7,6);"`
	T1 float64 `gorm:"column:t1" example:"100"`
	T2 float64 `gorm:"column:t2" example:"100"`
	T3 float64 `gorm:"column:t3" example:"100"`
	T4 float64 `gorm:"column:t4" example:"100"`
	T5 float64 `gorm:"column:t5" example:"100"`
	P  int     `gorm:"column:p;type:int(50);"`
	// 上傳檔案
	UserName  string    `gorm:"column:username;type:varchar(50);"`
	C         int       `gorm:"column:c;type:int(50);"`
	G0        string    `gorm:"column:g0;type:varchar(5);"`
	G1        string    `gorm:"column:g1;type:varchar(5);"`
	G2        string    `gorm:"column:g2;type:varchar(5);"`
	G3        string    `gorm:"column:g3;type:varchar(5);"`
	G4        string    `gorm:"column:g4;type:varchar(5);"`
	G5        string    `gorm:"column:g5;type:varchar(5);"`
	Y         float64   `gorm:"column:ty; example:"100"`
	M         float64   `gorm:"column:tm; example:"100"`
	D         float64   `gorm:"column:td; example:"100"`
	W         float64   `gorm:"column:tw; example:"100"`
	I         float64   `gorm:"column:ti; example:"100"`
	ZR        string    `gorm:"column:zr; type:varchar(15);"`
	ZS        string    `gorm:"column:zs; type:varchar(15);"`
	ZC        string    `gorm:"column:zc; type:varchar(15);"`
	ZN        string    `gorm:"column:zn; type:varchar(15);"`
	CreatedAt time.Time // gorm 格式
}

// TableName 資料表
func (m Operation) TableName() string {
	return TableOperation
}

// Database 資料庫
func (m Operation) Database() database.Type {
	return DB
}
