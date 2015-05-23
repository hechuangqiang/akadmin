package action

import (
	"github.com/hechuangqiang/ak"
)

type loginFilter struct {
	excludedUrls []string //不需要过滤的url
}

func (l *loginFilter) Execute(ctx *ak.Context) bool {
	rp := ctx.Request.URL.Path
	for _, u := range l.excludedUrls {
		if rp == u {
			return true
		}
	}
	if _, ok := ctx.Session.Get(SESSIONUSER); ok {
		return true
	} else {
		ctx.Redirect("/reqlogin")
		return false
	}

}

func init() {
	lf := &loginFilter{}
	lf.excludedUrls = make([]string, 2)
	lf.excludedUrls = append(lf.excludedUrls, "/reqlogin")
	lf.excludedUrls = append(lf.excludedUrls, "/login")
	ak.AddFilter(lf)
}
