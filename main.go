package main

import "github.com/gin-gonic/gin"

func main() {

	// 初始化gin框架
	r := gin.Default()

	// 初始化路由分配
	InitRouter(r)

	// 运行
	r.Run()
}
