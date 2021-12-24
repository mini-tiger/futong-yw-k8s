package main

import (
	"ftk8s/base/cfg"
	"ftk8s/router"

	"github.com/fvbock/endless"
)

func main() {
	cfg.InitAll()
	r := router.InitRouter()
	// 默认endless服务器会监听下列信号：
	// syscall.SIGHUP，syscall.SIGUSR1，syscall.SIGUSR2，syscall.SIGINT，syscall.SIGTERM和syscall.SIGTSTP
	// 接收到 SIGHUP 信号将触发`fork/restart` 实现优雅重启（kill -1 pid会发送SIGHUP信号）
	// 接收到 syscall.SIGINT或syscall.SIGTERM 信号将触发优雅关机
	// 接收到 SIGUSR2 信号将触发HammerTime
	// SIGUSR1 和 SIGTSTP 被用来触发一些用户自定义的hook函数
	if err := endless.ListenAndServe(cfg.AppConfObj.HttpAddr, r); err != nil {
		cfg.Mlog.Panic("failed to start service(futong-yw-k8s), error message: ", err.Error())
	}
}
