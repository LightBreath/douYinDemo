package controller

import (
	"douyinDemo/model"
	"douyinDemo/utils"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func Publish(c *gin.Context) {
	//首先验证userID是否存在
	rawId, _ := c.Get("user_id")

	userId, exist := rawId.(int64)
	if !exist {
		PublishVideoError(c, "no exist userID")
		return
	}
	//title := c.PostForm("title")
	//上传文件
	data, err := c.MultipartForm()
	if err != nil {
		PublishVideoError(c, err.Error())
		return
	}
	files := data.File["data"]
	for _, file := range files {

		filename := utils.NewFileName(userId) //根据userId得到唯一的文件名
		//存放文件的位置
		savePath := filepath.Join("./static", filename)
		err = c.SaveUploadedFile(file, savePath)
		if err != nil {
			PublishVideoError(c, err.Error())
			continue
		}
		//截取一帧画面作为封面
		err = utils.SaveImageFromVideo(filename, true)
		if err != nil {
			PublishVideoError(c, err.Error())
			continue
		}
		PublishVideoOk(c, file.Filename+"上传成功")
	}

}

// 正确返回
func PublishVideoOk(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, model.CommonResponse{StatusCode: 0, StatusMsg: msg})
}

// 异常返回
func PublishVideoError(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, model.CommonResponse{StatusCode: 1,
		StatusMsg: msg})
}
