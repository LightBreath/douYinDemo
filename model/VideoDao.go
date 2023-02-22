package model

import (
	"errors"
	//"gorm.io/gorm"
	//"log"
	"sync"
	"time"
)
type Video struct {
	Id            int64     `json:"id" gorm:"column:id;type:bigint;primaryKey"`                                 // 主键id
	UserId        int64     `json:"user_id" gorm:"column:user_id;type:varchar(64);notnull"`                     // 用户信息id
	PlayUrl       string    `json:"play_url" gorm:"column:play_url;type:varchar(200);notnull"`                  // 视频播放地址
	CoverUrl      string    `json:"cover_url" gorm:"column:cover_url;type:varchar(200);notnull"`                // 视频封面地址
	FavoriteCount string    `json:"favorite_count" gorm:"column:favorite_count;type:int(10);notnull;default:0"` // 视频的点赞总数
	CommentCount  string    `json:"comment_count" gorm:"column:comment_count;type:int(10);notnull;default:0"`   // 视频的评论总数
	Title         string    `json:"title" gorm:"column:title;type:varchar(50);notnull"`                         //视频标题
	CreatedAt     time.Time `json:"created_at" gorm:"column:created_at;type:datetime"`                          // 创建时间
	UpdatedAt     time.Time `json:"Updated_at" gorm:"column:updated_at;type:datetime"`                          // 修改时间
}
type VideoDAO struct {
}

var (
	videoDAO  *VideoDAO
	videoOnce sync.Once
)

func NewVideoDAO() *VideoDAO {
	videoOnce.Do(func() {
		videoDAO = new(VideoDAO)
	})
	return videoDAO
}
func (v *VideoDAO) QueryVideoCountByUserId(userId int64, count *int64) error {
	if count == nil {
		return errors.New("QueryVideoCountByUserId count 空指针")
	}
	return db.Model(&Video{}).Where("user_info_id=?", userId).Count(count).Error
}

func (v *VideoDAO) QueryVideoListByUserId(userId int64, videoList *[]*Video) error {
	if videoList == nil {

		return errors.New("QueryVideoListByUserId videoList 空指针")
	}
	return db.Where("user_info_id=?", userId).
		Select([]string{"id", "user_info_id", "play_url", "cover_url", "favorite_count", "comment_count", "is_favorite", "title"}).
		Find(videoList).Error
}
