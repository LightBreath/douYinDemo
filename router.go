package main

import (
	"douyinDemo/controller/user"
	"douyinDemo/controller/video"
	"douyinDemo/model"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	model.InitDB()

	apiRouter := r.Group("/douyin")

	apiRouter.POST("/user/register/", user.UserRegister)


	apiRouter.POST("/user/register/", video.Publish)
}
