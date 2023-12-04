package config

import (
	"github.com/robfig/cron/v3"
	"github.com/songcser/gingo/config/autoload"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Configuration struct {
	Domain string         `mapstructure:"domain" json:"domain" yaml:"domain"`
	DbType string         `mapstructure:"dbType" json:"dbType" yaml:"dbType"`
	Admin  autoload.Admin `mapstructure:"admin" json:"admin" yaml:"admin"`
	Mysql  autoload.Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Zap    autoload.Zap   `mapstructure:"zap" json:"zap" yaml:"zap"`
	JWT    autoload.JWT   `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}

var (
	GVA_CONFIG Configuration
	// GVA_DB gorm 用来连接数据库，也可以 根据 struct 的名字和属性，映射成数据库里的表
	GVA_DB *gorm.DB
	// GVA_LOG zap 用来打印日志
	GVA_LOG *zap.Logger
	// GVA_VP viper 是一个用来读取配置信息的库，可以从 yaml、json、ini...格式的文件里读取配置信息
	GVA_VP *viper.Viper

	GVA_JOB *cron.Cron
)
