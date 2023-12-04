package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/songcser/gingo/config"
	"github.com/songcser/gingo/pkg/response"
	"github.com/songcser/gingo/utils"
	"net/http"
)

func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		//如果 recovered 是 ValidationErrors，并且强制转换成功了（变成了 errs）
		/*
			相当于
			if (recovered instanceof validator.ValidationErrors) {
				ValidationErrors errs = (ValidationErrors)recovered;
				...打印日志
			}
		*/
		if errs, ok := recovered.(validator.ValidationErrors); ok {
			response.FailWithDetailed(errs.Translate(utils.Trans), "入参校验失败", c)
			return
		}
		//这个也是强转，上面那个相当于一个更加具体的异常，下面这个就是个通用的异常
		if err, ok := recovered.(error); ok {
			config.GVA_LOG.Error(err.Error())
			response.FailWithError(err, c)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	})
}
