package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Println("No command specified")
		return
	}

	var commandRegistered bool

	if args[0] == "rest" {
		commandRegistered = true
		if err := startRestAPI(); err != nil {
			log.Fatalf("Error starting REST API: %+v", err)
			return
		}
		return
	}

	if args[0] == "cron" {
		commandRegistered = true
		if err := startCron(); err != nil {
			log.Fatalf("Error starting CRON: %+v", err)
			return
		}
		return
	}

	if !commandRegistered {
		log.Println("Command not registered")
	}
}
