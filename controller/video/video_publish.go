package video

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)


func Publish()(c *gin.Context){
	//首先验证userID是否存在
	token := c.PostForm("token")
	_, exist :=用户登录信息[token]
	if  !exist{
		PublishVideoError(c,"no exist userID")
		return
	}
	//FormFile returns the first file for the provided form key
	data,err :=c.FormFile("data")
	if err != nil{
		PublishVideoError(c,err.Error())
		return
	}
	filename :=filepath.Base(data.filename)
	user :=用户登录信息[token]
	finalName :=fmt.Sprintf("%d_%s",user.Id,filename)//根据userId得到唯一的文件名
	savePath :=filepath.Join("./public",finalName)
	err = c.SaveUploadedFile(file, savePath)
	if err != nil {
		PublishVideoError(c, err.Error())
		return
	}
	//正确返回情形
    PublishVideoOk(c,finalName+"upload successfully")
}

//正确返回
func PublishVideoOk(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, models.CommonResponse{StatusCode: 0, StatusMsg: msg})
}
//异常返回
func PublishVideoError(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, models.CommonResponse{StatusCode: 1,
		StatusMsg: msg})
}
