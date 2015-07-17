package action

import (
	"github.com/hechuangqiang/ak"
)

func init() {
	ak.AddRoute("/", index)
	ak.AddRoute("/reqlogin", reqLogin)
	ak.AddRoute("/login", login)
	ak.AddRoute("/logout", logout)
	ak.AddRoute("/users",func(c *ak.Context){
		c.WriteTpl("users.html")
	})
}
