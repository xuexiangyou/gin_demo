package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xuexiangyou/gin_demo/apis/admin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.GET("/admin/info", admin.Info)
	r.GET("/admin/hello", admin.Hello)
	return r
}
