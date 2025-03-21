package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fenky-ng/edot-test-case/user/internal/handler/rest"
	"github.com/gin-gonic/gin"
)

func startRestAPI() error {
	config, usecase, err := Init()
	if err != nil {
		return fmt.Errorf("[REST API] error initializing: %w", err)
	}

	api := rest.InitRestAPI(rest.InitRestAPIOptions{
		Config:  config,
		Usecase: usecase,
	})

	router := initRouter(api)
	err = start(router)
	if err != nil {
		return fmt.Errorf("[REST API] error starting: %w", err)
	}

	return nil
}

func initRouter(api *rest.RestAPI) *gin.Engine {
	router := gin.New()
	rest.AssignRoute(api, router)
	return router
}

func start(router *gin.Engine) error {
	address := ":9000"

	srv := &http.Server{
		Addr:    address,
		Handler: router,
	}

	log.Println("[REST API] Starting server at", address)

	listenAndServeError := make(chan error, 1)
	go func() {
		listenAndServeError <- srv.ListenAndServe()
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)

	select {

	case err := <-listenAndServeError:
		return err

	case <-quit:
		log.Println("[REST API] Shutting down server...")

		// The context is used to inform the server it has 10 seconds to finish
		// the request it is currently handling
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalf("[REST API] server shutdown error: %+v", err)
		}

		select {
		case <-ctx.Done():
			log.Println("[REST API] server shutdown timeout")
		}

		log.Println("[REST API] server exiting")

	}

	return nil
}
