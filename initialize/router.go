package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/songcser/gingo/config"
	"github.com/songcser/gingo/internal/app"
	"github.com/songcser/gingo/middleware"
	"github.com/songcser/gingo/utils"
	"net/http"
)

func HealthCheck(g *gin.Context) {
	g.JSON(http.StatusOK, "ok...")
}

func Routers() *gin.Engine {

	if err := utils.Translator("zh"); err != nil {
		config.GVA_LOG.Error(err.Error())
		return nil
	}
	//只用 gin 默认的配置
	Router := gin.Default()
	//gin.SetMode(gin.DebugMode)

	//这里相当于添加 filter，添加一个 用来处理异常信息 的 filter
	Router.Use(middleware.Recovery())
	//添加一个 用来打印日志 的 filter
	Router.Use(middleware.Logger())
	HealthGroup := Router.Group("")
	{
		// 健康监测
		HealthGroup.GET("/health", HealthCheck)
	}
	//相当于指定 url 的前缀
	ApiGroup := Router.Group("api/v1")
	//初始化，绑定一些 router
	app.InitRouter(ApiGroup)

	return Router
}
