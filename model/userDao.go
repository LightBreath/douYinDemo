package model

import (
	"time"
)

type User struct {
	Id         int64     `json:"id" gorm:"column:id;type:varchar(64);primaryKey"`                   // 主键id
	Username   string    `json:"username" gorm:"column:username;type:varchar(32);notnull"`          // 用户名(邮箱)
	Password   string    `json:"password" gorm:"column:password;type:varchar(32);notnull"`          // 密码
	UserInfoId int64     `json:"user_info_id" gorm:"columun:user_info_id;type:varchar(64);notnull"` // 用户信息id
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at;type:datatime"`                 // 创建时间
	UpdatedAt  time.Time `json:"Updated_at" gorm:"column:updated_at;type:datatime"`                 // 修改时间
}

type UserInfo struct {
	Id        int64     `json:"id" gorm:"column:id;primaryKey"`                    // 主键id，与User表的UserInfoId配合
	Name      string    `json:"name" gorm:"column:name;"`                          // 用户的昵称
	IsFollow  bool      `json:"is_follow" gorm:"column:is_follew;"`                // 1：已关注，0：未关注
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:datatime"` // 创建时间
	UpdatedAt time.Time `json:"Updated_at" gorm:"column:updated_at;type:datatime"` // 修改时间
}
