package crontab

import (
	"informal/app/service/message"
	logger "informal/libs"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

var CrontabService = crontabService{}

type crontabService struct{}

func (p *crontabService) Register() {
	messions := g.Cfg().GetArray("crontabs.Missions")
	if len(messions) == 0 {
		logger.Info("[注册定时任务]没有定时任务可以注册")
		return
	}
	for _, each := range messions {
		mession := gconv.MapStrStr(each)
		name, ok := mession["name"]
		if !ok {
			logger.Warn("[注册定时任务][20001]注册定时任务错误，缺少key:name，已跳过")
			continue
		}
		cron, ok := mession["cron"]
		if !ok {
			logger.Warn("[注册定时任务][20001]注册定时任务错误，缺少key:cron，已跳过")
			continue
		}

		// _, err := gcron.Add(cron, func() {
		logger.Debug("[执行定时任务][10000]定时任务触发，场景：" + name + ",时间:" + cron)
		params, err := message.SceneService.GenerateParams(name)
		if err != nil {
			logger.Error(err.Error())
		}
		msg, err := message.SceneService.MatchScene(name, params)
		if err != nil {
			logger.Error(err.Error())
		}
		er := message.FeishuBotService.SendMessage(msg)
		if er != nil {
			logger.Error(err.Error())
		}
		// })
		// if err != nil {
		// 	logger.Error(fmt.Sprintf("[注册定时任务][20002]定时任务注册失败，场景：%s, 原因：%s", name, err.Error()))
		// } else {
		// 	logger.Info("[注册定时任务][10000]定时任务注册成功，场景：" + name)
		// }
	}
}
