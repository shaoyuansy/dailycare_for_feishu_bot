package message

import (
	"errors"
	"fmt"
	"informal/app/dao"
	"math/rand"
	"strings"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

var SceneService = sceneService{}

type sceneService struct{}

func (s *sceneService) GenerateParams(scene string) (g.Map, error) {
	tip, err := randTips(scene)
	if err != nil {
		return g.Map{}, err
	}

	switch scene {
	case "water":
		return g.Map{
			"time": g.MapStrStr{
				"content": gtime.Now().Format("Y-m-d H:i:s"),
				"type":    "string",
			},
			"tip": g.MapStrStr{
				"content": tip,
				"type":    "string",
			},
		}, nil
	case "finish_work":
		return g.Map{
			"time": g.MapStrStr{
				"content": gtime.Now().Format("Y-m-d H:i:s"),
				"type":    "string",
			},
			"tip": g.MapStrStr{
				"content": tip,
				"type":    "string",
			},
		}, nil
	default:
		return g.Map{}, errors.New("[生成消息参数][20003]消息模板不存在")
	}
}

func (s *sceneService) MatchScene(scene string, params g.Map) (string, error) {
	templateMap := g.Map{
		"water": g.Map{
			"params":   g.Slice{"time", "tip"},
			"template": `{"config":{"wide_screen_mode":true},"elements":[{"fields":[{"is_short":true,"text":{"content":"**⏰ 时间：** \n\t%s","tag":"lark_md"}}],"tag":"div"},{"fields":[{"is_short":true,"text":{"content":"**❤️ 非说：** \n\t%s","tag":"lark_md"}}],"tag":"div"}],"header":{"template":"wathet","title":{"content":"📢 喝水时间到啦","tag":"plain_text"}}}`,
		},
		"finish_work": g.Map{
			"params":   g.Slice{"time", "tip"},
			"template": `{"config":{"wide_screen_mode":true},"elements":[{"fields":[{"is_short":true,"text":{"content":"**⏰ 时间：** \n\t%s","tag":"lark_md"}}],"tag":"div"},{"fields":[{"is_short":true,"text":{"content":"**❗ 通知：** \n\t%s","tag":"lark_md"}}],"tag":"div"}],"header":{"template":"orange","title":{"content":"📢 号外号外～","tag":"plain_text"}}}`,
		},
	}

	// 根据场景进行校验
	value, ok := templateMap[scene]
	if !ok {
		return "", errors.New("[组装消息模板][20003]消息模板不存在")
	}
	template := gconv.Map(value)
	scanStrs := g.SliceAny{}
	for _, v := range gconv.SliceStr(template["params"]) {
		field, result := params[v]
		if !result {
			return "", errors.New("[组装消息模板][20004]消息模板参数错误：" + v)
		}
		fieldMap := gconv.Map(field)
		if fieldMap["type"] == "array" {
			scanStrs = append(scanStrs, strings.Join(gconv.SliceStr(fieldMap["content"]), `\n`))

		}
		if fieldMap["type"] == "string" {
			scanStrs = append(scanStrs, gconv.String(fieldMap["content"]))
		}
	}

	content := fmt.Sprintf(gconv.String(template["template"]), scanStrs...)
	return content, nil
}

// 查询数据库中某场景的随机文案
func randTips(scene string) (string, error) {
	tip := ""
	tipsPoolResults, err := dao.TipsPool.Fields("content").Where("type", scene).Array()
	if err != nil {
		return "", errors.New("[生成消息参数][20006]数据库查询错误，原因：" + err.Error())
	}
	if len(tipsPoolResults) > 0 {
		rand.Seed(time.Now().UnixNano())
		index := rand.Intn(len(tipsPoolResults))
		tip = gconv.String(tipsPoolResults[index])
		return tip, nil
	} else {
		return "", errors.New("[生成消息参数][20007]未配置提示信息，请尽快处理，场景:" + scene)
	}
}
