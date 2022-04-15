package cron

import (
	"github.com/nelsonkti/echo-framework/cron/base"
)

func RegisterCrons(address, password string) {
	cron := base.StartCronTab(address, password)

	//注册定时任务

	go cron.Run()
}
