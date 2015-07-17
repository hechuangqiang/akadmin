package action

import (
	"log"

	"github.com/hechuangqiang/ak"
)

//首页
func index(c *ak.Context) {
	log.Println("已经进入indexaction")
	c.WriteTpl("index.html")
}
