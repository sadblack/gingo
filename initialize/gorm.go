package initialize

import (
	"github.com/songcser/gingo/config"
	"github.com/songcser/gingo/internal/app"
	"github.com/songcser/gingo/pkg/auth"
	"gorm.io/gorm"
	"os"
)

// Gorm 这里写死了，怎么着都是 mysql
// 作用是 连接 mysql，并返回
func Gorm() *gorm.DB {
	switch config.GVA_CONFIG.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

// RegisterTables 注册数据库表专用
// 根据 BaseUser 和 App 这两个结构体的名称和属性，自动创建表
func RegisterTables(db *gorm.DB) {
	err := db.Set("gorm:table_options", "CHARSET=utf8mb4").AutoMigrate(
		// 系统模块表
		auth.BaseUser{},
		app.App{}, // app表注册
	)
	if err != nil {
		os.Exit(0)
	}
}
