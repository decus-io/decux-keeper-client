package main

import (
	"log"
	"time"

	"github.com/decus-io/decus-keeper-client/config"
	"github.com/decus-io/decus-keeper-client/db"
	"github.com/decus-io/decus-keeper-client/service"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("initial config error: %v", err)
	}
	if err := db.ConnectDatabase(); err != nil {
		log.Fatalf("connect database error: %v", err)
	}

	groupService := service.NewGroup()
	if err := groupService.SyncGroup(); err != nil {
		log.Fatalf("sync group error: %v", err)
	}

	keeperService := service.NewKeeper()
	go func() {
		for {
			if err := keeperService.Heartbeat(); err != nil {
				log.Printf("send heart error: %v", err)
			}
			time.Sleep(time.Minute)
		}
	}()

	withdrawService := service.NewWithdraw()
	go func() {
		for {
			if err := withdrawService.CheckNewWithdraw(); err != nil {
				log.Printf("check withdraw error: %v", err)
			}
			time.Sleep(time.Minute)
		}
	}()

	forever := make(chan bool)
	<-forever
}
