package cron

import (
    "echo-framework/cron/base"
    //github.com/robfig/cron      查看定时规则
)

func RegisterCrons(address, password string) {
    cron := base.StartCronTab(address, password)

    //注册定时任务

    go cron.Run()
}
