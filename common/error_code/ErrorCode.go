package errorcode

import (
	"fmt"
	"log"
	"runtime/debug"
	"strings"

	"BearApp/common/helper"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
)

var mappingAPIError = map[string]APIError{
	/**
	 * 共用相關
	 */
	"get_db_conn":           {"4009001", "取DB連線失敗"},
	"gorm_pool_is_timeout":  {"4009002", "Gorm連線池逾時"},
	"gorm_pool_is_closed":   {"4009003", "Gorm連線池已經關閉"},
	"gorm_pool_no_config":   {"4009004", "Gorm連線池尚未設定"},
	"db_rollback_error":     {"4009005", "DB ROLLBACK失敗"},
	"redis_pool_is_timeout": {"4009006", "Redis連線池逾時"},
	"redis_pool_is_closed":  {"4009007", "Redis連線池已經關閉"},
	"redis_pool_no_config":  {"4009008", "Redis連線池尚未設定"},
	"get_redis_conn":        {"4009009", "取Redis連線失敗"},
	"type_convert":          {"4009010", "型態轉換失敗"},
	"param_invalid":         {"4009011", "參數錯誤"},
	"parse_error":           {"4009012", "資料錯誤（資料解析錯誤）"},
	"template_set_error":    {"4009013", "模板設定錯誤"},
	"template_exe_error":    {"4009014", "模板執行錯誤"},
	"time_zone_error":       {"4009015", "時區錯誤"},
	"panic":                 {"4009016", "發生Panic錯誤！"},
	"encrypt_session_err":   {"4009100", "Token加密失敗"},
	"undefined_error":       {"4009998", "未知錯誤"},

	// 帳號密碼相關
	"create_success":       {"200", "帳號創建成功"},
	"success":              {"000200", "成功"},
	"primary_key":          {"000200", "成功"},
	"account_open_success": {"000200", "帳號開通成功"},

	"login_success":          {"000200", "帳號登入成功"},
	"mail_success":           {"000200", "郵件發送成功"},
	"set_success":            {"000200", "修改成功"},
	"verification_success":   {"000200", "令牌驗證成功"},
	"newpassword":            {"000200", "請設置新的密碼"},
	"open_successed":         {"000200", "帳號已開通"},
	"send_mail_success":      {"000200", "信箱發送成功"},
	"user_fmt_err":           {"000201", "帳號格式不符"},
	"password_fmt_err":       {"000202", "密碼格式不符"},
	"newpassword_fmt_err":    {"000203", "新密碼格式不符"},
	"alias_fmt_err":          {"000204", "使用者別名格式不符"},
	"account_not_found":      {"000206", "帳號未註冊"},
	"account_freeze":         {"000207", "帳號已凍結"},
	"account_deleted":        {"000208", "帳號已經刪除過"},
	"password_incorrect":     {"000209", "密碼錯誤"},
	"account_not_registered": {"000210", "帳號未註冊"},
	"account_not_open":       {"000211", "帳號尚未開通"},
	"open_success":           {"000212", "開通成功"},
	"not_agree":              {"000213", "請至信箱同意重置號碼"},
	"object_is_sell":         {"000214", "賣掉了"},
	"object_not_found":       {"000215", "無此物件"},
	"reset_error":            {"000216", "取回物品失敗"},
	"food_not_found":         {"000217", "找不到食物"},
	"access_not_found":       {"000218", "權限不足"},
	"reservation_is_exist":   {"000219", "無法訂位,預約存在"},
	"date_exists":            {"000200", "日期以存在"},
	"not_newpassword":        {"000217", "尚未同意修改密碼"},
	"token_verification_error": {"000218", "驗證失敗"},
	"order_canceled": {"000287", "驗證失敗"},
	"account_exists": {"000408", "帳號已註冊"},

	// 上傳檔案相關
	"upload_error":                           {"4000601", "上傳檔案錯誤"},
	"upload_file_write_db_error":             {"4000602", "上傳檔案寫入DB時錯誤"},
	"db_commit_error":                        {"4000603", "寫入DB註解時發生錯誤"},
	"upload_file_write_db_not_found_map_key": {"4000604", "上傳檔案寫入DB時找不到對應值"},
	"not_found_in_game_image_db":             {"4000605", "資料庫沒這筆game資料哦"},
}

// newAPIError 取API錯誤訊息
func newAPIError(text string, err error) Error {
	text = strings.TrimSpace(text)
	apiErroCode, ok := mappingAPIError[text]
	if !ok {
		apiErroCode = APIError{
			"4009999",
			fmt.Sprintf("未定義錯誤訊息 [ %s ]", text),
		}
	}

	if err != nil {
		apiErroCode.Text = fmt.Sprintf("%s (%s)", apiErroCode.Text, err.Error())
	}

	if err != nil {
		if text == "panic" {
			log.Println("🚧  🚧  🚧  ", string(debug.Stack()))
			log.Println("🚒  🚒  🚒")
		}
		log.Printf(
			"🎃  %s, 內部發生錯誤 [%s], %s 🎃\n",
			helper.MyCaller(),
			text,
			apiErroCode.Text,
		)
	}
	return &apiErroCode
}

// GetAPIError 由錯誤碼取得API回傳
func GetAPIError(text string, err error) APIError {
	if err != nil {
		log.Printf("發生錯誤: [%s] %s\n", text, err.Error())
	}

	api, ok := mappingAPIError[text]
	if !ok {
		return APIError{"9999", "Undefined Error (" + text + ")"}
	}
	return api
}

// GetGqlError 由錯誤碼取得Gql回傳
func GetGqlError(p graphql.ResolveParams, text string, err error) gqlerrors.FormattedError {
	if err != nil {
		log.Printf("GraphQL 發生錯誤: [%s] %s\n", text, err.Error())
	}

	// type FormattedError struct {
	// 	Message       string                    `json:"message"`
	// 	Locations     []location.SourceLocation `json:"locations"`
	// 	Path          []interface{}             `json:"path,omitempty"`
	// 	Extensions    map[string]interface{}    `json:"extensions,omitempty"`
	// 	originalError error
	// }

	apiErr := gqlerrors.NewFormattedError(text)
	if p.Info.Path != nil {
		apiErr.Path = []interface{}{p.Info.Path.Key}
	}

	return apiErr
}
