package utils
import (
	"douyinDemo/model"
	"douyinDemo/config"
	
	"fmt"
	
	"log"
	"path/filepath"
)

type VideResponse struct {
}
func NewFileName(userId int64) string {
	var count int64

	err := model.NewVideoDAO().QueryVideoCountByUserId(userId, &count)
	if err != nil {
		log.Println(err)
	}
	return fmt.Sprintf("%d-%d", userId, count)
}
// SaveImageFromVideo 将视频切一帧保存到本地
// isDebug用于控制是否打印出执行的ffmepg命令
func SaveImageFromVideo(name string, isDebug bool) error {
	v2i := NewVideo2Image()
	if isDebug {
		v2i.Debug()
	}
	v2i.InputPath = filepath.Join(config.Info.StaticSourcePath, name)
	v2i.OutputPath = filepath.Join(config.Info.StaticSourcePath, name)
	v2i.FrameCount = 1
	queryString, err := v2i.GetQueryString()
	if err != nil {
		return err
	}
	return v2i.ExecCommand(queryString)
}
