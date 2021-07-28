package validation

import "regexp"

const (
	usernamePattern = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	passwordPattern = `^[a-z0-9A-Z@._\-]{8,20}$`
)

// IsUsername 檢驗是否為合法用戶名,合法字符有 0-9, A-Z, a-z,合法字符長度{8,20}
func IsUsername(name string) bool {
	changeNameType := []byte(name)
	usernameRegexp := regexp.MustCompile(usernamePattern)
	return usernameRegexp.Match(changeNameType)
}

// IsPassword 檢驗是否為合法密碼,合法字符有 0-9, A-Z, a-z,合法字符長度{8,20}
func IsPassword(Password string) bool {
	changePasswordType := []byte(Password)
	userPasswordRegexp := regexp.MustCompile(passwordPattern)
	return userPasswordRegexp.Match(changePasswordType)
}
