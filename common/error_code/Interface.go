package errorcode

// Error 專案錯誤定義
type Error interface {
	error
	ErrorCode() string
	ErrorText() string
}
