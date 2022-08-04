package boot

import (
	"informal/app/service/crontab"
	_ "informal/packed"

	"github.com/gogf/gf/frame/g"
)

func init() {
	s := g.Server()
	// 启动时不打印路由表
	s.SetDumpRouterMap(false)
	crontab.CrontabService.Register()
}
