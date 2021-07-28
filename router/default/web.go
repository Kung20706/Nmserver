package router

import (

	// _ "./docs"

	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//LoadWebRouter 讀取路由
func LoadWebRouter(r *gin.RouterGroup) {
	// 靜態資源
	// 載入文件路由
	// docs.Init()
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//	創建帳號
	// docs.Init()
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

//Middleware 讀取路由
func Middleware(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("API_Key")
	if user == nil {
		// Abort the request with the appropriate error code
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	// 执行该中间件之后的逻辑
	c.Next()
}
