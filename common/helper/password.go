package helper

import (
	constant "BearApp/constant"
	"crypto/md5"
	"fmt"
	"log"
)

// CryptPassword 密碼加密
func CryptPassword(
	rawPassword string,
) (
	password string,
) {

	// 先將密碼MD5加密過一次
	firstMD5 := md5.Sum([]byte(rawPassword))
	firstPass := fmt.Sprintf("%x", firstMD5)

	// add some salt hehehe...
	salt := `"wEgyaPhGhbRfscwPWjpMHqpeHLHD7cK9"`
	secondMD5 := md5.Sum([]byte(firstPass + salt))
	password = fmt.Sprintf("%x", secondMD5)

	return
}

// CryptDIDTS 密碼加密
func CryptDIDTS(
	hashread string,
	// ts   string,
) (
	firstPass string,
) {
	// 先將密碼MD5加密過一次
	log.Print(hashread)
	firstMD5 := md5.Sum([]byte(hashread + constant.SecretKey))
	firstPass = fmt.Sprintf("%x", firstMD5)
	log.Print(firstPass)

	return
}
