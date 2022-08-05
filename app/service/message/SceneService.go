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
	switch scene {
	case "water":
		tip := ""
		tipsPoolResults, err := dao.TipsPool.Fields("content").Where("type", scene).Array()
		if err != nil {
			return g.Map{}, errors.New("[生成消息参数][20006]数据库查询错误，原因：" + err.Error())
		}
		if len(tipsPoolResults) > 0 {
			rand.Seed(time.Now().UnixNano())
			index := rand.Intn(len(tipsPoolResults))
			tip = gconv.String(tipsPoolResults[index])
		}
		return g.Map{
			"time": g.MapStrStr{
				"content": gtime.Now().Format("Y-m-d H:i:s"),
				"type":    "string",
			},
			"text": g.MapStrStr{
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
			"params":   g.Slice{"time", "text"},
			"template": `{"config":{"wide_screen_mode":true},"elements":[{"fields":[{"is_short":true,"text":{"content":"**⏰ 时间：** \n\t%s","tag":"lark_md"}}],"tag":"div"},{"fields":[{"is_short":true,"text":{"content":"**❤️ 非说：** \n\t%s","tag":"lark_md"}}],"tag":"div"}],"header":{"template":"wathet","title":{"content":"📢 喝水时间到啦","tag":"plain_text"}}}`,
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
