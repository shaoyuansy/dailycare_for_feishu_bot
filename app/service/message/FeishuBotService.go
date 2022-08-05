package message

import (
	"errors"
	logger "informal/libs"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

var FeishuBotService = feishuBotService{}

type feishuBotService struct{}

func (p *feishuBotService) SendMessage(msg string) error {
	if msg == "" {
		return errors.New("[生成消息参数][20005]消息发送失败，原因:消息体不能为空")
	}
	url := g.Cfg().GetString("urls.InformalBotHookUrl")
	c := g.Client()
	resp := c.ContentJson().PostContent(url, g.Map{
		"msg_type": "interactive",
		"card":     gconv.String(msg),
	})
	if j, err := gjson.DecodeToJson(resp); err != nil {
		return err
	} else {
		if j.GetInt("StatusCode") != 0 {
			return errors.New("[生成消息参数][20005]消息发送失败，原因:" + resp)
		} else {
			logger.Debug("[飞书Bot发送消息][10000]发送成功")
			return nil
		}

	}
}
