package router

import (
	defaultR "BearApp/router/default"

	"github.com/gin-gonic/gin"
)

// RouteProvider 路由提供者
func RouteProvider(r *gin.Engine) {
	defaultR.LoadRoutes(r)
}
