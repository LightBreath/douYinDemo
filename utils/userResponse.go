package utils

import "douyinDemo/model"

type UserRegister struct {
	UserId int64  `json:"user_id"` // 用户id
	Token  string `json:"token"`   // 用户鉴权token
}

type UserLogin struct {
	UserId int64  `json:"user_id"` // 用户id
	Token  string `json:"token"`   // 用户鉴权token
}

type UserInfo struct {
	User model.UserInfo `json:"user"`
}
