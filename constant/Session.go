package constant

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"log"
)

/*
| =======================================================
| 以下是加密、解密的部分
| =======================================================
*/

var (
	// 加密用
	commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
	salt     = "wEgyaPhGhbRfscwPWjpMHqpeHLHD7cK9"
)

// EncryptSession session加密
func EncryptSession(data []byte) (string, error) {
	// 創建加密算法aes
	c, err := aes.NewCipher([]byte(salt))
	if err != nil {
		return "", err
	}

	//加密字符串
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	session := make([]byte, len(data))
	cfb.XORKeyStream(session, data)
	// fmt.Printf("%s => %x\n", data, session)
	return fmt.Sprintf("%x", session), nil
}

// DecryptSession session解密
func DecryptSession(session string) ([]byte, error) {
	s, err := hex.DecodeString(session)
	// 創建加密算法aes
	c, err := aes.NewCipher([]byte(salt))
	if err != nil {
		log.Println("global.DecryptSession: 創建加密算法aes, 發生錯誤, ", session)
		return nil, err
	}

	// 解密字符串
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	data := make([]byte, len(s))
	cfbdec.XORKeyStream(data, s)
	return data, nil
}
