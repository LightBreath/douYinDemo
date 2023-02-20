package controller

import "github.com/gin-gonic/gin"

func UserRegister(c *gin.Context) {
	// 用户注册的逻辑：
	// 1. 查询传输来的用户名和密码
	//username := c.Query("username")
	//password := c.Query("password")

	// 2. 查询从数据库传来的用户名

	// 3. 对比两次用户名是否相同

	// 4.1 如果两次用户名相同，返回错误

	// 4.2 两次用户名不同，返回创建成功
}
