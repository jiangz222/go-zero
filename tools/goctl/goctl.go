package main

import (
	"github.com/jiangz222/go-zero/core/load"
	"github.com/jiangz222/go-zero/core/logx"
	"github.com/zeromicro/go-zero/tools/goctl/cmd"
)

func main() {
	logx.Disable()
	load.Disable()
	cmd.Execute()
}
