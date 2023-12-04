package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/songcser/gingo/config"
	"github.com/songcser/gingo/initialize"
	"go.uber.org/zap"
	"time"
)

type server interface {
	ListenAndServe() error
}

func runServer(r *gin.Engine) {
	address := ":8080"
	s := initServer(address, r)

	if err := s.ListenAndServe(); err != nil {
		config.GVA_LOG.Error(err.Error())
	}
}

func initServer(address string, router *gin.Engine) server {
	//通过 endless 来监听端口， endless 相当于 tomcat
	// NewServer 的 函数签名是 NewServer(addr string, handler http.Handler)
	//addr 就是地址， handler 在这里就是 gin.Engine
	// gin.Engine 实现了 http.Handler 接口，可以处理 http 请求
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}

func main() {
	// 初始化配置信息
	config.GVA_VP = initialize.Viper()

	// 初始化日志
	config.GVA_LOG = initialize.Zap()

	//把日志信息 设为 全局
	zap.ReplaceGlobals(config.GVA_LOG)

	// 初始化数据库
	config.GVA_DB = initialize.Gorm() // gorm连接数据库
	if config.GVA_DB != nil {
		// 获取 db
		db, _ := config.GVA_DB.DB()
		// 程序结束前关闭数据库连接
		defer db.Close()
	} else {
		fmt.Println("数据库启动失败...")
		return
	}

	// 初始化Router，建立 url -> controller 的映射关系
	router := initialize.Routers()
	if router == nil {
		return
	}

	// 初始化Admin
	initialize.Admin(router)
	//启动服务器，开始监听 8080端口，绑定 router，等待请求
	runServer(router)
}
