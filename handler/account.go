package Accountapi

import (
	"BearApp/business/auth"
	datastruct "BearApp/common/data_struct"
	errorcode "BearApp/common/error_code"
	"BearApp/common/helper"
	"BearApp/handler/common"
	"BearApp/model"
	"BearApp/repository"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// AccountLogin 登入帳號
// @Summary
// @Description
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body ExternalLoginInput true "輸入帳號密碼"
// @Success 200 {object} TokenRes "token"
// @Router /Account/External/Login [post]
func AccountExternalLogin(c *gin.Context) {

	//解析輸入資料
	var (
		result     []model.Account
		input      ExternalLoginInput
		totalCount int64
		sData      *datastruct.SessionData
	)

	err := json.NewDecoder(c.Request.Body).Decode(&input)
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
		dataAPI := datastruct.API{
			ErrorText: apiErr.ErrorText(),
			Data:      err,
		}
		c.JSON(http.StatusOK, dataAPI)
		log.Println(dataAPI)
		return
	}

	reply := TokenRes{}
	result = []model.Account{}
	//由客戶端傳送時間 確認裝置是否為此人使用

	err = db.Where(
		"facebook_id = ?", input.FacebookID,
	).Find(&result).Count(&totalCount).Error
	if err != nil {
		// 帳號不存在
		apiErr := errorcode.CheckGormConnError("account_not_found", err)
		dataAPI := datastruct.API{
			ErrorCode: apiErr.ErrorCode(),
			ErrorText: apiErr.ErrorText(),
			Data:      input,
		}
		c.JSON(http.StatusOK, dataAPI)
		log.Println(dataAPI)
		return
	}
	if totalCount != 0 {
		sData = new(datastruct.SessionData)
		sData, _ = repository.ResetSession(result[0].ID)

		reply.Token = sData.Session
	}

	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	if totalCount == 0 {
		newaccount := model.Account{
			UserID:     uuid[:],
			FacebookID: input.FacebookID,
			Open:       1,
		}

		err = db.Create(&newaccount).Error
		if err != nil {
			dataAPI := datastruct.DataApi{
				Status:  419,
				Message: "db_conn_fail",
			}
			c.JSON(http.StatusOK, dataAPI)
			log.Println(dataAPI)
			return
		}
		err = db.Where("facebook_id = ?", input.FacebookID).Find(&result).Error
		if err != nil {
			// 帳號不存在
			apiErr := errorcode.CheckGormConnError("account_not_found", err)
			dataAPI := datastruct.API{
				ErrorCode: apiErr.ErrorCode(),
				ErrorText: apiErr.ErrorText(),
				Data:      input,
			}
			c.JSON(http.StatusOK, dataAPI)
			log.Println(dataAPI)
			return
		}

		sData, _ = repository.ResetSession(result[0].ID)
		log.Print(sData.Session)
		// reply.Token = sData.Session
	}
	dataAPI := datastruct.DataApi{
		Status:  200,
		Message: "suscess",
		Data:    sData.Session,
	}
	c.JSON(http.StatusOK, dataAPI)
	return

}

// AccountLogin 登入帳號
// @Summary 登入帳號
// @Description 取id 將欄位內帳號資料刪除
// @Tags auth
// @Accept json
// @Produce json
// @Param body body AccountLoginInput true "登入參數"
// @Success 200 {object} Accountres "成功即可呼叫其他API"
// @Router /Account/Internal/Login [post]
func AccountLogin(c *gin.Context) {

	//解析輸入資料

	var (
		input AccountLoginInput
	)
	err := json.NewDecoder(c.Request.Body).Decode(&input)
	if err != nil {
		apiErr := errorcode.CheckGormConnError("parse_error", err)
		dataAPI := datastruct.API{
			ErrorText: apiErr.ErrorText(),
			Data:      input,
		}
		c.JSON(http.StatusOK, dataAPI)
		log.Println(dataAPI)
	}

	// 資料庫連線
	db, err := model.NewModelDB(model.Account{}, true)
	if err != nil {
		apiErr := errorcode.CheckGormConnError("get_db_conn", err)
		dataAPI := datastruct.API{
			ErrorText: apiErr.ErrorText(),
			Data:      err,
		}
		c.JSON(http.StatusOK, dataAPI)
		log.Println(dataAPI)
		return
	}

	//查詢相同名稱
	result := model.Account{}
	err = db.Where(
		"username = ?", input.Username,
	).Find(&result).Error
	if err != nil {
		// 帳號不存在
		dataAPI := datastruct.API{
			ErrorCode: "405",
			ErrorText: "AccountNotRegister",
		}
		c.JSON(http.StatusOK, dataAPI)
		log.Println(dataAPI)
		return
	}

	cryptpassword := helper.CryptPassword(input.Password)

	var usersession string
	//密碼若是一樣 檢查成功 取得session

	switch cryptpassword {
	case result.Password:
		var sData *datastruct.SessionData
		sData = new(datastruct.SessionData)
		sData, apiErr := repository.ResetSession(result.ID)

		log.Println("-------sData-------", sData, "----result----", result)
		// 要到使用者的redis
		if apiErr != nil {
			return
		}

		usersession = sData.Session
		token, apiErr := common.EncryptSession(
			result.ID,
			usersession,
		)
		if apiErr != nil {
			dataAPI := datastruct.API{

				ErrorText: apiErr.ErrorText(),
			}
			c.JSON(http.StatusOK, dataAPI)
			log.Println(dataAPI)
			return
		}

		log.Println("---->", token)
		res := IdTokenRes{}
		res.UserID = result.UserID
		res.Token = token
		c.JSON(http.StatusOK, datastruct.API{
			ErrorCode: "200",
			ErrorText: "登入成功",
			Data:      res,
		})

	default:
		res := IdTokenRes{}
		res.UserID = result.UserID
		res.Token = ""
		dataAPI := datastruct.API{
			ErrorCode: "400",
			ErrorText: "",
			Data:      res,
		}
		c.JSON(http.StatusOK, dataAPI)
	}

}

// AccountCreate 創建帳號
// @Summary
// @Description
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body AccountCreateInput true "輸入帳號密碼"
// @Success 200 {object} TokenRes "token"
// @Router /Account/Create [post]
func AccountCreate(c *gin.Context) {

	//解析輸入資料
	var (
		// result []model.Account
		input AccountCreateInput

		totalCount    int64
		missingParams string
	)

	err := json.NewDecoder(c.Request.Body).Decode(&input)
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
		dataAPI := datastruct.API{
			ErrorText: apiErr.ErrorText(),
			Data:      err,
		}
		c.JSON(http.StatusOK, dataAPI)
		log.Println(dataAPI)
		return
	}

	// result = []model.Account{}
	//驗證輸入參數
	{

		if input.Username == "" {
			missingParams = "missing_params"
		}

		if len(missingParams) > 0 {
			apiErr := errorcode.CheckGormConnError(missingParams, nil)
			dataAPI := datastruct.API{

				ErrorText: apiErr.ErrorText(),
			}
			c.JSON(http.StatusOK, dataAPI)
			log.Println(dataAPI)
			return
		}
	}
	password := helper.CryptPassword(input.Password)
	uuidWithHyphen := uuid.New()
	log.Print(input.Auth)
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	newaccount := model.Account{
		UserID:   uuid,
		Username: input.Username,
		Password: password,
		Phone:    input.Phone,
		Email:    input.Email,
		Birthday: input.Birthday,
		Gender:   input.Gender,
		Alias:    input.Alias,
		Auth:     input.Auth,
	}

	err = db.Where(
		"username = ?", input.Username,
	).Find(&newaccount).Count(&totalCount).Error
	if totalCount != 0 {
		apiErr := errorcode.CheckGormConnError("account_exists", nil)
		dataAPI := datastruct.API{
			ErrorCode: apiErr.ErrorCode(),
			ErrorText: apiErr.ErrorText(),
			Data:      input.Username,
		}
		c.JSON(http.StatusOK, dataAPI)
		log.Println(dataAPI)

	} else {
		// 寫入db
		err = db.Create(&newaccount).Error
		if err != nil {
			apiErr := errorcode.CheckGormConnError("create_db_conn", err)
			dataAPI := datastruct.API{

				ErrorText: apiErr.ErrorText(),
				Data:      input.Username,
			}
			c.JSON(http.StatusOK, dataAPI)
			log.Println(dataAPI)
			return
		}

		res := AccountCreateRes{}
		res.Username = newaccount.Username
		res.UID = newaccount.UserID

		// 創建帳號建立完成
		apiErr := errorcode.CheckGormConnError("create_success", nil)
		dataAPI := datastruct.API{
			ErrorCode: apiErr.ErrorCode(),
			ErrorText: apiErr.ErrorText(),
			Data:      res,
		}
		c.JSON(http.StatusOK, dataAPI)
		log.Println(dataAPI)
	}
}

// AccountUpdatePassword 修改密碼
// @Summary 修改密碼
// @Description
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body AccountUpdatePasswordInput true "重製密碼"
// @Success 200 {object} TokenRes "token"
// @Router /Account/ResetPassWord [post]
func AccountUpdatePassword(c *gin.Context) {
	var (
		input AccountUpdatePasswordInput
		err   error
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

	// password := helper.CryptPassword(input.Password)
	newpassword := helper.CryptPassword(input.NewPassword)

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

	// AccountInput.
	if len(input.Token) < 0 {
		log.Println("尚未加入token")
		return
	}
	authData, err := auth.DecryptSession(input.Token)
	if err != nil {
		log.Println("無效的內容 返回", authData)
		return
	}
	//使session 換取登入操作權限
	log.Println("Token", authData)
	result := model.Account{}
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

	}

	if result.IsFreeze {
		//帳號已凍結
		apiErr := errorcode.CheckGormConnError("account_freeze", err)
		dataAPI := datastruct.ErrAPI{
			ErrorText: apiErr.ErrorText(),
		}
		c.JSON(http.StatusOK, dataAPI)
		log.Println(dataAPI)
		return
	}

	err = db.Model(&result).Update("password", newpassword).Error
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
		ErrorText: "重製成功",
		Data:      result.Username,
	}
	c.JSON(http.StatusOK, dataAPI)
	log.Println(dataAPI)
}

// AccountUpdateData 修改會員資訊
// @Summary 修改會員資訊
// @Description
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body AccountUpdateDataInput true "修改會員資訊"
// @Success 200 {object} TokenRes "token"
// @Router /Account/UpdateData [post]
func AccountUpdateData(c *gin.Context) {
	var (
		input AccountUpdateDataInput
		err   error
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
	result := model.Account{}

	authData, err := auth.DecryptSession(input.Token)
	if err != nil {
		log.Println("無效的內容 返回", authData)
		return
	}
	token, apiErr := common.EncryptSession(
		result.ID,
		input.Token,
	)
	log.Print(token)
	if apiErr != nil {
		dataAPI := datastruct.API{
			ErrorCode: apiErr.ErrorCode(),
			ErrorText: apiErr.ErrorText(),
		}
		c.JSON(http.StatusOK, dataAPI)
		log.Println(dataAPI)
		return
	}
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
	err = db.Model(&result).Update(
		map[string]interface{}{
			"alias":    input.Alias,
			"phone":    input.Phone,
			"birthday": input.Birthday,
			"gender":   input.Gender,
			"auth":     input.Auth,
		}).Error
	if err != nil {
		apiErr := errorcode.CheckGormConnError("parse_error", err)
		dataAPI := datastruct.ErrAPI{
			ErrorText: apiErr.ErrorText(),
		}
		c.JSON(http.StatusOK, dataAPI)
		log.Println(dataAPI)
	}

	log.Print(result)
	dataAPI := datastruct.API{
		ErrorCode: "200",
		ErrorText: "",
		Data:      result,
	}
	c.JSON(http.StatusOK, dataAPI)
	log.Println(dataAPI)
}

// AccountList 回傳資訊列表
// @Summary 回傳資訊列表
// @Description
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body TokenRes true "回傳資訊列表"
// @Success 200 {object} TokenRes "token"
// @Router /Account/List [post]
func AccountList(c *gin.Context) {
	var (
		input      TokenRes
		err        error
		operresult []model.Operation
		us         []UserRes
		g          []model.GameName
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
	if len(input.Token) < 1 {
		log.Println("尚未加入token")
		return
	}
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
	authData, err := auth.DecryptSession(input.Token)
	if err != nil {
		log.Println("無效的內容 返回", authData)
		return
	}
	//使session 換取登入操作權限
	result := model.Account{}
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
	//用戶顯示自己
	if result.Auth != "1" {

		log.Print("rulechheck8888", result.Auth)

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
		log.Print(operresult)
		for i, _ := range operresult {

			operresult[i].UserName = result.Alias
		}
		if len(g) == 0 {
			dataAPI := datastruct.API{
				ErrorCode: "407",
				ErrorText: "尚未上船紀錄",
				Data:      us,
			}
			c.JSON(http.StatusOK, dataAPI)
			log.Println(dataAPI)
			return
		}

		dataAPI := datastruct.API{
			ErrorCode: "200",
			ErrorText: "更新成功",
			Data:      operresult,
		}
		c.JSON(http.StatusOK, dataAPI)
		log.Println(dataAPI)
		return
	}
	//管理者監看所有人列表
	if result.Auth == "1" {
		result := []model.Account{}
		err = db.Find(&result).Error
		log.Print("rulechheck", result)
		if err != nil {
			apiErr := errorcode.CheckGormConnError("account_not_found", err)
			dataAPI := datastruct.ErrAPI{
				ErrorText: apiErr.ErrorText(),
			}
			c.JSON(http.StatusOK, dataAPI)
			log.Println(dataAPI)
			return
		}
		err = rdb.Find(&operresult).Error
		if err != nil {
			apiErr := errorcode.CheckGormConnError("account_not_found", err)
			dataAPI := datastruct.ErrAPI{
				ErrorText: apiErr.ErrorText(),
			}
			c.JSON(http.StatusOK, dataAPI)
			log.Println(dataAPI)
			return
		}
		// us := []UserRes{}
		dataAPI := datastruct.API{
			ErrorCode: "200",
			ErrorText: "更新成功",
			Data:      operresult,
		}
		c.JSON(http.StatusOK, dataAPI)
		log.Println(result)
		return
	}
	token, apiErr := common.EncryptSession(
		result.ID,
		input.Token,
	)
	log.Print(token, apiErr)
}

// AccountQuery 回傳資訊列表
// @Summary 回傳資訊列表
// @Description
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body TokenRes true "回傳資訊列表"
// @Success 200 {object} TokenRes "token"
// @Router /Account/Query [post]
func AccountQuery(c *gin.Context) {
	var (
		input TokenQuery
		err   error
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
	if len(input.Token) < 1 {
		log.Println("尚未加入token")
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
	authData, err := auth.DecryptSession(input.Token)
	if err != nil {
		log.Println("無效的內容 返回", authData)
		return
	}
	//使session 換取登入操作權限
	result := model.Account{}
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
	//用戶顯示自己
	if result.Auth != "1" {

		log.Print("rulechheckkkk", result.Auth)
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
			"user_id = ?", input.UserID,
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
		for i, _ := range operresult {
			operresult[i].UserName = result.Username
		}
		dataAPI := datastruct.API{
			ErrorCode: "200",
			ErrorText: "更新成功",
			Data:      result,
		}
		c.JSON(http.StatusOK, dataAPI)
		log.Println(dataAPI)
		return
	}
	//管理者監看所有人列表
	if result.Auth == "1" {
		result := []model.Account{}
		err = db.Find(&result).Error
		log.Print("rulechheck", result)
		if err != nil {
			apiErr := errorcode.CheckGormConnError("account_not_found", err)
			dataAPI := datastruct.ErrAPI{
				ErrorText: apiErr.ErrorText(),
			}
			c.JSON(http.StatusOK, dataAPI)
			log.Println(dataAPI)
			return
		}
		c.JSON(http.StatusOK, result)
		log.Println(result)
		return
	}
	token, apiErr := common.EncryptSession(
		result.ID,
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
	operresult := []model.Operation{}
	err = db.Where(
		"id = ?", input.UserID,
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
	for i, _ := range operresult {
		operresult[i].UserName = result.Username
	}
	dataAPI := datastruct.API{
		ErrorCode: "200",
		ErrorText: "更新成功",
		Data:      operresult,
	}
	log.Print(token, dataAPI)
}

// GetUserData 取使用者資料
// @Summary 登入成功取使用者資訊
// @Description TOKEN => response(userdata)
// @Tags auth
// @Accept json
// @Produce json
// @Param body body AccountStatusInput true "登入參數"
// @Success 200 {object} Accountres "成功即可呼叫其他API"
// @Router /Account/GetUserData [post]
func GetUserData(c *gin.Context) {

	//解析輸入資料

	var (
		input AccountStatusInput
	)
	err := c.ShouldBindJSON(&input)
	if err != nil {
		apiErr := errorcode.CheckGormConnError("parse_error", err)
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
		dataAPI := datastruct.API{
			ErrorText: apiErr.ErrorText(),
			Data:      err,
		}
		c.JSON(http.StatusOK, dataAPI)
		log.Println(dataAPI)
		return
	}
	// 資料庫連線
	db, err := model.NewModelDB(model.Account{}, true)
	if err != nil {
		apiErr := errorcode.CheckGormConnError("get_db_conn", err)
		dataAPI := datastruct.API{
			ErrorText: apiErr.ErrorText(),
			Data:      err,
		}
		c.JSON(http.StatusOK, dataAPI)
		log.Println(dataAPI)
		return
	}
	authData, err := auth.DecryptSession(input.Token)
	if err != nil {
		log.Println("無效的內容 返回", authData)
		return
	}
	token, apiErr := common.EncryptSession(
		authData.UserID,
		input.Token,
	)

	//查詢相同名稱
	result := model.Account{}
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
	err = rdb.Where(
		"user_id = ?", result.UserID,
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
	var usersession string
	//密碼若是一樣 檢查成功 取得session

	var sData *datastruct.SessionData
	sData = new(datastruct.SessionData)
	sData, apiErr = repository.ResetSession(result.ID)

	log.Println("-------sData-------", sData, "----result----", result)
	// 要到使用者的redis
	if apiErr != nil {
		return
	}

	usersession = sData.Session
	token, apiErr = common.EncryptSession(
		result.ID,
		usersession,
	)
	if apiErr != nil {
		dataAPI := datastruct.API{

			ErrorText: apiErr.ErrorText(),
		}
		c.JSON(http.StatusOK, dataAPI)
		log.Println(dataAPI)
		return
	}

	log.Println("---->", token)
	res := AccountResponse{}
	res.Email = result.Email
	res.Birthday = result.Birthday
	res.Gender = result.Gender
	res.Phone = result.Phone
	res.Alias = result.Alias

	c.JSON(http.StatusOK, datastruct.API{
		ErrorCode: "200",
		ErrorText: "登入成功",
		Data:      res,
	})

}
