package route

import (
	"github.com/kataras/iris"
	"time"
	"hdg.com/db-demo/src/server/common"
	"hdg.com/db-demo/src/server/model"
	"fmt"
)

func SetContext() {
	iris.UseFunc(func(ctx *iris.Context) {
		startAt := time.Now().UnixNano() / 1000000
		ctx.Set("startAt", startAt)
		ctx.Response.Header.Set("X-Powered-By", fmt.Sprintf("%s/v%s", common.ConfigurationContext.Server.Name, common.ConfigurationContext.Server.Version))
		ctx.Next()
	})
}

func InitRoute() {
	SetContext()
	AddRoute()
	iris.Listen(fmt.Sprintf("%s:%d", common.ConfigurationContext.Server.Host, common.ConfigurationContext.Server.Port))
}

func AppendRoute(route *model.Route) {
	r := iris.Party(route.Root)
	for k, v := range route.Rout {
		switch v.Type {
		case model.GET:
			r.Get(k, v.Handle)
		case model.POST:
			r.Post(k, v.Handle)
		default:
			r.Get(k, v.Handle)
		}
	}
}
