package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/xuexiangyou/gin_demo/models"
	"github.com/xuexiangyou/gin_demo/pkgs/app"
)

func Info(c *gin.Context) {
	user := models.NewUser()
	data, err := user.GetOne()
	if err != nil {
		panic(err)
	}
	app.Ok(c, data, "")
}
