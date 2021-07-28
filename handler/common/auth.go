package common

import (
	"encoding/json"
	datastruct "BearApp/common/data_struct"
	errorCode "BearApp/common/error_code"
	constant "BearApp/constant"
)

// EncryptSession 加密session
func EncryptSession(ID int64, session string) (token string, apiErr errorCode.Error) {
	// 將資料組成json並加密後 設定Cookie
	// 組成json

	cookie, err := ToJSON(datastruct.SessionData{
		UserID:  ID,
		Session: session,
	})

	// 將資料組成json是否有誤
	if err != nil {
		return
	}

	// 加密
	token, err = constant.EncryptSession(cookie)

	// 加密cookie 是否有誤
	if err != nil {
		apiErr = errorCode.GetAPIError("encrypt_session_err", err)
		return
	}

	return
}

func ToJSON(v interface{}) (j []byte, err error) {
	j, err = json.Marshal(v)
	return
}
