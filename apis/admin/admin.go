package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/xuexiangyou/gin_demo/models"
	"github.com/xuexiangyou/gin_demo/pkgs/app"
)

func Info(c *gin.Context) {
	user := models.NewUser()
	data, err := user.GetOne()
	if err != nil && err != gorm.ErrRecordNotFound {
		panic(err)
	}
	if err == gorm.ErrRecordNotFound {
		app.Ok(c, "", "没有数据")
	} else {
		app.Ok(c, data, "")
	}
}
