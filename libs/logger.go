package libs

import (
	"github.com/gogf/gf/frame/g"
)

func Info(content string) {
	g.Log().Skip(1).Line(true).Info(content)
}

func Warn(content string) {
	g.Log().Skip(1).Line(true).Warning(content)
}

func Notice(content string) {
	g.Log().Skip(1).Line(true).Notice(content)
}

func Error(content string) {
	g.Log().Skip(1).Line(true).Error(content)
}

func Debug(content string) {
	g.Log().Skip(1).Line(true).Debug(content)
}
