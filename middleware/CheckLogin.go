package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

// 驗證舊系統登入
func checkGoogleLogin(c *gin.Context) {
	sid, err := c.Cookie("mysession")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(sid)
	return
}
