package main

import (
	"github.com/xuexiangyou/gin_demo/router"
)

func main() {
	r := router.InitRouter()
	r.Run()
}
