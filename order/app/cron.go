package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	handler_cron "github.com/fenky-ng/edot-test-case/order/internal/handler/cron"
	"github.com/robfig/cron"
)

func startCron() error {
	config, usecase, err := Init()
	if err != nil {
		return fmt.Errorf("[CRON] error initializing: %w", err)
	}

	ch := handler_cron.InitCronHandler(handler_cron.InitCronHandlerCOptions{
		Config:  config,
		Usecase: usecase,
	})

	cron := initCron(ch)
	runCron(cron)

	return nil
}

func initCron(ch *handler_cron.CronHandler) *cron.Cron {
	c := cron.New()
	handler_cron.AssignJob(ch, c)
	return c
}

func runCron(cron *cron.Cron) {
	cron.Start()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)

	select {
	case <-quit:
		log.Println("[CRON] Shutting down CRON...")

		cron.Stop()

		log.Println("[CRON] Stopped")
	}
}
