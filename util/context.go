package util

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

type ApiContext struct {
	iris.Context
}

type ApiHandler func(ctx ApiContext) error

// 成功的返回
func (c *ApiContext) Success(payload interface{}) error {
	_, err := c.JSON(iris.Map{
		"code": 0,
		"msg":  "ok",
		"data": payload,
	}, context.JSON{
		StreamingJSON: true,
		UnescapeHTML:  true,
	})
	return err
}

// 失败的返回
func (c *ApiContext) Fail(error error, statusCode ...int) error {
	code := 1
	if len(statusCode) > 0 {
		code = statusCode[0]
	}
	_, err := c.JSON(iris.Map{
		"code": code,
		"msg":  error.Error(),
		"data": nil,
	}, context.JSON{
		StreamingJSON: true,
		UnescapeHTML:  true,
	})
	return err
}

// 包装 iris.Handler 使其返回error 方便结束请求
func ApiHandlerWrap(handler ApiHandler) iris.Handler {
	return func(c iris.Context) {
		ctx := ApiContext{c}
		handler(ctx)
	}
}
