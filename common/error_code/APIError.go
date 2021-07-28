package errorcode

// APIError API錯誤格式
type APIError struct {
	Code string `json:"error_code"`
	Text string `json:"error_text"`
}

// ErrorCode 錯誤代碼
func (e APIError) ErrorCode() string {
	return e.Code
}

// ErrorText 錯誤訊息
func (e APIError) ErrorText() string {
	return e.Text
}

// Error API錯誤訊息
func (e APIError) Error() string {
	return e.Text + " (" + e.Code + ")"
}
