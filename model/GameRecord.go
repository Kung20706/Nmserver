package model

import (
	"BearApp/internal/database"
)

// GameName 帳號資料表
type GameName struct {
	ID     int64  `gorm:"column:id;primary_key;type:int(10);NOT NULL;DEFAULT:0"` // gorm 格式ID
	UserID string `gorm:"column:user_id;type:varchar(50);"`
	// OperationUserData 操作記錄資料表
	DateTime string `gorm:"column:datetime;type:varchar(50);"`
	G0       string `gorm:"column:g0;type:varchar(5);"`
	G1       string `gorm:"column:g1;type:varchar(5);"`
	G2       string `gorm:"column:g2;type:varchar(5);"`
	G3       string `gorm:"column:g3;type:varchar(5);"`
	G4       string `gorm:"column:g4;type:varchar(5);"`
	G5       string `gorm:"column:g5;type:varchar(5);"`
	// 完成後累積總得分
	Action   string  `gorm:"column:action;type:varchar(10);"`
	Duration float64 `gorm:"column:duration;example:"100"`

	Score int `gorm:"column:score;type:int(50);"`
}

// TableName 資料表
func (m GameName) TableName() string {
	return TableGameName
}

// Database 資料庫
func (m GameName) Database() database.Type {
	return DB
}
