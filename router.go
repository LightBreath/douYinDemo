package main

import (
	"douyinDemo/controller"

	"douyinDemo/model"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	model.InitDB()

	apiRouter := r.Group("/douyin")

	apiRouter.POST("/user/register/", controller.UserRegister)

	apiRouter.POST("/publish/action", controller.Publish)

}
