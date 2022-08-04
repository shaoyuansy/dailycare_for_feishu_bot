package message

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

var SceneService = sceneService{}

type sceneService struct{}

func (s *sceneService) GenerateParams(scene string) (g.Map, error) {
	switch scene {
	case "water":
		return g.Map{
			"time": g.MapStrStr{
				"content": gtime.Now().Format("Y-m-d H:i:s"),
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
			"params":   g.Slice{"time"},
			"template": `{"config":{"wide_screen_mode":true},"elements":[{"fields":[{"is_short":true,"text":{"content":"**⏰ 当前时间：** %s","tag":"lark_md"}}],"tag":"div"},{"fields":[{"is_short":true,"text":{"content":"**❤️ 温馨提示：** 工作再忙，也要按时补充水分","tag":"lark_md"}}],"tag":"div"},{"actions":[{"tag":"button","text":{"content":"👎 干一杯","tag":"lark_md"},"type":"default"},{"tag":"button","text":{"content":"🙂 干一口","tag":"lark_md"},"type":"default"},{"tag":"button","text":{"content":"👍 不渴","tag":"lark_md"},"type":"default"}],"tag":"action"}],"header":{"template":"orange","title":{"content":"📢 小非来啦","tag":"plain_text"}}}`,
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
