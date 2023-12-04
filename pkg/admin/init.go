package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/songcser/gingo/pkg/auth"
)

var (
	LoginPath    = "/admin/login/"
	RegisterPath = "/admin/register/"
	HomePath     = "/admin/"
)

func Init(r *gin.Engine, a Admin) {

	if a == nil {
		a = BaseAdmin{
			User: auth.BaseUser{},
		}
	}
	//添加一些 url 路径 -> controller 的映射，这里返回的是 html 页面，不是json
	r.HTMLRender = a.Render()
	r.GET(LoginPath, a.LoginView)
	r.GET(RegisterPath, a.RegisterView)
	r.POST(LoginPath, a.Login)
	r.POST(RegisterPath, a.Register)

	//这里也是添加了一些路径，不过这些带着参数呢
	g := r.Group(HomePath)
	g.Use(a.Auth())
	g.GET("/", a.Home)
	g.GET(":model", a.List)
	g.GET(":model/form", a.AddView)
	g.POST(":model/add", a.AddItem)
	g.GET(":model/:id", a.ViewItem)
	g.POST(":model/:id", a.EditItem)
	g.GET(":model/delete/:id", a.DeleteItem)
}
