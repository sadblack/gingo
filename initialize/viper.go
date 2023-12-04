package initialize

import (
	//flag 用来解析 命令行参数
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	//用来处理 http 请求，可根据 url，找到 handler，并将 response 返回给前端
	"github.com/gin-gonic/gin"
	"github.com/songcser/gingo/config"
	//用来读取配置文件
	"github.com/spf13/viper"
	"os"
)

// Viper //
// 优先级: 命令行 > 环境变量 > 默认值
func Viper(path ...string) *viper.Viper {
	var cfg string
	//如果没有指定配置文件的路径
	if len(path) == 0 {
		//定义命令行参数        参数名称为 c，默认值为 "",   usage 表示 这个参数的描述说明
		flag.StringVar(&cfg, "c", "", "choose config file.")
		//解析命令行参数
		flag.Parse()
		//如果没有通过命令行指定
		if cfg == "" {
			/*
			   判断 internal.ConfigEnv 常量存储的环境变量是否为空
			   比如我们启动项目的时候，执行：GVA_CONFIG=config.yaml go run main.go
			   这时候 os.Getenv(internal.ConfigEnv) 得到的就是 config.yaml
			   当然，也可以通过 os.Setenv(internal.ConfigEnv, "config.yaml") 在初始化之前设置
			*/
			if configEnv := os.Getenv(ConfigEnv); configEnv == "" {
				switch gin.Mode() {
				case gin.DebugMode:
					cfg = ConfigDefaultFile
					fmt.Printf("正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, ConfigDefaultFile)
				case gin.ReleaseMode:
					cfg = ConfigReleaseFile
					fmt.Printf("正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, ConfigReleaseFile)
				case gin.TestMode:
					cfg = ConfigTestFile
					fmt.Printf("正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, ConfigTestFile)
				}
			} else {
				// internal.ConfigEnv 常量存储的环境变量不为空 将值赋值于config
				cfg = configEnv
				fmt.Printf("正在使用%s环境变量,config的路径为%s\n", ConfigEnv, cfg)
			}
			//即使通过命令行指定了，也没啥用，仍然会给它 写死
		} else {
			cfg = "config/config.yaml"
			fmt.Printf("正在使用命令行的-c参数传递的值,config的路径为%s\n", cfg)
		}
	} else {
		//配置文件的路径
		cfg = path[0]
	}

	// 上面写了一堆，目的就是获取了 cfg 这个变量的值

	//新建一个 viper 对象
	v := viper.New()
	//设置配置文件的路径
	v.SetConfigFile(cfg)
	//设置配置文件的格式
	v.SetConfigType("yaml")
	//读取配置文件
	err := v.ReadInConfig()
	//如果读取失败，抛异常
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	//开启 “监听配置文件的变化” 这个功能
	v.WatchConfig()
	// 配置文件变化后，会执行这个方法，相当于 发布 - 订阅机制 （类似于 zookeeper）
	v.OnConfigChange(
		//这里是一个函数，有点类似于 java 里的 lambda
		//配置文件变化后，会把相关信息封装成一个 event(事件)，然后再 发布出去
		//fsnotify，全称是 file system notify
		func(e fsnotify.Event) {
			//打印日志
			fmt.Println("config file changed:", e.Name)
			//刷新 config.GVA_CONFIG 这个配置对象
			if err = v.Unmarshal(&config.GVA_CONFIG); err != nil {
				//如果有异常，就打印一下
				fmt.Println(err)
			}
		})

	//读取文件，映射成 config.GVA_CONFIG
	if err = v.Unmarshal(&config.GVA_CONFIG); err != nil {
		//有异常，就输出
		fmt.Println(err)
	}

	return v
}
