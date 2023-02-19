package main

import (
	"douyinDemo/controller/user"
	"douyinDemo/model"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	model.InitDB()

	apiRouter := r.Group("/douyin")

	apiRouter.POST("/user/register/", user.UserRegister)
}
