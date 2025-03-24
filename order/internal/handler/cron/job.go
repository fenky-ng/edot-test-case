package cron

import (
	"github.com/robfig/cron"
)

func AssignJob(
	handler *CronHandler,
	cron *cron.Cron,
) {
	cron.AddFunc("0 * * * * *", handler.UpdateOrderStatus)
}
