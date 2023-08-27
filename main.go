package main

import (
	"awesomeProject/common"
	"github.com/gin-gonic/gin"
)

func main() {
	//连接数据库
	common.InitDB()
	defer common.CloseDB()
	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run()) // listen and serve on
}
