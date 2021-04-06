package main

import (
	"github.com/sohaha/saiyan"
	"github.com/sohaha/zlsgo/zcli"
	"github.com/sohaha/zlsgo/zlog"
	"github.com/sohaha/zlsgo/znet"
)

var (
	port = zcli.SetVar("port", "Server Port").String(":8181")
)

func main() {
	err := zcli.LaunchServiceRun("saiyan", "", run)
	if err != nil {
		zlog.Error(err)
	} else {
		stop()
	}
}

func run() {
	r := znet.New()

	// 初始化服务
	w, err := saiyan.New()
	if err != nil {
		r.Log.Fatal(err)
	}

	// 程序退出时同时关闭服务
	defer w.Close()

	// 绑定服务
	w.BindHttpHandler(r)

	r.SetAddr(*port)

	// 启动之后直接访问 http://127.0.0.1:8181 即可访问到 php 程序
	znet.Run()
}

func stop() {}
