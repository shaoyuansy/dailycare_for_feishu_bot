package main

import (
	_ "informal/boot"
	_ "informal/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
