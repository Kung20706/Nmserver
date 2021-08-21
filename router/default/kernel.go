package router

import (
	docs "BearApp/docs"
	Accountapi "BearApp/handler"
	"log"
	"time"

	cors "github.com/itsjamie/gin-cors"

	// "net/http"
	"github.com/gin-contrib/sessions"
	redis "github.com/gin-contrib/sessions/redis"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	// 載入文檔
	ginSwagger "github.com/swaggo/gin-swagger"
	// _ "github.com/swaggo/gin-swagger/example/basic/docs"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// LoadRoutes 載入 routes
func LoadRoutes(r *gin.Engine) {

	docs.Init()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Use(static.Serve("/", static.LocalFile("./dist", true)))
	r.NoRoute(func(c *gin.Context) {
		c.File("./dist/index.html")
	})
	// r.Use(static.Serve("/", static.LocalFile("./static", true)))
	//	創建帳號
	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))
	//	刪除帳號
	// r.POST("/Account/Password", Accountapi.AccountUpdatePassword)
	// ws房間

	// r.GET("/wss/*token", Room)
	store, _ := redis.NewStore(10, "tcp", "redis:6379", "wdxr12345", []byte("secret"))
	r.Use(sessions.Sessions("Authorization", store))
	//改密碼確認信
	r.GET("/Account/MailReset/*username", Accountapi.AccountMailReset)
	r.POST("/Account/UpdateData", Accountapi.AccountUpdateData)
	r.POST("/Account/ResetPassWord", Accountapi.AccountUpdatePassword)
	r.POST("/Account/External/Login", Accountapi.AccountExternalLogin)
	r.POST("/Account/Internal/Login", Accountapi.AccountLogin)
	r.POST("/Account/Create", Accountapi.AccountCreate)
	r.POST("/Account/List", Accountapi.AccountList)
	r.POST("/Account/UpdateCharData", Accountapi.AccountUpdateCharData)
	r.POST("/Account/Status", Accountapi.AccountStatus)
	r.POST("/Account/UpdateStatus", Accountapi.AccountUpdateStatus)
	r.POST("/Account/GetCharData", Accountapi.AccountGetCharData)
	r.POST("/Account/GetUserData", Accountapi.GetUserData)
	r.POST("/Account/Query", Accountapi.AccountQuery)
	r.POST("/Account/PasswordRewrite", Accountapi.AccountPasswordRewrite)
	r.Run(":9487")
}
func middleware1(c *gin.Context) {

	log.Println("exec middleware1")
	//c.Next() 執行middleware後面接的function，執行完後再回到middleware繼續執行下去
	c.Next()

	if Accountapi.CheckInternal(c.Request.Header["Authorization"][0]) == 1 {
		log.Print("中介層 check authtype")
		c.Abort()
	}
	log.Println("after exec middleware1")
}

//外包api 透過這邊解析
func middleware2(c *gin.Context) {
	log.Println("exec middleware2")
	if Accountapi.CheckInternal(c.Request.Header["Authorization"][0]) == 1 {
		log.Print("check is invalid")

		c.Abort()
		log.Print("out is invalid")
	}

	c.Abort() //停止執行後面的hanlder，可以用來做auth
	// c.JSON(200, gin.H{"msg": "i'm fail..."})
}
