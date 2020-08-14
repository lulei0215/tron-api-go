package main

import (
	"trontool/router"
)

func main() {
	//model.Init()
	gin := router.Router()
	// 指定地址和端口号
	gin.Run("localhost:9090")
}
