package Accountapi

import (
	"BearApp/model"
	"time"
)

// Detail 帳號資料表
type Detail struct {
	ID      int64  `gorm:"column:id;primary_key;type:int(20);NOT NULL;DEFAULT:0"` // gorm 格式ID
	UserID  string `gorm:"column:user_id;type:varchar(50);"`
	Money   int    `gorm:"column:money;type:int(50);NOT NULL;"`
	UnLock1 bool   `gorm:"column:unlock1;type:bool(1);NOT NULL;"`
	UnLock2 bool   `gorm:"column:unlock2;type:bool(1);NOT NULL;"`
}

// IdTokenRes 帳號資料表
type IdTokenRes struct {
	UserID string `gorm:"column:user_id;primary_key;type:int(20);NOT NULL;DEFAULT:0"` // gorm 格式ID
	Token  string `gorm:"column:token;type:varchar(50);"`
}

// AccountInput 帳號管理 輸入結構
type AccountInput struct {
	UserID int64 `json:"user_id" example:"1"`
	// 使用者名稱
	Username string `json:"username" example:"jay@gmail.com"`
	// 密碼
	Password string `json:"password" example:"qwe12345"`
	// 新密碼
	NewPassword string `json:"newpassword" example:"qwe12345"`
	// 管理者別名
	Alias string `json:"alias" example:"大傑"`
	//是否凍結 0為正常 1為鎖定
	IsFreeze bool `json:"is_freeze" example:"0"`
	//是否凍結 1為正常 2為鎖定
	Status int64 `json:"status" example:"1"`
	// 建立時間
	CreatedAt string `json:"created_at" example:"2018-08-14 16:02:32"`
	// 更新時間
	UpdatedAt string `json:"updated_at" example:"2018-08-20 16:02:32"`
	// 頁數
	Page int64 `json:"page" example:"2"`
	// 筆數
	PerPage int64 `json:"per_page" example:"20"`
}

// AccountResponse 回傳資料
type AccountResponse struct {
	Email    string `json:"email" example:"WERTYUIDSDDQMQWLKDQW"`
	Birthday string `json:"birthday" example:"yyyy/mm/dd"`
	Gender   string `json:"gender" example:"male"`
	Phone    string `json:"phone" example:"0987654321"`
	Alias    string `json:"alias" example:""`
}

// AccountCreateRes 回傳資料
type AccountCreateRes struct {
	Username string `json:"username" example:"jay@gmail.com"`
	UID      string `json:"uid" example:"weqn213oq3"`
}

// AccountData 回傳資料
type AccountData struct {
	// 頁數
	Page int64 `json:"page" example:"2"`
	// 一頁筆數
	PerPage int64 `json:"per_page" example:"20"`
	// 總筆數
	Total       int          `json:"total" example:"20"`
	AccountList []Accountres `json:"accountlist" `
}

// AccountData 回傳資料
type Accountres struct {
	ID               int64     `json:"UserID" example:"1"`
	Username         string    `json:"UserName" example:"jay@gmail.com"`
	Alias            string    `json:"Alias" example:"大傑"`
	LoginFailedCount int       `json:"LoginFailedCount" example:"1"`
	Status           int64     `json:"Status" example:"1"`
	IsFreeze         bool      `json:"IsFreeze" example:"0"`
	CreatedAt        time.Time `json:"CreatedAt" example:"2018-08-14 16:02:32"`
	UpdatedAt        time.Time `json:"UpdatedAt" example:"2018-08-20 16:02:32"`
}

// AccountUpdatePasswordInput 修改密碼
type AccountUpdatePasswordInput struct {

	// 新密碼
	NewPassword string `json:"username" example:"qwer5678"`
	//用於驗證身分
	Mail string `json:"mail" example:"qwer5678@gmail.com"`
}

//AccountUpdatePasswordReturnAPI 修改密碼
type AccountUpdatePasswordReturnAPI struct {
	ErrorText string `json:"error_text" example:"account_update_password_completed"`
	Data      string `json:"data" example:"jay@gmail.com"`
}

// AccountExternalCreateInput 創建帳號
type AccountExternalCreateInput struct {

	// FID
	FacebookID string `json:"FacebookID" example:"sdfedx3-01"`
	// GID
	GoogleID string `json:"GoogleID" example:"2132190dsa"`
}

// AccountCreateInput 登入帳號
type AccountCreateInput struct {
	Auth string `json:"auth" example:"1"`
	// 別名
	Alias string `json:"alias" example:"傑森"`
	// 使用者帳號
	Username string `json:"username" example:"qwq111@gmail.com"`
	// 密碼
	Password string `json:"password" example:"qwe12345"`
	// Email
	Email string `json:"email" example:"0987@gmail.com"`
	// Phone
	Phone string `json:"phone" example:"2132190dsa"`
	// AreaCode
	AreaCode string `json:"areacode" example:"405"`
	// 生日
	Birthday string `json:"birthday" example:"2018-08-14"`
	// Gender
	Gender string `json:"gender" example:"male"`
}

// AccountLoginInput 登入帳號
type AccountLoginInput struct {
	// 使用者名稱
	Username string `json:"username" example:"qwq111"`
	// 密碼
	Password string `json:"password" example:"qwe12345"`
	// FID
	// FacebookID string `json:"FacebookID" example:"sdfedx3-01"`
	// GID
	// GoogleID string `json:"GoogleID" example:"2132190dsa"`
}

// ExternalLoginInput 登入帳號
type ExternalLoginInput struct {
	// Line
	LineID string `json:"LineID" example:"sdfedx3-01"`
	// FID
	FacebookID string `json:"FacebookID" example:"sdfedx3-01"`
	// GID
	GoogleID string `json:"GoogleID" example:"2132190dsa"`
}

// CreateInput 登入帳號
type CreateInput struct {
	// 使用者名稱
	Username string `json:"username" example:"beartest2021@gmail.com"`
	// 密碼
	Password string `json:"password" example:"qwe12345"`
}

// AccountCheckInput 登入帳號
type AccountCheckInput struct {
	// 使用者名稱
	Username string `json:"username" example:"jay@gmail.com"`
}

// AccountListInput 查找列表 輸入頁面要幾個物件 有幾頁
type AccountListInput struct {
	// 頁數
	Page int64 `form:"page"`
	// 筆數
	PerPage int64 `form:"per_page"`
}

// AccountListReturnAPI
type AccountListReturnAPI struct {
	ErrorText string `json:"error_text" example:"account_list_completed"`
	Data      string `json:"data" example:""`
}

// AccountListReturnAPI
type AccountUpdateAliasInput struct {
	// 使用者名稱
	Username string `json:"username" example:"jay@gmail.com"`
	// 暱稱
	Alias string `form:"alias"`
}

//
type AccountUpdateAliasReturnAPI struct {
	ErrorText string `json:"error_text" example:"account_update_alias_completed"`
	Data      string `json:"data" example:""`
}

//
type ObjectSetInput struct {
	Userid string `json:"UserId" example:"1"`

	Objtype int64 `json:"ObjType" example:"1"`
	// 物件類型
	Objname string `json:"ObjName" example:"穿山甲"`
	// 物件名稱
	Objurl string `json:"ObjUrl" example:"http://aws.s3.com/file.png"`
	// 物件圖片地址
}

//ObjectSetReturnAPI  存儲物件用
type ObjectSetReturnAPI struct {
	ErrorText string `json:"error_text" example:"account_update_alias_completed"`
	Data      string `json:"data" example:""`
}

//ObjectListstruct 顯示物件列
type ObjectListstruct struct {
	Userid string `json:"UserId" example:"80b9284d1a"`
	Token  string `json:"Token" example:"218f0b43e69c7390699cb0506e6b01c4eb6c842cbd38dca335702ae440dac576"`
}

// AccountForgotPasswordInput 忘記密碼
type AccountForgotPasswordInput struct {
	// 使用者名稱
	Password string `json:"Password" example:"qwe123"`
}

// AccountOpenOnInput 開通帳號
type AccountOpenOnInput struct {
	// 使用者名稱
	Username string `json:"username" example:"jay@gmail.com"`
}

// AccountResetPasswordInput 修改密碼
type AccountResetPasswordInput struct {
	// 使用者名稱
	Username string `json:"username" example:"jay@gmail.com"`
	// 使用者唯一碼
	UserID string `json:"userid" example:"qwer5678"`
	// 新密碼
	Token string `json:"token" example:"qwer5678"`
}

// AccountUpdateInput 更新帳號
type AccountUpdateInput struct {
	// 持久性密鑰
	Token string `json:"Token" example:"qwer5678"`
	// 使用者名稱
	Alias string `json:"Alias" example:"彩虹傑森"`
	// 姓
	FirstName string `json:"FirstName" example:"javier"`
	// 名
	LastName string `json:"LastName" example:"zhang"`
	// 生日
	Birthday string `json:"Birthday" example:"2018-08-14"`
	// Gender
	Gender string `json:"Gender" example:"male"`
	// Phone number
	Phone string `json:"Phone" example:"0987654321"`
	// 使用者名稱
	Password string `json:"Password" example:"qwe123"`
}

// AccountPasswordInput 更新帳號
type AccountPasswordInput struct {
	// 使用者唯一碼
	UserID string `json:"userid" example:"qwer5678"`
	// 持久性密鑰
	Token string `json:"Token" example:"qwer5678"`
	// 使用者名稱
	Password string `json:"Password" example:"qwe123"`
}

// AccountBindInput 更新帳號
type AccountBindInput struct {
	// 持久性密鑰
	Wuid string `json:"Wuid" example:"kj21120sd3"`
	// 使用者名稱
	Fuid string `json:"Fuid" example:"sdfedx3-01"`
	// Guid
	Guid string `json:"Guid" example:"2132190dsa"`

	DeletePlatform int64 `json:"DeletePlatform" example:"1"`
}

// AccountReBindInput 更新帳號
type AccountReBindInput struct {
	// 持久性密鑰
	Token string `json:"Token" example:"kj21120sd3"`
	// 持久性密鑰
	UserID string `json:"UserID" example:"f314fb64fc"`
	// FID
	FacebookID string `json:"FacebookID" example:"sdfedx3-01"`
	// GIDe
	GoogleID string `json:"GoogleID" example:"2132190dsa"`
}

// AccountTokenRes 修改密碼
type AccountTokenRes struct {

	// 使用者唯一碼
	UserID string `json:"userid" example:"qwer5678"`
	// 新密碼
	Token string `json:"token" example:"qwer5678"`
}

//TokenRes 顯示物件列
type TokenRes struct {
	Token string `json:"Token" example:"218f0b43e69c7390699cb0506e6b01c4eb6c842cbd38dca335702ae440dac576"`
}

//TokenRes 顯示物件列
type TokenQuery struct {
	Token string `json:"Token" example:"218f0b43e69c7390699cb0506e6b01c4eb6c842cbd38dca335702ae440dac576"`

	UserID string `json:"user_id" example:"7390699cb05"`
}

// AccountUpdateDataInput 更新資料
type AccountUpdateDataInput struct {
	// Token
	Token string `json:"Token" example:"218f0b43e69c7390699cb0506e6b01c4eb6c842cbd38dca335702ae440dac576"`

	// Auth
	Auth string `json:"auth" example:"1"`
	// 使用者名稱
	Alias string `json:"alias" example:"小河"`
	// Phone
	Phone string `json:"phone" example:"2132190dsa"`
	// 生日
	Birthday string `json:"birthday" example:"2018-08-14"`
	// Gender
	Gender string `json:"gender" example:"male"`
}

// AccountUpdateChar
type AccountUpdateChar struct {
	Token string `json:"Token" example:"WERTYUIDSDDQMQWLKDQW"`
	// 開始遊戲到完成本關累計的花費時間
	T  string `json:"t" example:"100"`
	T0 string `json:"t0" example:"100"`
	// 關卡按下開始至下一關的時間
	T2 string `json:"t2" example:"1"`
	// 點 年 置完成輸入花費的時間
	T2Y string `json:"t2y" example:"1"`
	// 點 月 置完成輸入花費的時間
	T2M string `json:"t2m" example:"1"`
	// 點 日 置完成輸入花費的時間
	T2D string `json:"t2d" example:"1"`
	// 點 周 置完成輸入花費的時間
	T2W string `json:"t2w" example:"1"`
	// 點 地點 置完成輸入花費的時間
	T2L string `json:"t2l" example:"1"`
	// 本次玩家所使用的腳色
	T2R string `json:"t2r" example:"1"`
	// 點 夜市 置完成輸入花費的時間
	T2N string `json:"t2n" example:"1"`
	// 點 答題間格時間 置完成輸入花費的時間
	T2I string `json:"t2i" example:"1"`
	// 紀錄玩家的答題順序
	O2 string `json:"o2" example:"1"`
	// 完成後累積總得分
	P string `json:"p" example:"1"`
	// 本關總得分
	P2 string `json:"p2" example:"1"`
	// 年正確
	PY string `json:"py" example:"1"`
	// 月
	PM string `json:"pm" example:"1"`
	// 日
	PD string `json:"pd" example:"1"`
	// 週
	PW string `json:"pw" example:"1"`
	// 縣市
	PL string `json:"pl" example:"1"`
	// 本次玩家所使用的腳色
	Z2R string `json:"z2r" example:"1"`
	// 本次玩家所使用的腳色
	Z2N string `json:"z2n" example:"1"`
	// 本次玩家所使用的腳色
	C string `json:"c" example:"1"`
	// 本次玩家所使用的腳色
	C2 string `json:"c2" example:"1"`
}

// AccountUpdateChars
// type AccountUpdateChars struct {
// 	Token string `json:"Token" example:"WERTYUIDSDDQMQWLKDQW"`
// 	// 開始遊戲到完成本關累計的花費時間
// 	Data interface{} `json:"Data"`
// }

// AccountStatusInput
type AccountStatusInput struct {
	Token string `json:"Token" example:"WERTYUIDSDDQMQWLKDQW"`
	// 開始遊戲到完成本關累計的花費時間dd
}

// AccountStatusInput
type AccountUpdateStatusInput struct {
	Token string `json:"Token" example:"WERTYUIDSDDQMQWLKDQW"`

	//當前金幣-money int
	Money   int  `json:"money" example:1`
	UnLock1 bool `json:"unlock1" example:false`
	//道具(是否已購買)bool
	UnLock2 bool `json:"unlock2" example:false`
	//道具(是否已購買)bool

	// 開始遊戲到完成本關累計的花費時間dd
}
type ActionRes struct {
	TotalCoin  int     `json:"totalcoin"`
	TotalScore int     `json:"totalscore"`
	Times      float32 `json:"times"`
}

//使用者前端回傳
type UserRes struct {
	DateTime string
	UserName string
	Action   string
	Duration float64
	Score    int
	G0       string
	G1       string
	G2       string
	G3       string
	G4       string
	G5       string
}

type AccountUpdateChars struct {
	Token  string `json:"Token"`
	Scores struct {
		P2  int `json:"P2"`
		P3  int `json:"P3"`
		P1  int `json:"P1"`
		P4  int `json:"P4"`
		P5  int `json:"P5"`
		P6  int `json:"P6"`
		P7  int `json:"P7"`
		P8  int `json:"P8"`
		P9  int `json:"P9"`
		P10 int `json:"P10"`
		P11 int `json:"P11"`
		P12 int `json:"P12"`
	} `json:"Scores"`
	Coins struct {
		C2  int `json:"C2"`
		C3  int `json:"C3"`
		C1  int `json:"C1"`
		C4  int `json:"C4"`
		C5  int `json:"C5"`
		C6  int `json:"C6"`
		C7  int `json:"C7"`
		C8  int `json:"C8"`
		C9  int `json:"C9"`
		C10 int `json:"C10"`
		C11 int `json:"C11"`
		C12 int `json:"C12"`
	} `json:"Coins"`
	Others struct {
		Z2R string `json:"Z2r"`
		Z5S string `json:"Z5s"`
		Z5C string `json:"Z5c"`
		Z7N string `json:"Z7n"`
	} `json:"Others"`
	Times struct {
		T     float64 `json:"T"`
		T0    float64 `json:"T0"`
		T1    float64 `json:"T1"`
		T2    float64 `json:"T2"`
		T2_Y  float64 `json:"T2_y"`
		T2_M  float64 `json:"T2_m"`
		T2_D  float64 `json:"T2_d"`
		T2_W  float64 `json:"T2_w"`
		T2_I  float64 `json:"T2_i"`
		T3    float64 `json:"T3"`
		T3_0  float64 `json:"T3_0"`
		T3_1  float64 `json:"T3_1"`
		T3_2  float64 `json:"T3_2"`
		T3_3  float64 `json:"T3_3"`
		T3_4  float64 `json:"T3_4"`
		T3_5  float64 `json:"T3_5"`
		T4    float64 `json:"T4"`
		T4_0  float64 `json:"T4_0"`
		T4_1  float64 `json:"T4_1"`
		T4_2  float64 `json:"T4_2"`
		T4_3  float64 `json:"T4_3"`
		T5    float64 `json:"T5"`
		T5_0  float64 `json:"T5_0"`
		T5_1  float64 `json:"T5_1"`
		T5_2  float64 `json:"T5_2"`
		T6    float64 `json:"T6"`
		T7    float64 `json:"T7"`
		T8    float64 `json:"T8"`
		T8_0  float64 `json:"T8_0"`
		T8_1  float64 `json:"T8_1"`
		T8_2  float64 `json:"T8_2"`
		T8_3  float64 `json:"T8_3"`
		T8_4  float64 `json:"T8_4"`
		T9    float64 `json:"T9"`
		T10   float64 `json:"T10"`
		T10_0 float64 `json:"T10_0"`
		T10_1 float64 `json:"T10_1"`
		T10_2 float64 `json:"T10_2"`
		T11   float64 `json:"T11"`
		T11_0 float64 `json:"T110"`
		T12   float64 `json:"T12"`
		T12_0 float64 `json:"T12_0"`
		T12_1 float64 `json:"T12_1"`
		T12_2 float64 `json:"T12_2"`
		T12_3 float64 `json:"T12_3"`
		T12_4 float64 `json:"T12_4"`
	} `json:"Times"`
	Graphic struct {
		G0 string `json:"G0"`
		G1 string `json:"G1"`
		G2 string `json:"G2"`
		G3 string `json:"G3"`
		G4 string `json:"G4"`
		G5 string `json:"G5"`
	} `json:"Graphic"`
}

// OperationUserData 操作記錄資料表
type OperationUserData struct {
	DateTime time.Time `gorm:"column:t" example:"100"`

	// 完成後累積總得分
	Duration float64 `gorm:"column:p;type:int(50);"`

	Score int `gorm:"column:c;type:int(50);"`
}

// OperationUserDataS 操作記錄資料表
type OperationUserDataS struct {
	// 開始遊戲到完成本關累計的花費時間
	T float64 `gorm:"column:t" example:"100"`
	// 完成後累積總得分
	P int `gorm:"column:p;type:int(50);"`
	// 本次玩家所使用的腳色
	C int `gorm:"column:c;type:int(50);"`
	// 本次玩家所使用的腳色
	CreatedAt time.Time  // gorm 格式
	UpdatedAt time.Time  // gorm 格式
	DeletedAt *time.Time `sql:"index"` // gorm 格式
}

// Operationres 操作記錄資料表
type Operationres struct {
	ID       int64   `gorm:"column:id;primary_key;type:int(10);NOT NULL;DEFAULT:0"` // gorm 格式ID
	UserID   string  `gorm:"column:user_id;type:varchar(50);"`
	UserName string  `gorm:"column:user_id;type:varchar(50);"`
	T        float64 `gorm:"column:t" example:"100"`
	// 完成後累積總得分
	P int `gorm:"column:p;type:int(50);"`
	// 上傳檔案
	URL string `gorm:"column:url;type:varchar(50);NOT NULL;"`

	C  int `gorm:"column:c;type:int(50);"`
	G0 string
	G1 string
	G2 string
	G3 string
	G4 string
	G5 string

	CreatedAt time.Time // gorm 格式
}

// OperationUser 操作記錄資料表
type OperationUser struct {
	UserName string `gorm:"column:username;type:varchar(50);"`

	Opes []model.Operation
}

