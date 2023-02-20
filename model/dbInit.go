package model

import (
	"douyinDemo/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() {
	db, err := gorm.Open(mysql.Open(config.ToString()), &gorm.Config{
		PrepareStmt:            true, //缓存预编译命令
		SkipDefaultTransaction: true, //禁用默认事务操作
	})
	if err != nil {
		panic(err)
	}

	// 自动迁移，会自动创建表、缺失的外键、约束等
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}
}
