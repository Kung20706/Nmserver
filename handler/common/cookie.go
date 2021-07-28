package common

import (
	"log"
	"net/http"
	"net/url"
	"time"
	"BearApp/constant"

	"github.com/gin-gonic/gin"
)

/*
| =======================================================
| 以下是設定cookie的部分
| =======================================================
*/

// SetSessionCookie 設定session的cookie
func SetSessionCookie(c *gin.Context, token string, expire time.Time) {
	SetCookie(c, constant.SessionCookieName, token, expire)
}

// CleanSessionCookie 清除session的Cookie
func CleanSessionCookie(c *gin.Context) {
	domain := ""
	referer, err := url.Parse(c.Request.Referer())
	if err == nil {
		domain = referer.Host
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    constant.SessionCookieName,
		Path:    "/",
		Expires: time.Unix(1, 0),
		Domain:  domain,
		MaxAge:  -1,
		// Secure:  true, // ---> 打開註解的話，只有在 HTTPS 才能Set-Cookie
	})
}

// SetCookie 設定cookie
func SetCookie(c *gin.Context, name, value string, expire time.Time) {
	domain := ""
	referer, err := url.Parse(c.Request.Referer())
	if err == nil {
		domain = referer.Host
		log.Println("API_Host:", referer)
	}
	log.Println("Key:", name, "Value:", value)
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    name,
		Value:   value,
		Expires: expire,
		Path:    "/",
		Domain:  domain,
		// Secure:  true, // ---> 打開註解的話，只有在 HTTPS 才能Set-Cookie
	})
}
// func CheckHeader(c *gin.Context,token){

// }