package main

import (
	"os"
	"os/signal"
	"syscall"

	"GO/работы/Скилбокс/finalwork/pkg/config/"

	"github.com/astrviktor/skillbox_diploma/pkg/server"
)

func main() {
	config.GlobalConfig = config.NewConfig("config.yaml")
	config.GlobalConfig = config.ForHerokuConfig(config.GlobalConfig)

	go server.StartServer()

	exit := make(chan os.Signal, 0)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	<-exit
}
