package Accountapi

import (
	"BearApp/business/auth"
	datastruct "BearApp/common/data_struct"
	errorcode "BearApp/common/error_code"
	constant "BearApp/constant"
	"BearApp/handler/common"
	"BearApp/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
)

var uploadBucket = "kungside"
var Access_key = "AKIAJP5H6ZIN43MUVLGA"
var Secret_key = "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"

// AccountUpdateCharData 過關資料
// @Summary 修改會員資訊
// @Description
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body AccountUpdateChars true "修改會員資訊"
// @Success 200 {object} TokenRes "token"
// @Router /Account/UpdateCharData [post]
func AccountUpdateCharData(c *gin.Context) {
	var (
		input AccountUpdateChars
		err   error
		g     model.GameName
		res   OperationUserData
	)
	Gdb, err := model.NewModelDB(model.GameName{}, true)
	if err != nil {
		apiErr := errorcode.CheckGormConnError("get_db_conn", err)
		dataAPI := datastruct.ErrAPI{
			ErrorText: apiErr.ErrorText(),
		}
		c.JSON(http.StatusOK, dataAPI)
		log.Println(dataAPI)
		return

	}
	// 判斷輸入參數
	err = json.NewDecoder(c.Request.Body).Decode(&input)
	if err != nil {
		apiErr := errorcode.CheckGormConnError("parse_error", err)
		dataAPI := datastruct.API{
			ErrorText: apiErr.ErrorText(),
			Data:      input,
		}
		c.JSON(http.StatusOK, dataAPI)
		log.Println(dataAPI)
	}

	db, err := model.NewModelDB(model.Account{}, true)
	if err != nil {
		apiErr := errorcode.CheckGormConnError("get_db_conn", err)
		dataAPI := datastruct.ErrAPI{
			ErrorText: apiErr.ErrorText(),
		}
		c.JSON(http.StatusOK, dataAPI)
		log.Println(dataAPI)
		return

	}
	log.Print(input.Token)
	if len(input.Token) < 1 {
		log.Println("尚未加入token")
		return
	}

	//使session 換取登入操作權限
	result := model.Account{}

	authData, err := auth.DecryptSession(input.Token)
	if err != nil {
		log.Println("無效的內容 返回", authData)
		return
	}
	token, apiErr := common.EncryptSession(
		authData.UserID,
		input.Token,
	)
	if apiErr != nil {
		dataAPI := datastruct.API{
			ErrorCode: apiErr.ErrorCode(),
			ErrorText: apiErr.ErrorText(),
		}
		c.JSON(http.StatusOK, dataAPI)
		log.Println(dataAPI)
		return
	}

	totalcoin := input.Coins.C1 + input.Coins.C2 + input.Coins.C3 + input.Coins.C4 + input.Coins.C5 + input.Coins.C6 + input.Coins.C7 + input.Coins.C8 + input.Coins.C9 + input.Coins.C10 + input.Coins.C11 + input.Coins.C12
	totalscore := input.Scores.P1 + input.Scores.P2 + input.Scores.P3 + input.Scores.P4 + input.Scores.P5 + input.Scores.P6 + input.Scores.P7 + input.Scores.P8 + input.Scores.P9 + input.Scores.P10 + input.Scores.P11 + input.Scores.P12

	log.Print(input.Coins.C2, token, totalcoin, totalscore, input.Times.T)

	err = db.Where(
		"id = ?", authData.UserID,
	).Find(&result).Error
	if err != nil {
		apiErr := errorcode.CheckGormConnError("account_not_found", err)
		dataAPI := datastruct.ErrAPI{
			ErrorText: apiErr.ErrorText(),
		}
		c.JSON(http.StatusOK, dataAPI)
		log.Println(dataAPI)
		return
	}

	rdb, err := model.NewModelDB(model.Operation{}, true)
	if err != nil {
		apiErr := errorcode.CheckGormConnError("get_db_conn", err)
		dataAPI := datastruct.ErrAPI{
			ErrorText: apiErr.ErrorText(),
		}
		c.JSON(http.StatusOK, dataAPI)
		log.Println(dataAPI)
		return

	}
	//創了一個操作陣列
	items := []model.Operation{}

	//加入每個陣列內容數據
	//tt
	items = append(items, model.Operation{
		UserName: result.Alias,
		UserID:   result.UserID,
		T:        input.Times.T,
		P:        totalscore,
		C:        totalcoin,
		G0:       input.Graphic.G0,
		G1:       input.Graphic.G1,
		G2:       input.Graphic.G2,
		G3:       input.Graphic.G3,
		G4:       input.Graphic.G4,
		G5:       input.Graphic.G5,
	})
	//t1
	items = append(items, model.Operation{
		UserName: result.Alias,
		UserID:   result.UserID,
		T:        input.Times.T1,
		P:        input.Scores.P1,
		C:        input.Coins.C1,
	})
	//t2
	items = append(items, model.Operation{
		UserName: result.Alias,
		UserID:   result.UserID,
		P:        input.Scores.P1,
		ZR:       input.Others.Z2R,
		Y:        input.Times.T2_Y,
		M:        input.Times.T2_M,
		D:        input.Times.T2_D,
		W:        input.Times.T2_W,
		I:        input.Times.T2_I,
	})
	//t3
	items = append(items, model.Operation{
		UserName: result.Alias,
		UserID:   result.UserID,
		T:        input.Times.T3,
		P:        input.Scores.P3,
		C:        input.Coins.C3,
		T0:       input.Times.T3_0,
		T1:       input.Times.T3_1,
		T2:       input.Times.T3_2,
		T3:       input.Times.T3_3,
		T4:       input.Times.T3_4,
		T5:       input.Times.T3_5,
	})
	//t4
	items = append(items, model.Operation{
		UserName: result.Alias,
		UserID:   result.UserID,
		T:        input.Times.T4,
		P:        input.Scores.P4,
		C:        input.Coins.C4,
		T0:       input.Times.T4_0,
		T1:       input.Times.T4_1,
		T2:       input.Times.T4_2,
		T3:       input.Times.T4_3,
	})
	//t5
	items = append(items, model.Operation{
		UserName: result.Alias,
		UserID:   result.UserID,
		T:        input.Times.T5,
		P:        input.Scores.P5,
		C:        input.Coins.C5,
		T0:       input.Times.T5_0,
		T1:       input.Times.T5_1,
		T2:       input.Times.T5_2,
	})

	//t6
	items = append(items, model.Operation{
		UserName: result.Alias,
		UserID:   result.UserID,
		T:        input.Times.T6,
		P:        input.Scores.P6,
		C:        input.Coins.C6,
	})

	//t7
	items = append(items, model.Operation{
		UserName: result.Alias,
		UserID:   result.UserID,
		T:        input.Times.T7,
		P:        input.Scores.P7,
		C:        input.Coins.C7,
		ZN:       input.Others.Z7N,
	})
	//t8
	items = append(items, model.Operation{
		UserName: result.Alias,
		UserID:   result.UserID,
		T:        input.Times.T8,
		P:        input.Scores.P8,
		C:        input.Coins.C8,

		T0: input.Times.T8_0,
		T1: input.Times.T8_1,
		T2: input.Times.T8_2,
		T3: input.Times.T8_3,
		T4: input.Times.T8_4,
	})
	//t9
	items = append(items, model.Operation{
		UserName: result.Alias,
		UserID:   result.UserID,
		T:        input.Times.T9,
		P:        input.Scores.P9,
		C:        input.Coins.C9,
	})
	//t10
	items = append(items, model.Operation{
		UserName: result.Alias,
		UserID:   result.UserID,
		T:        input.Times.T10,
		P:        input.Scores.P10,
		C:        input.Coins.C10,
		T0:       input.Times.T10_0,
		T1:       input.Times.T10_1,
		T2:       input.Times.T10_2,
	})
	//t11
	items = append(items, model.Operation{
		UserName: result.Alias,
		UserID:   result.UserID,
		T:        input.Times.T11,
		T0:       input.Times.T11_0,
		P:        input.Scores.P11,
		C:        input.Coins.C11,
	})
	//t12
	items = append(items, model.Operation{
		UserName: result.Alias,
		UserID:   result.UserID,
		T:        input.Times.T12,
		P:        input.Scores.P12,
		C:        input.Coins.C12,
		T0:       input.Times.T12_0,
		T1:       input.Times.T12_1,
		T2:       input.Times.T12_2,
		T3:       input.Times.T12_3,
		T4:       input.Times.T12_4,
	})
	//第二關
	//第三關
	log.Print(items, "THIS WAY IN APPEND SUSS")
	// m[0].UserID = result.UserID
	// m[0].T = input.Times.T
	// m[0].P = totalscore
	// m[0].C = totalcoin
	// m[0].G0 = input.Graphic.G0
	// m[0].G1 = input.Graphic.G1
	// m[0].G2 = input.Graphic.G2
	// m[0].G3 = input.Graphic.G3
	// m[0].G4 = input.Graphic.G4
	// m[0].G5 = input.Graphic.G5
	// m = append(m, m[0]{
	// 	UserID:result.Username
	// })

	res.DateTime = time.Now()
	res.Duration = input.Times.T
	res.Score = totalscore
	g.DateTime = time.Now().Format("2006-01-02 15:04:05")
	g.Duration = input.Times.T
	g.Score = totalscore
	g.UserID = result.UserID
	g.G0 = input.Graphic.G0
	g.G1 = input.Graphic.G1
	g.G2 = input.Graphic.G2
	g.G3 = input.Graphic.G3
	g.G4 = input.Graphic.G4
	g.G5 = input.Graphic.G5
	// res := model.Operation{
	// 	UserID: result.UserID,
	// 	// T:      input.T,
	// 	// T0:     input.T0,
	// 	// T2:     input.T2,
	// 	// T2Y:    input.T2Y,
	// 	// T2M:    input.T2M,
	// 	// T2D:    input.T2D,
	// 	// T2W:    input.T2W,
	// 	// T2L:    input.T2L,
	// 	// T2R:    input.T2R,
	// 	// T2N:    input.T2N,
	// 	// T2I:    input.T2I,
	// 	// O2:     input.O2,
	// 	// P:      input.P,
	// 	// P2:     input.P2,
	// 	// PY:     input.PY,
	// 	// PM:     input.PM,
	// 	// PD:     input.PD,
	// 	// PW:     input.PW,
	// 	// PL:     input.PL,
	// 	// Z2R:    input.Z2R,
	// 	// Z2N:    input.Z2N,
	// 	// C:      input.C,
	// 	// C2:     input.C2,
	// }
	for i, _ := range items {
		err = rdb.Create(&items[i]).Error
		if err != nil {
			apiErr := errorcode.CheckGormConnError("parse_error", err)
			dataAPI := datastruct.ErrAPI{
				ErrorText: apiErr.ErrorText(),
			}
			c.JSON(http.StatusOK, dataAPI)
			log.Println(dataAPI)
		}
	}
	actionstate := strconv.Itoa(len(items)) + "/12"
	g.Action = actionstate
	err = Gdb.Create(&g).Error
	if err != nil {
		apiErr := errorcode.CheckGormConnError("parse_error", err)
		dataAPI := datastruct.ErrAPI{
			ErrorText: apiErr.ErrorText(),
		}
		c.JSON(http.StatusOK, dataAPI)
		log.Println(dataAPI)
	}
	dataAPI := datastruct.API{
		ErrorCode: "200",
		ErrorText: "更新成功",
		Data:      res,
	}
	c.JSON(http.StatusOK, dataAPI)
	log.Println(dataAPI)
}

// AccountStatus 登入取狀態
// @Summary 取金幣
// @Description
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body AccountStatusInput true "範例TOKEN取金幣"
// @Success 200 {object} TokenRes "token"
// @Router /Account/Status [post]
func AccountStatus(c *gin.Context) {
	var (
		input AccountStatusInput
		err   error
		// m      model.Operation
		result model.Account
	)

	// 判斷輸入參數
	err = c.ShouldBindJSON(&input)
	if err != nil {
		apiErr := errorcode.CheckGormConnError("parse_error", err)
		dataAPI := datastruct.ErrAPI{
			ErrorText: apiErr.ErrorText(),
		}
		c.JSON(http.StatusOK, dataAPI)
		log.Println(dataAPI)
		return
	}

	db, err := model.NewModelDB(model.Account{}, true)
	if err != nil {
		apiErr := errorcode.CheckGormConnError("get_db_conn", err)
		dataAPI := datastruct.ErrAPI{
			ErrorText: apiErr.ErrorText(),
		}
		c.JSON(http.StatusOK, dataAPI)
		log.Println(dataAPI)
		return
	}
	log.Print(input.Token)
	if len(input.Token) < 1 {
		log.Println("尚未加入token")
		return
	}

	authData, err := auth.DecryptSession(input.Token)
	if err != nil {
		log.Println("無效的內容 返回", authData)
		return
	}
	log.Print(authData.UserID)
	token, apiErr := common.EncryptSession(
		authData.UserID,
		input.Token,
	)
	log.Print("wwwws", token)
	if apiErr != nil {
		dataAPI := datastruct.API{
			ErrorCode: apiErr.ErrorCode(),
			ErrorText: apiErr.ErrorText(),
		}
		c.JSON(http.StatusOK, dataAPI)
		log.Println(dataAPI)
		return
	}

	// rdb, err := model.NewModelDB(model.Operation{}, true)
	// if err != nil {
	// 	apiErr := errorcode.CheckGormConnError("get_db_conn", err)
	// 	dataAPI := datastruct.ErrAPI{
	// 		ErrorText: apiErr.ErrorText(),
	// 	}
	// 	c.JSON(http.StatusOK, dataAPI)
	// 	log.Println(dataAPI)
	// 	return
	// }
	err = db.Where(
		"id = ?", authData.UserID,
	).Find(&result).Error
	if err != nil {
		apiErr := errorcode.CheckGormConnError("account_not_found", err)
		dataAPI := datastruct.ErrAPI{
			ErrorText: apiErr.ErrorText(),
		}
		c.JSON(http.StatusOK, dataAPI)
		log.Println(dataAPI)
		return
	}
	log.Print(result)
	res := Detail{}
	res.Money = result.Money
	res.UnLock1 = result.UnLock1
	res.UnLock2 = result.UnLock2
	res.UserID = result.UserID
	res.ID = authData.UserID

	dataAPI := datastruct.API{
		ErrorCode: "200",
		ErrorText: "",
		Data:      res,
	}
	c.JSON(http.StatusOK, dataAPI)
	log.Println(dataAPI)
}

// AccountUpdateStatus 上傳狀態{金幣,道具-1,道具-2}
// @Summary 修改金錢取狀態
// @Description
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body AccountUpdateStatusInput true "修改狀態"
// @Success 200 {object} TokenRes "token"
// @Router /Account/UpdateStatus [post]
func AccountUpdateStatus(c *gin.Context) {
	var (
		input AccountUpdateStatusInput
		err   error
		// m      model.Operation
		result model.Account
	)

	// 判斷輸入參數
	err = c.ShouldBindJSON(&input)
	if err != nil {
		apiErr := errorcode.CheckGormConnError("parse_error", err)
		dataAPI := datastruct.ErrAPI{
			ErrorText: apiErr.ErrorText(),
		}
		c.JSON(http.StatusOK, dataAPI)
		log.Println(dataAPI)
		return
	}

	db, err := model.NewModelDB(model.Account{}, true)
	if err != nil {
		apiErr := errorcode.CheckGormConnError("get_db_conn", err)
		dataAPI := datastruct.ErrAPI{
			ErrorText: apiErr.ErrorText(),
		}
		c.JSON(http.StatusOK, dataAPI)
		log.Println(dataAPI)
		return
	}
	log.Print(input.Token)
	if len(input.Token) < 1 {
		log.Println("尚未加入token")
		return
	}

	//使session 換取登入操作權限
	//

	authData, err := auth.DecryptSession(input.Token)
	if err != nil {
		log.Println("無效的內容 返回", authData)
		return
	}
	token, apiErr := common.EncryptSession(
		authData.UserID,
		input.Token,
	)
	log.Print("wwwws", token)
	if apiErr != nil {
		dataAPI := datastruct.API{
			ErrorCode: apiErr.ErrorCode(),
			ErrorText: apiErr.ErrorText(),
		}
		c.JSON(http.StatusOK, dataAPI)
		log.Println(dataAPI)
		return
	}

	// rdb, err := model.NewModelDB(model.Operation{}, true)
	// if err != nil {
	// 	apiErr := errorcode.CheckGormConnError("get_db_conn", err)
	// 	dataAPI := datastruct.ErrAPI{
	// 		ErrorText: apiErr.ErrorText(),
	// 	}
	// 	c.JSON(http.StatusOK, dataAPI)
	// 	log.Println(dataAPI)
	// 	return
	// }
	err = db.Where(
		"id = ?", authData.UserID,
	).Find(&result).Error
	if err != nil {
		apiErr := errorcode.CheckGormConnError("account_not_found", err)
		dataAPI := datastruct.ErrAPI{
			ErrorText: apiErr.ErrorText(),
		}
		c.JSON(http.StatusOK, dataAPI)
		log.Println(dataAPI)
		return
	}
	// vs := strconv.FormatInt(authData.UserID, 16)
	// log.Print(vs)
	err = db.Model(&result).Update(
		map[string]interface{}{
			"money":   input.Money,
			"unlock1": input.UnLock1,
			"unlock2": input.UnLock2,
		}).Error
	if err != nil {
		apiErr := errorcode.CheckGormConnError("parse_error", err)
		dataAPI := datastruct.ErrAPI{
			ErrorText: apiErr.ErrorText(),
		}
		c.JSON(http.StatusOK, dataAPI)
		log.Println(dataAPI)
	}

	res := Detail{}
	res.ID = authData.UserID
	res.Money = input.Money
	res.UnLock1 = input.UnLock1
	res.UnLock2 = input.UnLock2
	log.Print(res.UnLock1)

	res.UserID = result.UserID

	log.Print(authData.UserID, res.ID)
	dataAPI := datastruct.API{
		ErrorCode: "200",
		ErrorText: "",
		Data:      res,
	}
	c.JSON(http.StatusOK, dataAPI)
	log.Println(dataAPI)
}

// AccountGetCharData 取以往過關資料
// @Summary token get chatdata
// @Description
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body AccountStatusInput true "修改會員資訊"
// @Success 200 {object} TokenRes "token"
// @Router /Account/GetCharData [post]
func AccountGetCharData(c *gin.Context) {
	var (
		input  AccountStatusInput
		err    error
		result model.Account
		g      []model.GameName
		// res   OperationUserData
	)
	Gdb, err := model.NewModelDB(model.GameName{}, true)
	if err != nil {
		apiErr := errorcode.CheckGormConnError("get_db_conn", err)
		dataAPI := datastruct.ErrAPI{
			ErrorText: apiErr.ErrorText(),
		}
		c.JSON(http.StatusOK, dataAPI)
		log.Println(dataAPI)
		return

	}
	// 判斷輸入參數
	// 判斷輸入參數
	err = c.ShouldBindJSON(&input)
	if err != nil {
		apiErr := errorcode.CheckGormConnError("parse_error", err)
		dataAPI := datastruct.ErrAPI{
			ErrorText: apiErr.ErrorText(),
		}
		c.JSON(http.StatusOK, dataAPI)
		log.Println(dataAPI)
		return
	}

	db, err := model.NewModelDB(model.Account{}, true)
	if err != nil {
		apiErr := errorcode.CheckGormConnError("get_db_conn", err)
		dataAPI := datastruct.ErrAPI{
			ErrorText: apiErr.ErrorText(),
		}
		c.JSON(http.StatusOK, dataAPI)
		log.Println(dataAPI)
		return
	}
	log.Print(input.Token)
	if len(input.Token) < 1 {
		log.Println("尚未加入token")
		return
	}

	authData, err := auth.DecryptSession(input.Token)
	if err != nil {
		log.Println("無效的內容 返回", authData)
		return
	}
	log.Print(authData.UserID)
	token, apiErr := common.EncryptSession(
		authData.UserID,
		input.Token,
	)
	log.Print("wwwws", token)
	if apiErr != nil {
		dataAPI := datastruct.API{
			ErrorCode: apiErr.ErrorCode(),
			ErrorText: apiErr.ErrorText(),
		}
		c.JSON(http.StatusOK, dataAPI)
		log.Println(dataAPI)
		return
	}

	rdb, err := model.NewModelDB(model.Operation{}, true)
	if err != nil {
		apiErr := errorcode.CheckGormConnError("get_db_conn", err)
		dataAPI := datastruct.ErrAPI{
			ErrorText: apiErr.ErrorText(),
		}
		c.JSON(http.StatusOK, dataAPI)
		log.Println(dataAPI)
		return
	}
	operresult := []model.Operation{}
	err = db.Where(
		"id = ?", authData.UserID,
	).Find(&result).Error
	if err != nil {
		apiErr := errorcode.CheckGormConnError("account_not_found", err)
		dataAPI := datastruct.ErrAPI{
			ErrorText: apiErr.ErrorText(),
		}
		c.JSON(http.StatusOK, dataAPI)
		log.Println(dataAPI)
		return
	}
	err = rdb.Where(
		"user_id = ?", result.UserID,
	).Find(&operresult).Error
	if err != nil {
		apiErr := errorcode.CheckGormConnError("account_not_found", err)
		dataAPI := datastruct.ErrAPI{
			ErrorText: apiErr.ErrorText(),
		}
		c.JSON(http.StatusOK, dataAPI)
		log.Println(dataAPI)
		return
	}
	err = Gdb.Where(
		"user_id = ?", result.UserID,
	).Find(&g).Error
	if err != nil {
		apiErr := errorcode.CheckGormConnError("account_not_found", err)
		dataAPI := datastruct.ErrAPI{
			ErrorText: apiErr.ErrorText(),
		}
		c.JSON(http.StatusOK, dataAPI)
		log.Println(dataAPI)
		return
	}
	for i, _ := range operresult {
		operresult[i].UserName = result.Username
	}
	dataAPI := datastruct.API{
		ErrorCode: "200",
		ErrorText: "更新成功",
		Data:      g,
	}
	c.JSON(http.StatusOK, dataAPI)
	log.Println(dataAPI)
}

// UploadImage 上傳圖片
// @Summary  Request.FormFile("file") Header "tableid" "time"
// @Description 解析時間來上傳圖片
// @Tags UserAssets
// @Accept json
// @Produce json
// @Param body body AccountCreateInput true "修改的結構"
// @Success 200 {object} Accountres "成功即可呼叫其他API"
// @Router /Account/Assets/UploadImage [post]
func UploadImage(c *gin.Context) {
	var (
		jsonRawEscaped   json.RawMessage
		jsonRawUnescaped json.RawMessage
	)
	file, header, err := c.Request.FormFile("file")
	//upload to the s3 bucket
	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(Access_key, Secret_key, ""),
		// Endpoint:         aws.String(end_point),
		Region:           aws.String("ap-northeast-1"),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(false), //virtual-host style方式，不要修改
	})
	if err != nil {
		log.Print("w")
	}
	log.Print(c.GetHeader("tableid"), c.GetHeader("time"))
	jsonRawEscaped = []byte(c.GetHeader("time"))                                   // "\\u263a"
	jsonRawUnescaped, _ = constant.UnescapeUnicodeCharactersInJSON(jsonRawEscaped) // "☺"

	fmt.Println(string(jsonRawEscaped)) // {"HelloWorld": "\uC548\uB155, \uC138\uC0C1(\u4E16\u4E0A). \u263a"}
	fmt.Println(string(jsonRawUnescaped))
	// newpost := []model.AliceBooking{}
	db, err := model.NewModelDB(model.Account{}, true)
	if err != nil {
		apiErr := errorcode.CheckGormConnError("get_db_conn", err)
		dataAPI := datastruct.API{
			ErrorText: apiErr.ErrorText(),
		}
		c.JSON(http.StatusOK, dataAPI)
		log.Println(dataAPI)
		return
	}
	now := time.Now()
	starttime := now.Format("2006-01-02")
	realtime := strings.Split(string(jsonRawUnescaped), " ")[0]
	newpost := []model.Account{}

	domain := ".s3-ap-northeast-1.amazonaws.com/"
	imgurl := "https://" + uploadBucket + domain + header.Filename
	err = db.Where("reservation_time=?", starttime+" "+realtime).Find(&newpost).Update("url", imgurl).Error
	if err != nil {
		apiErr := errorcode.CheckGormConnError("account_not_found", err)
		dataAPI := datastruct.ErrAPI{
			ErrorText: apiErr.ErrorText(),
		}
		c.JSON(http.StatusOK, dataAPI)
		log.Println(dataAPI)
		return
	}

	uploader := s3manager.NewUploader(sess)
	up, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(uploadBucket),
		Key:    aws.String(header.Filename),
		Body:   file,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":    "Failed to upload file",
			"uploader": up,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"filepath": up,
	})
}
