package action

import (
	"github.com/hechuangqiang/ak"
)

var SESSIONUSER = "SESSIONUSER"

//请求登录
func reqLogin(c *ak.Context) {
	c.WriteTpl("login.html")
}

//登录
func login(c *ak.Context) {
	if c.Params["loginName"] == "admin" && c.Params["pwd"] == "123" {
		c.Session.Put(SESSIONUSER, 1)
		c.Redirect("/")
	} else {
		c.Data["msg"] = "用户名或密码错误"
		c.Redirect("/reqlogin")
	}
}

func logout(c *ak.Context) {
	c.Session.Del(SESSIONUSER)
	c.Redirect("/reqlogin")
}
