package action

import (
	"github.com/hechuangqiang/ak"
)

//首页
func index(c *ak.Context) {
	c.WriteTpl("index.html")
}
