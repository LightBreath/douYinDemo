package model

type User struct {
	Id         int64 `gorm:"primary_key"`
	Username   string
	Password   string `gorm:"size:50;notnull"`
	UserInfoId int64
}
