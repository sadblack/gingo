package app

import (
	"github.com/gin-gonic/gin"
	"github.com/songcser/gingo/pkg/router"
)

func InitRouter(g *gin.RouterGroup) {
	r := router.NewRouter(g.Group("app"))
	a := NewApi()
	// 给 这个空白路径， 绑定 增删查改一套 api
	r.BindApi("", a)
	//为 /api/v1/app/hello 建立映射
	r.BindGet("hello", a.Hello) //
}
