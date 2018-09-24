package middleware

import (
	"github.com/kataras/iris"
)

func ParserPrimaryKey(key string) iris.Handler {
	return func(c iris.Context) {
		//cc := util.ApiContext{ c }
		id, _ := c.Params().GetInt(key)
		//if err != nil {
		//	cc.Fail(errors.New("ID格式不正确"))
		//	return
		//}
		c.Values().Set(key, id)
		c.Next()
	}
}

type Paging struct {
	Page   int // 第几页
	Limit  int // 数量
	Offset int // 计算出的偏移值
}

// 方便从URLParam中 解析分页参数c
func ParsePaging(c iris.Context) {
	page := c.URLParamIntDefault("page", 1)
	limit := c.URLParamIntDefault("per_page", 20)
	c.Values().Set("paging", Paging{
		page,
		limit,
		(page - 1) * limit,
	})
	c.Next()
}
